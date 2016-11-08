package phantomjs

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

var client http.Client

func init() {
	client = http.Client{Transport: &http.Transport{}}
}

type _Commands struct {
	Status Command
	NewSession
	GetallSessions
	DeleteSession
	Close
	Quit
	Get
	GoBack
	GoForward
	Refresh
	AddCookie
	GetCookie
	GetAllCookies
	DeleteCookie
	DeleteAllCookies
	FindElement
	FindElements
	FindChildElement
	FindChildElements
	ClearElement
	ClickElement
	SendKeysToElement
	SendKeysToActiveElement
	SubmitElement
	UploadFile
	GetCurrentWindowHandle
	GetWindowHandles
	GetWindowSize
	W3CGetWindowSize
	GetWindowPosition
	SetWindowSize
	W3CSetWindowSize
	SetWindowPosition
	SwitchToWindow
	SwitchToFrame
	SwitchToParentFrame
	GetActiveElement
	GetCurrenturl
	GetPageSource
	GetTitle
	ExecuteScript
	GetElementText
	GetElementValue
	GetElementTagName
	SetElementSelected
	IsElementSelected
	IsElementEnabled
	IsElementDisplayed
	GetElementLocation
	GetElementLocationOnceScrolledIntoView
	GetElementSize
	GetElementRect
	GetElementAttribute
	GetElementValueOfCssProperty
	ElementEquals
	Screenshot
	ElementScreenshot
	ImplicitlyWait
	ExecuteAsyncScript
	SetScriptTimeout
	SetTimeouts
	WindowMaximize
	W3cMaximizeWindow
	GetLog
	GetAvailableLogTypes
	DismissAlert
	AcceptAlert
	SetAlertValue
	GetAlertText
	SetAlertCredentials
	MouseClick
	MouseDoubleClick
	MouseButtonDown
	MouseButtonUp
	MouseMoveTo
	SetScreenOrientation
	GetScreenOrientation
	TouchSingleTap
	TouchDown
	TouchUp
	TouchMove
	TouchScroll
	TouchDoubleTap
	TouchLongPress
	TouchFlick
	ExecuteSql
	GetLocation
	SetLocation
	GetAppCache
	GetAppCacheStatus
	ClearAppCache
	GetLocalStorageItem
	RemoveLocalStorageItem
	GetLocalStorageKeys
	SetLocalStorageItem
	ClearLocalStorage
	GetLocalStorageSize
	GetSessionStorageItem
	RemoveSessionStorageItem
	GetSessionStorageKeys
	SetSessionStorageItem
	ClearSessionStorage
	GetSessionStorageSize
	GetNetworkConnection
	SetNetworkConnection
	GetCurrentContextHandle
	GetContextHandles
	SwitchToContext
}

