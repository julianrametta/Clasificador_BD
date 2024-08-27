package config

import (
    "github.com/joho/godotenv"
    "fmt"
    "os"
)

type Config struct {
    DBUser     string
    DBPassword string
    DBHost     string
    DBPort     string
    DBName     string
}

func (c *Config) Load() {
    _ = godotenv.Load()

    c.DBUser = os.Getenv("DB_USER")
    c.DBPassword = os.Getenv("DB_PASSWORD")
    c.DBHost = os.Getenv("DB_HOST")
    c.DBPort = os.Getenv("DB_PORT")
    c.DBName = os.Getenv("DB_NAME")

    if c.DBUser == "" || c.DBPassword == "" || c.DBHost == "" || c.DBPort == "" || c.DBName == "" {
        fmt.Printf("Error when loading configuration env variables")
        os.Exit(1)
    }
}