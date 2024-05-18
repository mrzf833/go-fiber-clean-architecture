package config

import (
	"go-fiber-clean-architecture/application/utils/helper2"
	"gorm.io/gorm"
)

// yang akan di gunakan ketika kita ingin mengakses database
var (
	DB         *gorm.DB
	DbPort     = helper2.GetEnv("DB_PORT", "3306")      // default port 3306 yaitu port mysql
	DbDriver   = helper2.GetEnv("DB_DRIVER", "mysql")   // default driver mysql
	DbUser     = helper2.GetEnv("DB_USER", "root")      // default user root
	DbPassword = helper2.GetEnv("DB_PASSWORD", "")      // default password root
	DbHost     = helper2.GetEnv("DB_HOST", "localhost") // default host localhost
	DbName     = helper2.GetEnv("DB_NAME", "go_fiber")  // default database go_fiber
)