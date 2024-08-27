package repositories

import (
	"clasificador_bd/models"
	"fmt"
)

type DBConfigRepository struct{}


func (r DBConfigRepository) GetByID(id string) (*models.DBConfig, error) {
	var db_config models.DBConfig
    result := DB.First(&db_config, id)

    if result.Error != nil {
        fmt.Println("Error when reading DB Config registry", result.Error)
		return nil, result.Error
    }
	fmt.Println("DB Config registry was succesfully obtained", db_config.ID, db_config.ConnectionString)
	return &db_config, nil
}

func (r DBConfigRepository) Create(connectionString string) (*models.DBConfig, error) {
	db_config := models.DBConfig{ConnectionString: connectionString}
    result := DB.Create(&db_config)

    if result.Error != nil {
        fmt.Println("Error when inserting New DB Config registry", result.Error)
		return nil, result.Error

    }
	fmt.Println("New DB Config registry was succesfully inserted", db_config.ID)
	return &db_config, nil
}