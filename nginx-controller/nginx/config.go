package nginx

// Config holds NGINX configuration parameters
type Config struct {
	ProxyConnectTimeout           string
	ProxyReadTimeout              string
	ClientMaxBodySize             string
	HTTP2                         bool
	MainServerNamesHashBucketSize string
	MainServerNamesHashMaxSize    string
	MainLogFormat                 string
	ProxyBuffering                bool
	ProxyBuffers                  string
	ProxyBufferSize               string
	ProxyMaxTempFileSize          string
}

// NewDefaultConfig creates a Config with default values
func NewDefaultConfig() *Config {
	return &Config{
		ProxyConnectTimeout:        "60s",
		ProxyReadTimeout:           "60s",
		ClientMaxBodySize:          "1m",
		MainServerNamesHashMaxSize: "512",
		ProxyBuffering:             true,
	}
}
