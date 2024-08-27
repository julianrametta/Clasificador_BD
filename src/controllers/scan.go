package controllers

import (
	"fmt"
	"net/http"
	"clasificador_bd/services"
	"github.com/gin-gonic/gin"
)

func CreateDBScan(c *gin.Context) {
	db_id := c.Param("id")

	err := services.DBScanService{}.ExecuteScan(db_id)
	if err != nil {
		fmt.Println("Error executing DB scan:", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error performing database scan analysis."})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "New database scan was successfully executed."})
}

func GetDBScan(c *gin.Context) {
	db_id := c.Param("id")

	scanResults, err := services.DBScanService{}.GetScanResults(db_id)
	if err != nil {
		fmt.Println("Error reading database scan results:", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error reading database scan results."})
		return
	}

	c.JSON(http.StatusOK, scanResults)
}