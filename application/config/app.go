package config

var (
	AppMode = Config("APP_MODE", "development")
	AppPort = Config("APP_PORT", "3000")
	AppUrl = Config("APP_URL", "localhost")
)