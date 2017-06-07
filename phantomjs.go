package phantomjs

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"syscall"
	"time"
)

type PhantomJS struct {
	name     string
	execPath string
	addr     string
	pid      string
	logFile  *os.File
	process  *os.Process

	options *Options
}

func (phantomJS *PhantomJS) Name() string {
	return phantomJS.name
}

func (phantomJS *PhantomJS) NewSession(cap *Capabilities) (*Session, error) {
	if cap == nil {
		cap = DesiredCapabilities()
	}
	args := map[string]interface{}{
		"requiredCapabilities": map[string]interface{}{},
		"desiredCapabilities":  cap,
	}

	res, err := Commands.NewSession.Execute(phantomJS.addr, args, phantomJS.options)
	if err != nil {
		return nil, err
	}
	if res.Code != ErrorCode_Success {
		return nil, fmt.Errorf("%s", res.Data)
	}

	var ret struct {
		Response
		Cap *Capabilities `json:"value"`
	}

	err = json.Unmarshal(res.Data, &ret)
	if err != nil {
		return nil, err
	}

	return NewSession(phantomJS.addr, ret.SessionId, ret.Cap, phantomJS.options), nil
}

func (phantomJS *PhantomJS) GetAllSessions() ([]*Session, error) {
	res, err := Commands.GetAllSessions.Execute(phantomJS.addr, nil, phantomJS.options)
	if err != nil {
		return nil, err
	}
	if res.Code != ErrorCode_Success {
		return nil, errors.New(string(res.Data))
	}

	var ret struct {
		Response
		Value []struct {
			Id  string        `json:"id"`
			Cap *Capabilities `json:"capabilities"`
		} `json:"value"`
	}

	err = json.Unmarshal(res.Data, &ret)
	if err != nil {
		return nil, err
	}
	if ret.Status != 0 {
		return nil, fmt.Errorf("%s", res.Data)
	}

	sessions := []*Session{}
	for _, item := range ret.Value {
		sessions = append(sessions, NewSession(phantomJS.addr, item.Id, item.Cap, phantomJS.options))
	}

	return sessions, nil
}

func (phantomJS *PhantomJS) start() error {
	var err error

	attr := os.ProcAttr{}
	attr.Dir = phantomJS.options.WorkDir
	attr.Env = os.Environ()
	attr.Files = []*os.File{nil, phantomJS.logFile, phantomJS.logFile}
	attr.Sys = &syscall.SysProcAttr{Setpgid: true}

	args := phantomJS.options.Args()
	phantomJS.process, err = os.StartProcess(phantomJS.execPath, args, &attr)
	if err != nil {
		return errors.New("start phantomjs service failed")
	}

	for i := 0; ; i++ {
		select {
		case <-time.After(1 * time.Second):
			if IsConnectAble(phantomJS.addr) {
				return nil
			}
		}
		if i == 30 {
			phantomJS.Quit()
			return fmt.Errorf("CAN NOT connect to service %s", phantomJS.addr)
		}
	}
}

func (phantomJS *PhantomJS) quit() error {
	defer func() {
		if phantomJS.logFile != nil {
			phantomJS.logFile.Close()
			os.RemoveAll(phantomJS.options.LogFilePath)
		}
		os.RemoveAll(phantomJS.options.CookiesFile)
		os.RemoveAll(phantomJS.options.WorkDir)
	}()

	if phantomJS.process != nil {
		err := phantomJS.process.Kill()
		if err != nil {
			return err
		}
		_, err = phantomJS.process.Wait()
		return err
	}
	return nil
}

func (phantomJS *PhantomJS) Quit() error {
	return phantomJS.quit()
}

func NewPhantomJS(port int, options *Options) (*PhantomJS, error) {
	cmd := "phantomjs"
	execPath, err := exec.LookPath(cmd)
	if err != nil {
		return nil, fmt.Errorf("please install phantomjs first")
	}

	return NewPhantomJSWithExecutePath(execPath, port, options)
}

func NewPhantomJSWithExecutePath(execPath string, port int, options *Options) (*PhantomJS, error) {
	var err error

	phantomJS := PhantomJS{}
	phantomJS.execPath = execPath
	phantomJS.options = &Options{}
	*phantomJS.options = *options

	if phantomJS.options.WorkDir == "" {
		phantomJS.options.WorkDir, err = ioutil.TempDir("/tmp", "phantomjs")
		if err != nil {
			return nil, fmt.Errorf("create workspace failed")
		}
	}
	_, err = os.Stat(phantomJS.options.WorkDir)
	if err == os.ErrNotExist {
		err = os.MkdirAll(phantomJS.options.WorkDir, 0755)
		if err == os.ErrPermission {
			return nil, fmt.Errorf("have no permission to create %s", phantomJS.options.WorkDir)
		} else {
			return nil, err
		}
	} else if err != nil {
		return nil, err
	}

	if phantomJS.options.LogFilePath == "" {
		phantomJS.logFile, err = ioutil.TempFile(phantomJS.options.WorkDir, "log")
		if err != nil {
			return nil, errors.New("open temp log file failed")
		}
		phantomJS.options.LogFilePath = phantomJS.logFile.Name()
	} else {
		phantomJS.logFile, err = os.OpenFile(phantomJS.options.LogFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			return nil, errors.New("open log file failed")
		}
	}

	if phantomJS.options.CookiesFile == "" {
		cookieFile, err := ioutil.TempFile(phantomJS.options.WorkDir, "cookie")
		if err != nil {
			return nil, errors.New("open temp cookie file failed")
		}
		phantomJS.options.CookiesFile = cookieFile.Name()
		cookieFile.Close()
	} else {
		info, err := os.Stat(phantomJS.options.CookiesFile)
		if err != nil && info.IsDir() {
			return nil, fmt.Errorf("can not access cookie file %s", phantomJS.options.CookiesFile)
		}
	}

	if port == 0 {
		port = FindFreePort()
	}
	phantomJS.addr = fmt.Sprintf("localhost:%d", port)
	phantomJS.options.Webdriver = phantomJS.addr

	err = phantomJS.start()
	if err != nil {
		return nil, err
	}

	return &phantomJS, nil
}
