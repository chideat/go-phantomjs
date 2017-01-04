package phantomjs

type Capabilities struct {
	BrowserName              string `json:"browserName"`
	DriverName               string `json:"driverName"`
	Version                  string `json:"version"`
	Platform                 string `json:"platform"`
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
	// page.settings
	JavscriptEnabled              bool   `json:"phantomjs.page.settings.javascriptEnabled"`
	LoadImages                    bool   `json:"phantomjs.page.settings.loadImages"`
	LocalToRemoteUrlAccessEnabled bool   `json:"phantomjs.page.settings.localToRemoteUrlAccessEnabled"`
	UserAgent                     string `json:"phantomjs.page.settings.userAgent,omitempty,omitempty"`
	UserName                      string `json:"phantomjs.page.settings.userName,omitempty,omitempty"`
	Password                      string `json:"phantomjs.page.settings.password,omitempty,omitempty"`
	XSSAuditingEnabled            bool   `json:"phantomjs.page.settings.XSSAuditingEnabled"`
	WebSecurityEnabled            bool   `json:"phantomjs.page.settings.webSecurityEnabled"`
	ResourceTimeout               int64  `json:"phantomjs.page.settings.resourceTimeout,omitempty"` // milli-secs
	// page.customHeaders
	Headers map[string]string `json:"phantomjs.page.customHeaders,omitempty"`

	Proxy Proxy `json:"proxy"`
}

func DesiredCapabilities() *Capabilities {
	cap := Capabilities{}
	cap.BrowserName = "phantomjs"
	cap.Version = ""
	cap.Platform = ""
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
	// page.settings
	cap.JavscriptEnabled = true
	cap.LoadImages = false
	cap.UserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/54.0.2840.71 Safari/537.36"
	cap.ResourceTimeout = 2 * 60 * 1000 // default timeout in 2 minutes

	return &cap
}

type Options struct {
	WorkDir         string
	LogFilePath     string
	CookiesFilePath string

	KeepAlive bool
	IsRemote  bool
}

func DefaultOptions() *Options {
	options := Options{}

	options.WorkDir = ""
	options.KeepAlive = true
	options.IsRemote = true

	return &options
}
