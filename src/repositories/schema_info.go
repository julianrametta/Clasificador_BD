package repositories

import (
	"clasificador_bd/models"
	"fmt"
)

type SchemaInfoRepository struct{}

func (r SchemaInfoRepository) GetInfoByConfig(dbConfig models.DBConfig) ([]models.SchemaInfo, error) {
	var dbInfo []models.SchemaInfo
	db, err := ConnectDbLocalSession(dbConfig.ConnectionString)
	if err != nil {
		fmt.Println("Failed to connect to database:", err)
		return nil, err
	}

	result := db.Raw(`
		SELECT 
			TABLE_SCHEMA AS esquema,
			TABLE_NAME AS tabla,
			COLUMN_NAME AS columna
		FROM 
			INFORMATION_SCHEMA.COLUMNS
		WHERE
			TABLE_SCHEMA NOT IN ('information_schema', 'mysql', 'performance_schema', 'sys')
		ORDER BY 
			TABLE_SCHEMA, TABLE_NAME, ORDINAL_POSITION`).Scan(&dbInfo)
	
	if result.Error != nil {
		fmt.Println("Error retrieving database info:", result.Error)
		return nil, result.Error
	}
	return dbInfo, nil
}