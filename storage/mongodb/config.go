package mongodb

// Config defines the config for storage.
type Config struct {
	// Connection string to use for DB. Will override all other authentication values if used
	//
	// Optional. Default is ""
	ConnectionURI string

	// Host name where the DB is hosted
	//
	// Optional. Default is "127.0.0.1"
	Host string

	// Port where the DB is listening on
	//
	// Optional. Default is 27017
	Port int

	// Server username
	//
	// Optional. Default is ""
	Username string

	// Server password
	//
	// Optional. Default is ""
	Password string

	// Database name
	//
	// Optional. Default is "doutok"
	Database string

	// Collection name
	//
	// Optional. Default is "default"
	Collection string

	// Reset clears any existing keys in existing Table
	//
	// Optional. Default is false
	Reset bool
}

// ConfigDefault is the default config
var ConfigDefault = Config{
	ConnectionURI: "",
	Host:          "127.0.0.1",
	Port:          27017,
	Database:      "doutok",
	Collection:    "default",
	Reset:         false,
}

// Helper function to set default values
func configDefault(config ...Config) Config {
	// Return default config if nothing provided
	if len(config) < 1 {
		return ConfigDefault
	}

	// Override default config
	cfg := config[0]

	// Set default values
	if cfg.Host == "" {
		cfg.Host = ConfigDefault.Host
	}
	if cfg.Port <= 0 {
		cfg.Port = ConfigDefault.Port
	}
	if cfg.Database == "" {
		cfg.Database = ConfigDefault.Database
	}
	if cfg.Collection == "" {
		cfg.Collection = ConfigDefault.Collection
	}
	return cfg
}
