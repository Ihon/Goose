package system

import (
    "fmt"
    "log"
    // "gorm.io/driver/postgres"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB(config *Config) {
    var err error
    // dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", config.DBHost, config.DBUserName, config.DBUserPassword, config.DBName, config.DBPort)
    dsn := fmt.Sprintf(
        "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        config.DBUserName,
        config.DBUserPassword,
        config.DBHost,
        config.DBPort,
        config.DBName,
    )

    DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to the Database")
    }
    fmt.Println("? Connected Successfully to the Database")
}