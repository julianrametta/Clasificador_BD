package routes

import (
    "clasificador_bd/controllers"
    "github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
    r.GET("/api/v1/ping", controllers.Ping)
    r.POST("/api/v1/database", controllers.CreateDBConfig)
    r.POST("/api/v1/database/scan/:id", controllers.CreateDBScan)
    r.GET("/api/v1/database/scan/:id", controllers.GetDBScan)
    r.POST("/api/v1/database/infotype/column", controllers.CreateColumnInfoType)
    r.GET("/api/v1/database/infotype/column", controllers.GetColumnInfoType)
    r.POST("/api/v1/database/infotype/table", controllers.CreateTableInfoType)
    r.GET("/api/v1/database/infotype/table", controllers.GetTableInfoType)
}

