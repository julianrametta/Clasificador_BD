package controllers

import (
	"fmt"
	"net/http"
	"clasificador_bd/dto"
	"clasificador_bd/services"
	"github.com/gin-gonic/gin"
)

func CreateDBConfig(c *gin.Context) {
	var req dto.DBConfigRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println("Invalid request: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dbConfig, err := services.DBConfigService{}.CreateNewDBConfig(req.Host, req.Port, req.Username, req.Password)
	if err != nil {
		fmt.Println("Error creating database config:", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating database config"})
		return
	}

	response := gin.H{"message": "New db config was successfully stored", "id": dbConfig.ID}
	c.JSON(http.StatusCreated, response)
}