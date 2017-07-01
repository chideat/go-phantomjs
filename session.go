package phantomjs

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Session struct {
	Id      string
	addr    string
	options *Options
	cap     *Capabilities
}

func (session *Session) ID() string {
	return session.Id
}

func (session *Session) Get(urlStr string) error {
	args := map[string]interface{}{
		"url":       urlStr,
		"sessionId": session.Id,
	}

	res, err := Commands.Get.Execute(session.addr, args, session.options)
	if err != nil {
		return err
	}
	if res.Code != ErrorCode_Success {
		return errors.New(string(res.Data))
	}
	var ret Response
	err = json.Unmarshal(res.Data, &ret)
	if err != nil {
		return err
	}
	if ret.Status != 0 {
		return fmt.Errorf("%s", res.Data)
	}
	return nil
}

func (session *Session) GoBack() {
}

func (session *Session) GoForward() {
}

func (session *Session) Refresh() error {
	args := map[string]interface{}{
		"sessionId": session.Id,
	}

	res, err := Commands.Refresh.Execute(session.addr, args, session.options)
	if err != nil {
		return err
	}
	if res.Code != ErrorCode_Success {
		return errors.New(string(res.Data))
	}

	var ret Response
	err = json.Unmarshal(res.Data, &ret)
	if err != nil {
		return err
	}
	if ret.Status != 0 {
		return fmt.Errorf("%s", res.Data)
	}
	return nil
}

func (session *Session) AddCookie(cookie *Cookie) error {
	args := map[string]interface{}{
		"sessionId": session.Id,
		"cookie":    cookie,
	}

	res, err := Commands.AddCookie.Execute(session.addr, args, session.options)
	if err != nil {
		return err
	}
	if res.Code != ErrorCode_Success {
		return errors.New(string(res.Data))
	}

	var ret Response

	err = json.Unmarshal(res.Data, &ret)
	if err != nil {
		return err
	}
	if ret.Status != 0 {
		return fmt.Errorf("%s", res.Data)
	}
	return nil
}

func (session *Session) GetAllCookies() ([]*Cookie, error) {
	args := map[string]interface{}{
		"sessionId": session.Id,
	}

	res, err := Commands.GetAllCookies.Execute(session.addr, args, session.options)
	if err != nil {
		return nil, err
	}
	if res.Code != ErrorCode_Success {
		return nil, errors.New(string(res.Data))
	}

	var ret struct {
		Response
		Value []*Cookie `json:"value"`
	}

	err = json.Unmarshal(res.Data, &ret)
	if err != nil {
		var ret struct {
			Response
			Value interface{} `json:"value"`
		}
		err = json.Unmarshal(res.Data, &ret)
		if err != nil {
			return nil, err
		}

		if ret.Status != 0 {
			return nil, fmt.Errorf("%s", res.Data)
		}
		return []*Cookie{}, nil
	}

	if ret.Status != 0 {
		return nil, fmt.Errorf("%s", res.Data)
	}
	return ret.Value, nil
}

func (session *Session) GetCookie(name string) (*Cookie, error) {
	cookies, err := session.GetAllCookies()
	if err != nil {
		return nil, err
	}
	for _, cookie := range cookies {
		if cookie.Name == name {
			return cookie, nil
		}
	}
	return nil, nil
}

func (session *Session) DeleteCookie(name string) error {
	args := map[string]interface{}{
		"sessionId": session.Id,
		"name":      name,
	}

	res, err := Commands.DeleteCookie.Execute(session.addr, args, session.options)
	if err != nil {
		return err
	}

	var ret Response
	err = json.Unmarshal(res.Data, &ret)
	if err != nil {
		return err
	}
	if ret.Status != 0 {
		return fmt.Errorf("%s", res.Data)
	}
	return nil
}

func (session *Session) DeleteAllCookies() error {
	args := map[string]interface{}{
		"sessionId": session.Id,
	}

	res, err := Commands.DeleteAllCookies.Execute(session.addr, args, session.options)
	if err != nil {
		return err
	}

	var ret Response
	err = json.Unmarshal(res.Data, &ret)
	if err != nil {
		return err
	}
	if ret.Status != 0 {
		return fmt.Errorf("%s", res.Data)
	}
	return nil
}

func (session *Session) FindElement(using, value string) (*Element, error) {
	args := map[string]interface{}{
		"sessionId": session.Id,
		"using":     using,
		"value":     value,
	}

	res, err := Commands.FindElement.Execute(session.addr, args, session.options)
	if err != nil {
		return nil, err
	}

	var ret Response
	err = json.Unmarshal(res.Data, &ret)
	if err != nil {
		return nil, err
	}

	if ret.Status != 0 {
		return nil, fmt.Errorf("%s", res.Data)
	}

	var element *Element
	err = json.Unmarshal(ret.Value, &element)
	if err != nil {
		return nil, err
	}

	element.Session = session
	return element, nil
}

