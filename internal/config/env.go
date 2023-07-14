package config

import (
	"endeus/wawan/pkg/utils"
	"fmt"
)

type config struct {
	Loglevel   string // TRACE, DEBUG, INFO, WARN, ERROR, FATAL, PANIC
	Serverport string
	Envi       string
	Dbname     string
	Dbhost     string
	Dbport     string
	Dbuser     string
	Dbpass     string
}

type Config struct {
	config
}

func New() Config {
	var cfg config
	cfg.Loglevel = utils.GetEnv("LOG_LEVEL", "DEBUG")
	cfg.Envi = utils.GetEnv("ENVIRONMENT", "dev")
	cfg.Serverport = utils.GetEnv("SERVER_PORT", "5000")
	cfg.Dbname = utils.GetEnv("DB_DATABASE", "")
	cfg.Dbhost = utils.GetEnv("DB_HOST", "127.0.0.1")
	cfg.Dbport = utils.GetEnv("DB_PORT", "3306")
	cfg.Dbuser = utils.GetEnv("DB_USERNAME", "")
	cfg.Dbpass = utils.GetEnv("DB_PASSWORD", "")
	return Config{config: cfg}
}

func (c Config) IsDev() bool {
	return c.Envi == "dev"
}

func (c Config) Environment() string {
	return c.Envi
}

func (c Config) LogLevel() string {
	if c.Loglevel == "" {
		return "DEBUG"
	}
	return c.Loglevel
}

func (c Config) ServerPort() string {
	return c.Serverport
}

func (c Config) DBName() string {
	return c.Dbname
}

func (c Config) DBUsername() string {
	return c.Dbuser
}

func (c Config) DBPassword() string {
	return c.Dbpass
}

func (c Config) DBHost() string {
	return c.Dbhost
}

func (c Config) DBPort() string {
	return c.Dbport
}

func (c Config) DBAddress() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.DBUsername(),
		c.DBPassword(),
		c.DBHost(),
		c.DBPort(),
		c.DBName(),
	)
}
