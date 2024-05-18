package utils

import (
	"flag"
	"fmt"
	"go-fiber-clean-architecture/application/config"
	"go-fiber-clean-architecture/application/utils/helper2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"strconv"
	"time"
)

// ConnectDB untuk menghubungkan ke database
func ConnectDB() {
	// jika MODE test dijalankan
	if flag.Lookup("test.v") != nil {
		config.DbPort     = helper2.GetEnv("DB_PORT_TEST", "3306")          // default port 3306 yaitu port mysql
		config.DbDriver   = helper2.GetEnv("DB_DRIVER_TEST", "mysql")       // default driver mysql
		config.DbUser     = helper2.GetEnv("DB_USER_TEST", "root")          // default user root
		config.DbPassword = helper2.GetEnv("DB_PASSWORD_TEST", "")          // default password root
		config.DbHost     = helper2.GetEnv("DB_HOST_TEST", "localhost")     // default host localhost
		config.DbName     = helper2.GetEnv("DB_NAME_TEST", "go_fiber_test") // default database go_fiber
	}

	var err error

	// inisialisasi database
	// pengecekan port ada atau tidak
	port, err := strconv.ParseUint(config.DbPort, 10, 32)
	if err != nil {
		panic("port error")
	}

	// inisialisasi dsn
	var dsn string

	// mengkonekkan ke database
	if config.DbDriver == "mysql" {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.DbUser, config.DbPassword, config.DbHost, port, config.DbName)
		config.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		})
	}
	//else if Config("DB_DRIVER") == "postgres" {
	//	dsn = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", Config("DB_HOST"), port, Config("DB_USER"), Config("DB_PASSWORD"), Config("DB_NAME"))
	//	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	//}

	fmt.Println(dsn)

	// pengecekan error
	if err != nil {
		panic("failed to connect database")
	}

	// set max idle connection
	sqlDB, err := config.DB.DB()
	if err != nil {
		panic("error db in 46 line")
	}
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetMaxOpenConns(20)
	sqlDB.SetConnMaxLifetime(60 * time.Minute)
	sqlDB.SetConnMaxIdleTime(10 * time.Minute)

	fmt.Println("Connection Opened to Database")

	// Migrate the schema
	//DB.AutoMigrate(&model.Todo{})
	//fmt.Println("Database Migrated")
}
