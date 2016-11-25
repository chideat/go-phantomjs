package phantomjs

import "testing"

func TestNewPhantomJSWithExecutePath(t *testing.T) {
	options := &Options{}
	phantomJS, err := NewPhantomJSWithExecutePath("/usr/local/bin/phantomjs", 0, options)
	if err != nil {
		t.Fatal(err)
	}
	err = phantomJS.Quit()
	if err != nil {
		t.Fatal(err)
	}
}

func TestNewPhantomJS(t *testing.T) {
	options := &Options{}
	phantomJS, err := NewPhantomJS(0, options)
	if err != nil {
		t.Fatal(err)
	}
	err = phantomJS.Quit()
	if err != nil {
		t.Fatal(err)
	}
}

func TestQuit(t *testing.T) {
	options := &Options{}
	phantomJS, err := NewPhantomJS(0, options)
	if err != nil {
		t.Fatal(err)
	}
	err = phantomJS.Quit()
	if err != nil {
		t.Fatal(err)
	}
}

func TestNewSession(t *testing.T) {
	options := &Options{}
	phantomJS, err := NewPhantomJS(0, options)
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		err = phantomJS.Quit()
		if err != nil {
			t.Fatal(err)
		}
	}()

	session, err := phantomJS.NewSession(nil)
	if err != nil {
		t.Fatal(err)
	}
	err = session.Close()
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetAllSessions(t *testing.T) {
	options := &Options{}
	phantomJS, err := NewPhantomJS(0, options)
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		err := phantomJS.Quit()
		if err != nil {
			t.Fatal(err)
		}
	}()

	session, err := phantomJS.NewSession(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		err = session.Close()
		if err != nil {
			t.Fatal(err)
		}
	}()

	sessions, err := phantomJS.GetAllSessions()
	if err != nil {
		t.Fatal(err)
	}
	if len(sessions) != 1 {
		t.Fatal("get all sessions failed")
	}
}
