package controllers

import (
	"fmt"
	"net/http"
	"clasificador_bd/dto"
	"clasificador_bd/services"
	"github.com/gin-gonic/gin"
)

func CreateColumnInfoType(c *gin.Context) {
	var req dto.InfoType
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println("Invalid request: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	infoTypeService := services.InfoTypeService{}
	isValidRegex := infoTypeService.ValidateRegex(req.Regex)
	if !isValidRegex {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Regex"})
		return
	}

	newColumnInfoType, err := infoTypeService.CreateColumnInfoType(req)

	if err != nil {
		fmt.Println("Error creating column info type: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating column info type"})
		return
	}


	c.JSON(http.StatusOK, gin.H{"message": "New column info type successfully stored", "id": newColumnInfoType.ID})

}

func GetColumnInfoType(c *gin.Context) {
	infoTypeService := services.InfoTypeService{}
	columnInfoTypes, err := infoTypeService.GetColumnInfoTypes()
	if err != nil {
		fmt.Println("Error reading column info types: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error reading column info types"})
		return
	}

	c.JSON(http.StatusOK, columnInfoTypes)
}

func CreateTableInfoType(c *gin.Context) {
	var req dto.InfoType
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println("Invalid request: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	infoTypeService := services.InfoTypeService{}
	isValidRegex := infoTypeService.ValidateRegex(req.Regex)
	if !isValidRegex {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Regex"})
		return
	}

	newTableInfoType, err := infoTypeService.CreateTableInfoType(req)

	if err != nil {
		fmt.Println("Error creating table info type", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating table info type"})
		return
	}


	c.JSON(http.StatusOK, gin.H{"message": "New table info type successfully stored", "id": newTableInfoType.ID})

}

func GetTableInfoType(c *gin.Context) {
	infoTypeService := services.InfoTypeService{}
	tableInfoTypes, err := infoTypeService.GetTableInfoTypes()
	if err != nil {
		fmt.Println("Error reading table info types: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error reading table info types"})
		return
	}

	c.JSON(http.StatusOK, tableInfoTypes)
}