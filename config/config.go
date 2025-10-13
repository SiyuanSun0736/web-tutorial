package config

type Config struct {
	Database DatabaseConfig
	Server   ServerConfig
}

type DatabaseConfig struct {
	Host     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

type ServerConfig struct {
	Port string
}

func LoadConfig() *Config {
	return &Config{
		Database: DatabaseConfig{
			Host:     "localhost",
			User:     "user",
			DBName:   "test",
			SSLMode:  "disable",
			Password: "passwd",
		},
		Server: ServerConfig{
			Port: ":8080",
		},
	}
}
