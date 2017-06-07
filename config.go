package phantomjs

import (
	"fmt"
	"reflect"
	"strings"
)

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
	cap.ResourceTimeout = 60 * 1000 // default timeout in 60 seconds

	return &cap
}

type Options struct {
	// Sets the file name to store the persistent cookies
	CookiesFile string `json:"--cookies-file,omitempty"`
	// Specifies JSON-formatted configuration file
	Config string `json:"--config,omitempty"`
	// Prints additional warning and debug message: 'true' or 'false' (default)
	Debug bool `json:"--debug,omitempty" default:"false"`
	// Enables disk cache: 'true' or 'false' (default)
	DiskCache bool `json:"--disk-cache,omitempty" default:"false"`
	// Specifies the location for the disk cache
	DiskCachePath string `json:"--disk-cache-path,omitempty"`
	// Ignores SSL errors (expired/self-signed certificate errors): 'true' or 'false' (default)
	IgnoreSSLErrors bool `json:"--ignore-ssl-errors,omitempty" default:"false"`
	// Loads all inlined images: 'true' (default) or 'false'
	LoadImages bool `json:"--load-images" default:"true"`
	// Allows use of 'file:///' URLs: 'true' (default) or 'false'
	LocalURLAccess bool `json:"--local-url-access" default:"true"`
	// Specifies the location for local storage
	LocalStoragePath string `json:"--local-storage-path,omitempty"`
	// Sets the maximum size of the local storage (in KB)
	LocalStorageQuota string `json:"--local-storage-quota,omitempty"`
	// Specifies the location for offline storage
	OfflineStoragePath string `json:"--offline-storage-path,omitempty"`
	// set the maximum size of the local storage (in KB)
	OfflineStorageQuota int `json:"--offline-storage-quota,omitempty"`
	// Allows local content to access remote URL: 'true' or 'false' (default)
	LocalToRemoteURLAccess bool `json:"--local-to-remote-access,omitempty" default:"false"`
	// Limits the size of the disk cache (in KB)
	MaxDiskCacheSize int `json:"--max-disk-cache-size,omitempty"`
	// Sets the encoding for the terminal output, default is 'utf8'
	OutputEncoding string `json:"--output-encoding,omitempty" default:"utf8"`
	// Starts the script in a debug harness and listens on the specified port
	RemoteDebuggerPort int `json:"--remote-debugger-port,omitempty"`
	// Runs the script in the debugger immediately: 'true' or 'false' (default)
	RemoteDebuggerAutorun bool `json:"--remote-debugger-autorun,omitempty" default:"false"`
	// Specifies the proxy type, 'http' (default), 'none' (disable completely), or 'socks5'
	ProxyType string `json:"--proxy-type,omitempty" default:"http"`
	// Sets the proxy server, format: proxy.company.com:8080
	Proxy string `json:"--proxy,omitempty"`
	// Provides authentication information for the proxy, format: username:password
	ProxyAuth string `json:"--proxy-auth,omitempty"`
	// Sets the encoding used for the starting script, default is 'utf8'
	ScriptEncoding string `json:"--script-encoding,omitempty" default:"utf8"`
	// Sets the script language instead of detecting it: 'javascript'
	ScriptLanguage string `json:"--script-language,omitempty" default:"javascript"`
	// Enables web security, 'true' (default) or 'false'
	WebSecurity bool `json:"--web-security" default:"true"`
	// Selects a specific SSL protocol version to offer.
	// Values (case insensitive): TLSv1.2, TLSv1.1, TLSv1.0, TLSv1 (same as v1.0), SSLv3, or ANY.
	// Default is to offer all that Qt thinks are secure (SSLv3 and up).
	// Not all values may be supported, depending on the system OpenSSL library.
	SSLProtocol string `json:"--ssl-protocol,omitempty" default:"ANY"`
	// Sets supported TLS/SSL ciphers.
	// Argument is a colon-separated list of OpenSSL cipher names (macros like ALL, kRSA, etc. may not be used).
	// Default matches modern browsers
	SSLCiphers string `json:"--ssl-cipher,omitempty"`
	// Sets the location for custom CA certificates
	// If none set, uses environment variable SSL_CERT_DIR. If none set too, uses system default.
	SSLCertificatesPath string `json:"--ssl-certificates-path,omitempty"`
	// Sets the location of a client certificate
	SSLClientCertificateFile string `json:"--ssl-client-certificate-file,omitempty"`
	// Sets the location of a clients' private key
	SSLClientKeyFile string `json:"--ssl-client-key-file,omitempty"`
	// Sets the passphrase for the clients' private key
	SSLClientKeyPassphrase string `json:"--ssl-client-key-passphrase,omitempty"`
	// Starts in 'Remote WebDriver mode' (embedded GhostDriver): '[[<IP>:]<PORT>]' (default '127.0.0.1:8910')
	Webdriver string `json:"--webdriver" default:"127.0.0.1:8910"`
	// File where to write the WebDriver's Log (default 'none') (NOTE: needs '--webdriver')
	WebdriverLogFile string `json:"--webdriver-logfile,omitempty" default:"none"`
	// URL to the Selenium Grid HUB: 'URL_TO_HUB' (default 'none') (NOTE: needs '--webdriver')
	WebdriverSeleniumGridHub string `json:"--webdriver-selenium-grid-hub,omitempty" default:"none"`

	WorkDir     string `json:"-"`
	LogFilePath string `json:"-"`
	KeepAlive   bool   `json:"-"`
	IsRemote    bool   `json:"-"`
}

func (options *Options) Args() []string {
	ret := []string{}

	structType := reflect.TypeOf(*options)
	structVal := reflect.ValueOf(*options)

LABEL_NEXT_FIELD:
	for i := 0; i < structType.NumField(); i++ {
		typ := structType.Field(i)
		valField := structVal.Field(i)
		defaultVal := typ.Tag.Get("default")
		name := ""

		for _, tag := range strings.Split(typ.Tag.Get("json"), ",") {
			switch tag {
			case "-":
				break LABEL_NEXT_FIELD
			case "omitempty":
				continue
			default:
				name = tag
				break
			}
		}
		if name == "" {
			break
		}

		val := fmt.Sprintf("%v", valField.Interface())
		if valField.Kind() != reflect.String && val == "0" || val == "" || val == "false" || val == defaultVal {
			continue
		}
		ret = append(ret, fmt.Sprintf("%s=%s", name, val))
	}
	return ret
}

func DefaultOptions() *Options {
	options := Options{}

	options.LoadImages = true
	options.LocalURLAccess = true
	options.WebSecurity = true

	options.WorkDir = ""
	options.KeepAlive = true
	options.IsRemote = true

	return &options
}
