package phantomjs

type Proxy struct {
	ProxyType          string `json:"proxyType,omitempty"`
	AuthDetect         string `json:"autoDetect,omitempty"`
	FtpProxy           string `json:"ftpProxy,omitempty"`
	HttpProxy          string `json:"httpProxy,omitempty"`
	ProxyAutoconfigUrl string `json:"proxyAutoconfigUrl,omitempty"`
	SslProxy           string `json:"sslProxy,omitempty"`
	NoProxy            string `json:"noProxy,omitempty"`
	SocksProxy         string `json:"socksProxy,omitempty"`
	SocksUsername      string `json:"socksUsername,omitempty"`
	SocksPassword      string `json:"socksPassword,omitempty"`
}
