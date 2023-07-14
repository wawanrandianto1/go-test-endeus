package endeus

//go:generate mockgen -source=config.go -package=mocks -destination=../../internal/mocks/config.go

type Config interface {
	IsDev() bool
	Environment() string
	ServerPort() string
	DBHost() string
	DBPort() string
	DBName() string
	DBUsername() string
	DBPassword() string
	DBAddress() string
	LogLevel() string
}
