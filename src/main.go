package main

import (
    "clasificador_bd/repositories"
    "clasificador_bd/routes"
    "clasificador_bd/config"
    "github.com/gin-gonic/gin"
    "fmt"
)
func main() {
    config := config.Config{}
    config.Load()

    dsn := config.DBUser + ":" + config.DBPassword + "@tcp(" + config.DBHost + ":" + config.DBPort + ")/" + config.DBName + "?charset=utf8mb4&parseTime=True&loc=Local"
    
    fmt.Println("Debug:", dsn)

    repositories.ConnectDbGlobalSession(dsn)
    r := gin.Default()

    routes.SetupRoutes(r)

    r.Run(":8081")
}