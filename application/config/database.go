package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strconv"
	"time"
)

// yang akan di gunakan ketika kita ingin mengakses database
var (
	DB *gorm.DB
	DbPort = Config("DB_PORT", "3306") // default port 3306 yaitu port mysql
	DbDriver = Config("DB_DRIVER", "mysql") // default driver mysql
	DbUser = Config("DB_USER", "root") // default user root
	DbPassword = Config("DB_PASSWORD", "password") // default password root
	DbHost = Config("DB_HOST", "localhost") // default host localhost
	DbName = Config("DB_NAME", "go_fiber") // default database go_fiber
)

// ConnectDB untuk menghubungkan ke database
func ConnectDB() {
	var err error

	// inisialisasi database
	// pengecekan port ada atau tidak
	port, err := strconv.ParseUint(DbPort, 10, 32)
	if err != nil {
		panic("port error")
	}

	// inisialisasi dsn
	var dsn string

	// mengkonekkan ke database
	if DbDriver == "mysql" {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, port, DbName)
		DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
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
	sqlDB, err := DB.DB()
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