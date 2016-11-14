package phantomjs

type Cookie struct {
	Name     string      `json:"name"`
	Value    string      `json:"value"`
	Path     string      `json:"path"`
	Domain   string      `json:"domain"`
	Expiry   interface{} `json:"expiry"`
	Secure   bool        `json:"secure"`
	HttpOnly bool        `json:"httponly"`
}

func (cookie *Cookie) Map() map[string]interface{} {
	ret := map[string]interface{}{}

	ret["name"] = cookie.Name
	ret["value"] = cookie.Value
	ret["path"] = cookie.Path
	ret["domain"] = cookie.Domain
	ret["expiry"] = cookie.Expiry
	ret["secure"] = cookie.Secure
	ret["httponly"] = cookie.HttpOnly

	return ret
}
