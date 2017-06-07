package phantomjs

type ProxyType string

const (
	ProxyType_Direct     ProxyType = "direct"
	ProxyType_Pac                  = "pac"
	ProxyType_AutoDetect           = "autodetect"
	ProxyType_System               = "system"
	ProxyType_Manual               = "manual"
)

type Proxy struct {
	ProxyType ProxyType `json:"proxyType,omitempty"`
	// defines the url for a proxy auto-config file if proxyType is equal to "pac"
	ProxyAutoconfigUrl string `json:"proxyAutoconfigUrl,omitempty"`
	FTPProxy           string `json:"ftpProxy,omitempty"`
	HTTPProxy          string `json:"httpProxy,omitempty"`
	SSLProxy           string `json:"sslProxy,omitempty"`
	SocksProxy         string `json:"socksProxy,omitempty"`
	SocksVersion       int    `json:"socksVersion,omitempty"`
	SocksUsername      string `json:"socksUsername,omitempty"`
	SocksPassword      string `json:"socksPassword,omitempty"`
}

func NewHTTPProxy(addr string) *Proxy {
	proxy := Proxy{}
	proxy.ProxyType = ProxyType_Manual
	proxy.HTTPProxy = addr

	return &proxy
}

func NewSSLProxy(addr string) *Proxy {
	proxy := Proxy{}
	proxy.ProxyType = ProxyType_Manual
	proxy.SSLProxy = addr

	return &proxy
}

func NewSocks4Proxy(addr string) *Proxy {
	proxy := Proxy{}
	proxy.ProxyType = ProxyType_Manual
	proxy.SocksProxy = addr
	proxy.SocksVersion = 4

	return &proxy
}

func NewSocks5Proxy(addr string) *Proxy {
	proxy := Proxy{}
	proxy.ProxyType = ProxyType_Manual
	proxy.SocksProxy = addr
	proxy.SocksVersion = 5

	return &proxy
}
