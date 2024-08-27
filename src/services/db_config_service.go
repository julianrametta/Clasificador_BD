package services

import (
	"clasificador_bd/models"
	"clasificador_bd/repositories"
	"fmt"
)

type DBConfigService struct{}

func (s DBConfigService) CreateNewDBConfig(host, port, username, password string) (*models.DBConfig, error) {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/", username, password, host, port)

	dbConfig, err := repositories.DBConfigRepository{}.Create(connectionString)
	if err != nil {
		return nil, fmt.Errorf("Error when creating new db config: %v", err)
	}

	return dbConfig, nil
}