func (session *Session) FindElements() {
}

func (session *Session) FindChildElement() {
}

func (session *Session) FindChildElements() {
}

func (session *Session) ClearElement() {
}

func (session *Session) GetCurrentWindowHandle() (string, error) {
	args := map[string]interface{}{
		"sessionId": session.Id,
	}

	res, err := Commands.GetCurrentWindowHandle.Execute(session.addr, args, session.options)
	if err != nil {
		return "", err
	}

	var ret struct {
		Response
		Value string `json:"value"`
	}

	err = json.Unmarshal(res.Data, &ret)
	if err != nil {
		return "", err
	}
	if ret.Status != 0 {
		return "", fmt.Errorf("%s", res.Data)
	}

	return ret.Value, nil
}

func (session *Session) GetWindowHandles() ([]string, error) {
	args := map[string]interface{}{
		"sessionId": session.Id,
	}

	res, err := Commands.GetCurrentWindowHandle.Execute(session.addr, args, session.options)
	if err != nil {
		return nil, err
	}

	var ret struct {
		Response
		Value []string `json:"value"`
	}

	err = json.Unmarshal(res.Data, &ret)
	if err != nil {
		return nil, err
	}
	if ret.Status != 0 {
		return nil, fmt.Errorf("%s", res.Data)
	}

	return ret.Value, nil
}

func (session *Session) GetWindowSize() (int, int, error) {
	args := map[string]interface{}{
		"sessionId":    session.Id,
		"windowHandle": "current",
	}

	res, err := Commands.GetWindowSize.Execute(session.addr, args, session.options)
	if err != nil {
		return 0, 0, err
	}

	var ret struct {
		Response
		Value struct {
			Width  int `json:"width"`
			Height int `json:"height"`
		} `json:"value"`
	}

	err = json.Unmarshal(res.Data, &ret)
	if err != nil {
		return 0, 0, err
	}
	if ret.Status != 0 {
		return 0, 0, fmt.Errorf("%s", res.Data)
	}
	return ret.Value.Width, ret.Value.Height, nil
}

func (session *Session) SetWindowSize(width, height int) error {
	args := map[string]interface{}{
		"sessionId":    session.Id,
		"width":        width,
		"height":       height,
		"windowHandle": "current",
	}

	res, err := Commands.SetWindowSize.Execute(session.addr, args, session.options)
	if err != nil {
		return err
	}

	var ret Response
	err = json.Unmarshal(res.Data, &ret)
	if err != nil {
		return err
	}
	if ret.Status != 0 {
		return fmt.Errorf("%s", res.Data)
	}
	return nil
}

func (session *Session) GetWindowPosition() (int, int, error) {
	args := map[string]interface{}{
		"sessionId":    session.Id,
		"windowHandle": "current",
	}

	res, err := Commands.GetWindowPosition.Execute(session.addr, args, session.options)
	if err != nil {
		return 0, 0, err
	}

	var ret struct {
		Response
		Value struct {
			X int `json:"x"`
			Y int `json:"y"`
		} `json:"value"`
	}
	err = json.Unmarshal(res.Data, &ret)
	if err != nil {
		return 0, 0, err
	}
	if ret.Status != 0 {
		return 0, 0, fmt.Errorf("%s", res.Data)
	}
	return ret.Value.X, ret.Value.Y, nil
}

func (session *Session) SetWindowPosition(x, y int) error {
	args := map[string]interface{}{
		"sessionId":    session.Id,
		"windowHandle": "current",
		"x":            x,
		"y":            y,
	}

	res, err := Commands.SetWindowPosition.Execute(session.addr, args, session.options)
	if err != nil {
		return err
	}

	var ret Response
	err = json.Unmarshal(res.Data, &ret)
	if err != nil {
		return err
	}
	if ret.Status != 0 {
		return fmt.Errorf("%s", res.Data)
	}
	return nil
}

func (session *Session) GetCurrentUrl() (string, error) {
	args := map[string]interface{}{
		"sessionId": session.Id,
	}

	res, err := Commands.GetCurrentUrl.Execute(session.addr, args, session.options)
	if err != nil {
		return "", err
	}

	var ret struct {
		Response
		Value string `json:"value"`
	}
	err = json.Unmarshal(res.Data, &ret)
	if err != nil {
		return "", err
	}
	if ret.Status != 0 {
		return "", fmt.Errorf("%s", res.Data)
	}
	return ret.Value, nil
}