var Commands = _Commands{
	Status:                                 Command{Cmd: "status", Method: "GET", Route: "/status"},
	NewSession:                             Command{Cmd: "newSession", Method: "POST", Route: "/session"},
	GetAllSessions:                         Command{Cmd: "getAllSessions", Method: "GET", Route: "/sessions"},
	DeleteSession:                          Command{Cmd: "deleteSession", Method: "DELETE", Route: "/session/:session_id"},
	Close:                                  Command{Cmd: "close", Method: "POST", Route: ""},
	Quit:                                   Command{Cmd: "quit", Method: "DELETE", Route: "/session/:session_id"},
	Get:                                    Command{Cmd: "get", Method: "GET", Route: "/session/:session_id/url"},
	GoBack:                                 Command{Cmd: "goBack", Method: "POST", Route: "/session/:session_id/back"},
	GoForward:                              Command{Cmd: "goForward", Method: "POST", Route: "/session/:session_id/forward"},
	Refresh:                                Command{Cmd: "refresh", Method: "POST", Route: "/session/:session_id/refresh"},
	AddCookie:                              Command{Cmd: "addCookie", Method: "POST", Route: "/session/:session_id/cookie"},
	GetCookie:                              Command{Cmd: "getCookie", Method: "GET", Route: ""},
	GetAllCookies:                          Command{Cmd: "getCookies", Method: "GET", Route: "/session/:session_id/cookie"},
	DeleteCookie:                           Command{Cmd: "deleteCookie", Method: "DELETE", Route: "/session/:session_id/cookie/:name"},
	DeleteAllCookies:                       Command{Cmd: "deleteAllCookies", Method: "DELETE", Route: "/session/:session_id/cookie"},
	FindElement:                            Command{Cmd: "findElement", Method: "POST", Route: "/session/:session_id/element"},
	FindElements:                           Command{Cmd: "findElements", Method: "POST", Route: "/session/:session_id/elements"},
	FindChildElement:                       Command{Cmd: "findChildElement", Method: "POST", Route: "/session/:session_id/element/:id/element"},
	FindChildElements:                      Command{Cmd: "findChildElements", Method: "POST", Route: "/session/:session_id/element/:id/elements"},
	ClearElement:                           Command{Cmd: "clearElement", Method: "POST", Route: "/session/:session_id/element/:id/clear"},
	ClickElement:                           Command{Cmd: "clickElement", Method: "POST", Route: "/session/:session_id/element/:id/click"},
	SendKeysToElement:                      Command{Cmd: "sendKeysToElement", Method: "POST", Route: ""},
	SendKeysToActiveElement:                Command{Cmd: "sendKeysToActiveElement", Method: "POST", Route: ""},
	SubmitElement:                          Command{Cmd: "submitElement", Method: "POST", Route: ""},
	UploadFile:                             Command{Cmd: "uploadFile", Method: "POST", Route: ""},
	GetCurrentWindowHandle:                 Command{Cmd: "getCurrentWindowHandle", Method: "GET", Route: "/session/:session_id/window_handle"},
	GetWindowHandles:                       Command{Cmd: "getWindowHandles", Method: "GET", Route: "/session/:session_id/window_handles"},
	GetWindowSize:                          Command{Cmd: "getWindowSize", Method: "GET", Route: "/session/:session_id/window/:window_handle/size"},
	W3CGetWindowSize:                       Command{Cmd: "w3cGetWindowSize", Method: "GET", Route: "/session/:session_id/window/size"},
	SetWindowSize:                          Command{Cmd: "setWindowSize", Method: "POST", Route: "/session/:session_id/window/:window_handle/size"},
	W3CSetWindowSize:                       Command{Cmd: "w3cSetWindowSize", Method: "POST", Route: "/session/:session_id/window/size"},
	GetWindowPosition:                      Command{Cmd: "getWindowPosition", Method: "GET", Route: "/session/:session_id/window/:window_handle/position"},
	SetWindowPosition:                      Command{Cmd: "setWindowPosition", Method: "POST", Route: "/session/:session_id/window/:window_handle/position"},
	SwitchToWindow:                         Command{Cmd: "switchToWindow", Method: "POST", Route: ""},
	SwitchToFrame:                          Command{Cmd: "switchToFrame", Method: "POST", Route: ""},
	SwitchToParentFrame:                    Command{Cmd: "switchToParentFrame", Method: "POST", Route: ""},
	GetActiveElement:                       Command{Cmd: "getActiveElement", Method: "GET", Route: ""},
	GetCurrentUrl:                          Command{Cmd: "getCurrentUrl", Method: "GET", Route: "/session/:session_id/url"},
	GetPageSource:                          Command{Cmd: "getPageSource", Method: "GET", Route: "/session/:session_id/source"},
	GetTitle:                               Command{Cmd: "getTitle", Method: "GET", Route: "/session/:session_id/title"},
	ExecuteScript:                          Command{Cmd: "executeScript", Method: "POST", Route: "/session/:session_id/execute"},
	GetElementText:                         Command{Cmd: "getElementText", Method: "GET", Route: "/session/:session_id/element/:id/text"},
	GetElementValue:                        Command{Cmd: "getElementValue", Method: "GET", Route: "/session/:session_id/element/:id/value"},
	GetElementTagName:                      Command{Cmd: "getElementTagName", Method: "GET", Route: "/session/:session_id/element/:id/name"},
	SetElementSelected:                     Command{Cmd: "setElementSelected", Method: "POST", Route: ""},
	IsElementSelected:                      Command{Cmd: "isElementSelected", Method: "GET", Route: "/session/:session_id/element/:id/selected"},
	IsElementEnabled:                       Command{Cmd: "isElementEnabled", Method: "GET", Route: "/session/:session_id/element/:id/enabled"},
	IsElementDisplayed:                     Command{Cmd: "isElementDisplayed", Method: "GET", Route: "/session/:session_id/element/:id/displayed"},
	GetElementLocation:                     Command{Cmd: "getElementLocation", Method: "GET", Route: "/session/:session_id/element/:id/location"},
	GetElementLocationOnceScrolledIntoView: Command{Cmd: "getElementLocationOnceScrolledIntoView", Method: "GET", Route: "/session/:session_id/element/:id/location_in_view"},
	GetElementSize:                         Command{Cmd: "getElementSize", Method: "GET", Route: "/session/:session_id/element/:id/size"},
	GetElementRect:                         Command{Cmd: "getElementRect", Method: "GET", Route: "/session/:session_id/element/:id/rect"},
	GetElementAttribute:                    Command{Cmd: "getElementAttribute", Method: "GET", Route: "/session/:session_id/element/:id/attribute/:name"},
	GetElementValueOfCssProperty:           Command{Cmd: "getElementValueOfCssProperty", Method: "GET", Route: ""},
	ElementEquals:                          Command{Cmd: "elementEquals", Method: "POST", Route: ""},
	Screenshot:                             Command{Cmd: "screenshot", Method: "POST", Route: ""},
	ElementScreenshot:                      Command{Cmd: "elementScreenshot", Method: "POST", Route: ""},
	ImplicitlyWait:                         Command{Cmd: "implicitlyWait", Method: "POST", Route: ""},
	ExecuteAsyncScript:                     Command{Cmd: "executeAsyncScript", Method: "POST", Route: "/session/:session_id/execute_async"},
	SetScriptTimeout:                       Command{Cmd: "setScriptTimeout", Method: "POST", Route: "/session/:session_id/timeouts/async_script"},
	SetTimeouts:                            Command{Cmd: "setTimeouts", Method: "POST", Route: "/session/:session_id/timeouts"},
	WindowMaximize:                         Command{Cmd: "windowMaximize", Method: "POST", Route: ""},
	W3cMaximizeWindow:                      Command{Cmd: "w3cMaximizeWindow", Method: "POST", Route: ""},
	GetLog:                                 Command{Cmd: "getLog", Method: "GET", Route: ""},
	GetAvailableLogTypes:                   Command{Cmd: "getAvailableLogTypes", Method: "GET", Route: ""},
	DismissAlert:                           Command{Cmd: "dismissAlert", Method: "POST", Route: ""},
	AcceptAlert:                            Command{Cmd: "acceptAlert", Method: "POST", Route: ""},
	SetAlertValue:                          Command{Cmd: "setAlertValue", Method: "POST", Route: ""},
	GetAlertText:                           Command{Cmd: "getAlertText", Method: "GET", Route: ""},
	SetAlertCredentials:                    Command{Cmd: "setAlertCredentials", Method: "POST", Route: ""},
	MouseClick:                             Command{Cmd: "mouseClick", Method: "POST", Route: ""},
	MouseDoubleClick:                       Command{Cmd: "mouseDoubleClick", Method: "POST", Route: ""},
	MouseButtonDown:                        Command{Cmd: "mouseButtonDown", Method: "POST", Route: ""},
	MouseButtonUp:                          Command{Cmd: "mouseButtonUp", Method: "POST", Route: ""},
	MouseMoveTo:                            Command{Cmd: "mouseMoveTo", Method: "POST", Route: ""},
	SetScreenOrientation:                   Command{Cmd: "setScreenOrientation", Method: "POST", Route: ""},
	GetScreenOrientation:                   Command{Cmd: "getScreenOrientation", Method: "GET", Route: ""},
	TouchSingleTap:                         Command{Cmd: "touchSingleTap", Method: "POST", Route: ""},
	TouchDown:                              Command{Cmd: "touchDown", Method: "POST", Route: ""},
	TouchUp:                                Command{Cmd: "touchUp", Method: "POST", Route: ""},
	TouchMove:                              Command{Cmd: "touchMove", Method: "POST", Route: ""},
	TouchScroll:                            Command{Cmd: "touchScroll", Method: "POST", Route: ""},
	TouchDoubleTap:                         Command{Cmd: "touchDoubleTap", Method: "POST", Route: ""},
	TouchLongPress:                         Command{Cmd: "touchLongPress", Method: "POST", Route: ""},
	TouchFlick:                             Command{Cmd: "touchFlick", Method: "POST", Route: ""},
	ExecuteSql:                             Command{Cmd: "executeSql", Method: "POST", Route: ""},
	GetLocation:                            Command{Cmd: "getLocation", Method: "GET", Route: ""},
	SetLocation:                            Command{Cmd: "setLocation", Method: "POST", Route: ""},
	GetAppCache:                            Command{Cmd: "getAppCache", Method: "GET", Route: ""},
	GetAppCacheStatus:                      Command{Cmd: "getAppCacheStatus", Method: "GET", Route: ""},
	ClearAppCache:                          Command{Cmd: "clearAppCache", Method: "POST", Route: ""},
	GetLocalStorageItem:                    Command{Cmd: "getLocalStorageItem", Method: "GET", Route: ""},
	RemoveLocalStorageItem:                 Command{Cmd: "removeLocalStorageItem", Method: "POST", Route: ""},
	GetLocalStorageKeys:                    Command{Cmd: "getLocalStorageKeys", Method: "GET", Route: ""},
	SetLocalStorageItem:                    Command{Cmd: "setLocalStorageItem", Method: "POST", Route: ""},
	ClearLocalStorage:                      Command{Cmd: "clearLocalStorage", Method: "POST", Route: ""},
	GetLocalStorageSize:                    Command{Cmd: "getLocalStorageSize", Method: "GET", Route: ""},
	GetSessionStorageItem:                  Command{Cmd: "getSessionStorageItem", Method: "GET", Route: ""},
	RemoveSessionStorageItem:               Command{Cmd: "removeSessionStorageItem", Method: "POST", Route: ""},
	GetSessionStorageKeys:                  Command{Cmd: "getSessionStorageKeys", Method: "GET", Route: ""},
	SetSessionStorageItem:                  Command{Cmd: "setSessionStorageItem", Method: "POST", Route: ""},
	ClearSessionStorage:                    Command{Cmd: "clearSessionStorage", Method: "POST", Route: ""},
	GetSessionStorageSize:                  Command{Cmd: "getSessionStorageSize", Method: "GET", Route: ""},
	GetNetworkConnection:                   Command{Cmd: "getNetworkConnection", Method: "GET", Route: ""},
	SetNetworkConnection:                   Command{Cmd: "setNetworkConnection", Method: "POST", Route: ""},
	GetCurrentContextHandle:                Command{Cmd: "getCurrentContextHandle", Method: "GET", Route: ""},
	GetContextHandles:                      Command{Cmd: "getContextHandles", Method: "GET", Route: ""},
	SwitchToContext:                        Command{Cmd: "switchToContext", Method: "POST", Route: ""},
}

