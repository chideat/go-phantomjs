package phantomjs

import (
	"encoding/json"
	"fmt"
)

type Element struct {
	Id      string `json:"ELEMENT"`
	Session *Session
}

func (element *Element) dispatchWithoutReturn(command Command) error {
	_, err := element.dispatchWithReturn(command)
	return err
}

func (element *Element) dispatchWithReturn(command Command) ([]byte, error) {
	args := map[string]interface{}{
		"sessionId": element.Session.Id,
		"id":        element.Id,
	}

	res, err := command.Execute(element.Session.addr, args, element.Session.options)
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

	return ret.Value, nil
}

func (element *Element) Click() error {
	return element.dispatchWithoutReturn(Commands.ClickElement)
}

func (element *Element) GetText() (string, error) {
	res, err := element.dispatchWithReturn(Commands.GetElementText)
	return string(res), err
}

func (element *Element) GetValue() (string, error) {
	res, err := element.dispatchWithReturn(Commands.GetElementValue)
	return string(res), err
}

func (element *Element) GetTagName() (string, error) {
	res, err := element.dispatchWithReturn(Commands.GetElementTagName)
	return string(res), err
}
