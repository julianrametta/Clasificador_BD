package repositories

import (
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "fmt"
)
var DB *gorm.DB

func ConnectDbGlobalSession(dsn string) {
    DB, _ = ConnectDbLocalSession(dsn)
}

func ConnectDbLocalSession(dsn string) (*gorm.DB, error) {
    fmt.Println("db connect string", dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        fmt.Println("Failed to connect database", err)
		return nil, err
    }

	fmt.Println("Connection to DB was succesfully done!")
    return db, nil
}