type Result struct {
	Code ErrorCode
	Data []byte
}

type Command struct {
	Cmd    string
	Method string
	Route  string
}

func (cmd *Command) Execute(server string, params map[string]string, option *ConnectionOption) (*Result, error) {
	if cmd.Route == "" {
		return nil, fmt.Errorf("cmd %s is NOT supported", cmd.Cmd)
	}

	route, err := RouteTemplate(cmd.Route, params)
	if err != nil {
		return nil, err
	}

	urlStr := fmt.Sprintf("%s%s", server, route)

	body, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	return cmd.http_request(cmd.Method, urlStr, body, option)
}

func (cmd *Command) http_request(method, urlStr string, body []byte, option *ConnectionOption) (*Result, error) {
	u, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	var (
		req *http.Request
	)

	if method != "POST" && method != "PUT" {
		req, err = http.NewRequest(method, urlStr, nil)
	} else {
		req, err = http.NewRequest(method, urlStr, bytes.NewBuffer(body))
	}
	if err != nil {
		return nil, err
	}
	req.Header.Set(method, u.Path)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-type", "application/json;charset=\"UTF-8\"")
	req.Header.Set("User-Agent", "Golang HTTP Client")

	if option.KeepAlive {
		req.Header.Set("Connection", "keep-alive")

		if u.User.Username() != "" {
			name := u.User.Username()
			passwd, _ := u.User.Password()
			auth := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", name, passwd)))
			req.Header.Set("Authorization", fmt.Sprintf("Basic %s", auth))
		}
	} else {
		// TODO add support for password manager
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	dataRaw, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, err
	}

	if 300 <= resp.StatusCode && resp.StatusCode < 304 {
		return cmd.http_request("GET", resp.Header.Get("location"), nil, option)
	}
	if 399 < resp.StatusCode && resp.StatusCode < 500 {
		return &Result{Code: ErrorCode(resp.StatusCode), Data: dataRaw}, nil
	}
	return &Result{Code: ErrorCode_Success, Data: dataRaw}, nil
}
