package config

import (
	"go-fiber-clean-architecture/application/helper/helper2"
)

var (
	AppMode = helper2.GetEnv("APP_MODE", "development")
	AppPort = helper2.GetEnv("APP_PORT", "3000")
	AppUrl = helper2.GetEnv("APP_URL", "localhost")
)