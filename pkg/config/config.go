package config

// Server defines the general server configuration.
type Server struct {
	Addr       string
	Kubeconfig string
}

// Logs defines the level and color for log configuration.
type Logs struct {
	Level  string
	Pretty bool
}

// Config is a combination of all available configurations.
type Config struct {
	Server Server
	Logs   Logs
}

// Load initializes a default configuration struct.
func Load() *Config {
	return &Config{}
}
