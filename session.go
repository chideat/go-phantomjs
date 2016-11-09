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

func (session *Session) GetAllCookies() ([]map[string]interface{}, error) {
	args := map[string]interface{}{
		"sessionId": session.Id,
	}

	res, err := Commands.Refresh.Execute(session.addr, args, session.options)
	if err != nil {
		return nil, err
	}
	if res.Code != ErrorCode_Success {
		return nil, errors.New(string(res.Data))
	}

	var ret struct {
		Response
		Value []map[string]interface{} `json:"value"`
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
	if ret.Status == 0 {
		return fmt.Errorf("%s", res.Data)
	}
	return nil
}

func (session *Session) DeleteAllCookies() error {
	// {"sessionId":"ccb17e40-a580-11e6-826c-4bb6080aa8c0","status":0,"value":{}}
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
	if ret.Status == 0 {
		return fmt.Errorf("%s", res.Data)
	}
	return nil
}

func (session *Session) FindElement() {
}

func (session *Session) FindElements() {
}

func (session *Session) FindChildElement() {
}

func (session *Session) FindChildElements() {
}

func (session *Session) ClearElement() {
}

func (session *Session) ClickElement() {
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

func (session *Session) GetWindowSize(width, height int) (int, int, error) {
	// {"width": 100, "windowHandle": "current", "sessionId": "ccb17e40-a580-11e6-826c-4bb6080aa8c0", "height": 100}
	// {"sessionId":"ccb17e40-a580-11e6-826c-4bb6080aa8c0","status":0,"value":{}}
	args := map[string]interface{}{
		"sessionId":    session.Id,
		"width":        width,
		"height":       height,
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
	// {"sessionId":"209aa290-a584-11e6-a29c-5d4dcda7337f","status":0,"value":{"x":0,"y":0}}
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
	// {"sessionId":"ccb17e40-a580-11e6-826c-4bb6080aa8c0","status":0,"value":"https://www.baidu.com/"}
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
	// {"sessionId":"21f71130-a577-11e6-b9eb-8f2f50092c66","status":0,"value":"<html></html>"}
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
	return ret.Value, nil
}

func (session *Session) GetTitle() (string, error) {
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
	return ret.Value, nil
}

func (session *Session) ExecuteScript(script string, _args []interface{}) error {
	// {"sessionId":"eab2af60-a57e-11e6-8678-73f6cd39aa02","status":0,"value":null}
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

func (session *Session) ExecuteAsyncScript(script string, args map[string]interface{}) {
}

func (session *Session) GetElementText() {
}

func (session *Session) GetElementValue() {
}

func (session *Session) GetElementTagValue() {
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