func (session *Session) GetPageSource() (string, error) {
	args := map[string]interface{}{
		"sessionId": session.Id,
	}

	res, err := Commands.GetPageSource.Execute(session.addr, args, session.options)
	if err != nil {
		return "", err
	}

	var ret struct {
		Response
		Value string `json:"value"`
	}
	err = json.Unmarshal(res.Data, &ret)
	if err != nil {
		return "", err
	}
	if ret.Status != 0 {
		return "", fmt.Errorf("%s", res.Data)
	}
	if ret.Value == "<html><head></head><body></body></html>" {
		return "", fmt.Errorf("Session.GetPageSource: get empty page source")
	}
	return ret.Value, nil
}

func (session *Session) GetTitle() (string, error) {
	args := map[string]interface{}{
		"sessionId": session.Id,
	}

	res, err := Commands.GetTitle.Execute(session.addr, args, session.options)
	if err != nil {
		return "", err
	}

	var ret struct {
		Response
		Value string `json:"value"`
	}
	err = json.Unmarshal(res.Data, &ret)
	if err != nil {
		return "", err
	}
	if ret.Status != 0 {
		return "", fmt.Errorf("%s", res.Data)
	}
	return ret.Value, nil
}

func (session *Session) ExecuteScript(script string, _args []interface{}) error {
	args := map[string]interface{}{
		"sessionId": session.Id,
		"script":    script,
		"args":      _args,
	}

	res, err := Commands.ExecuteScript.Execute(session.addr, args, session.options)
	if err != nil {
		return err
	}

	var ret Response
	err = json.Unmarshal(res.Data, &ret)
	if err != nil {
		return err
	}
	if ret.Status != 0 {
		return fmt.Errorf("%s", res.Data)
	}
	return nil
}

// not known how to use
func (session *Session) ExecuteAsyncScript(script string, _args map[string]interface{}) error {
	args := map[string]interface{}{
		"sessionId": session.Id,
		"script":    script,
		"args":      _args,
	}

	res, err := Commands.ExecuteAsyncScript.Execute(session.addr, args, session.options)
	if err != nil {
		return err
	}

	var ret Response
	err = json.Unmarshal(res.Data, &ret)
	if err != nil {
		return err
	}
	if ret.Status != 0 {
		return fmt.Errorf("%s", res.Data)
	}
	return nil
}

func (session *Session) IsElementSelected() {
}

func (session *Session) IsElementEnabled() {
}

func (session *Session) IsElementDisplayed() {
}

func (session *Session) GetElementLocation() {
}

func (session *Session) GetElementLocationOnceScrolledIntoView() {
}

func (session *Session) GetElementSize() {
}

func (session *Session) GetElementRect() {
}

func (session *Session) GetElementAttribute() {
}

func (session *Session) SetPageLoadTimeouts(ms int64) error {
	args := map[string]interface{}{
		"sessionId": session.Id,
		"ms":        ms,
		"type":      "page load",
	}

	res, err := Commands.SetTimeouts.Execute(session.addr, args, session.options)
	if err != nil {
		return err
	}

	var ret Response
	err = json.Unmarshal(res.Data, &ret)
	if err != nil {
		return err
	}
	if ret.Status != 0 {
		return fmt.Errorf("%s", res.Data)
	}
	return nil
}

func (session *Session) SetScriptTimeout(ms int64) error {
	args := map[string]interface{}{
		"sessionId": session.Id,
		"ms":        ms,
	}

	res, err := Commands.SetScriptTimeout.Execute(session.addr, args, session.options)
	if err != nil {
		return err
	}

	var ret Response
	err = json.Unmarshal(res.Data, &ret)
	if err != nil {
		return err
	}
	if ret.Status != 0 {
		return fmt.Errorf("%s", res.Data)
	}
	return nil
}

func (session *Session) Close() error {
	args := map[string]interface{}{
		"sessionId": session.Id,
	}

	res, err := Commands.Quit.Execute(session.addr, args, session.options)
	if err != nil {
		return err
	}

	var ret Response
	err = json.Unmarshal(res.Data, &ret)
	if err != nil {
		return err
	}
	if ret.Status != 0 {
		return fmt.Errorf("%s", res.Data)
	}
	return nil
}

func NewSession(addr, id string, cap *Capabilities, options *Options) *Session {
	session := Session{}
	session.addr = addr
	session.Id = id
	session.cap = cap
	session.options = options

	return &session
}
