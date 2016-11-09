package phantomjs

type Capabilities struct {
	BrowserName              string `json:"browserName"`
	DriverName               string `json:"driverName"`
	Version                  string `json:"version"`
	Platform                 string `json:"platform"`
	JavscriptEnabled         bool   `json:"javascriptEnabled"`
	TakesScreenshot          bool   `json:"takesScreenshot"`
	HandlesAlerts            bool   `json:"handlesAlerts"`
	DatabaseEnabled          bool   `json:"databaseEnabled"`
	LocationContextEnabled   bool   `json:"locationContextEnabled"`
	ApplicationCacheEnabled  bool   `json:"applicationCacheEnabled"`
	BrowserConnectionEnabled bool   `json:"browserConnectionEnabled"`
	CssSelectorsEnabled      bool   `json:"cssSelectorsEnabled"`
	WebStorageEnabled        bool   `json:"webStorageEnabled"`
	Rotatable                bool   `json:"rotatable"`
	AcceptSslCerts           bool   `json:"acceptSslCerts"`
	NativeEvents             bool   `json:"nativeEvents"`
	Proxy                    Proxy  `json:"proxy"`
}

func DesiredCapabilities() *Capabilities {
	cap := Capabilities{}
	cap.BrowserName = "phantomjs"
	cap.Version = ""
	cap.Platform = ""
	cap.JavscriptEnabled = true
	cap.TakesScreenshot = true
	cap.HandlesAlerts = false
	cap.DatabaseEnabled = false
	cap.LocationContextEnabled = false
	cap.ApplicationCacheEnabled = false
	cap.BrowserConnectionEnabled = false
	cap.CssSelectorsEnabled = true
	cap.WebStorageEnabled = false
	cap.Rotatable = false
	cap.AcceptSslCerts = false
	cap.NativeEvents = true
	cap.Proxy.ProxyType = "DIRECT"

	return &cap
}

type Options struct {
	WorkDir         string
	LogFilePath     string
	CookiesFilePath string

	KeepAlive bool
	IsRemote  bool
	UserAgent string
}

func DefaultOptions() *Options {
	options := Options{}

	options.WorkDir = ""
	options.KeepAlive = true
	options.IsRemote = true
	options.UserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/54.0.2840.71 Safari/537.36"

	return &options
}
