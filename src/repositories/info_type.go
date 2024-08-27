package repositories

import (
	"clasificador_bd/models"
	"fmt"
)

type InfoTypeRepository struct{}

func (r InfoTypeRepository) GetAllColumnInfoTypes() ([]models.InfoTypeColName, error) {
	var records []models.InfoTypeColName
	result := DB.Find(&records)
	if result.Error != nil {
		fmt.Println("Error retrieving column info types:", result.Error)
		return nil, result.Error
	}

	return records, nil
}

func (r InfoTypeRepository) GetMapColumnInfoTypeIDToRegex() (map[uint]string, error) {
	records, err := r.GetAllColumnInfoTypes()
	if err != nil {
		return nil, err
	}

	infoTypeMap := make(map[uint]string)
	for _, record := range records {
		infoTypeMap[record.ID] = record.Regex
	}

	return infoTypeMap, nil
}

func (r InfoTypeRepository) InsertColumnInfoType(colInfoType models.InfoTypeColName) (*models.InfoTypeColName, error) {
    result := DB.Create(&colInfoType)

    if result.Error != nil {
        fmt.Println("Error when creating new column info type", result.Error)
		return nil, result.Error

    }
	fmt.Println("New column info type was succesfully created", colInfoType.ID)
	return &colInfoType, nil
}

func (r InfoTypeRepository) GetAllTableInfoTypes() ([]models.InfoTypeTableName, error) {
	var records []models.InfoTypeTableName
	result := DB.Find(&records)
	if result.Error != nil {
		fmt.Println("Error retrieving table info types:", result.Error)
		return nil, result.Error
	}

	return records, nil
}

func (r InfoTypeRepository) GetMapTableInfoTypeIDToRegex() (map[uint]string, error) {
	records, err := r.GetAllTableInfoTypes()
	if err != nil {
		return nil, err
	}

	infoTypeMap := make(map[uint]string)
	for _, record := range records {
		infoTypeMap[record.ID] = record.Regex
	}

	return infoTypeMap, nil
}


func (r InfoTypeRepository) InsertTableInfoType(tableInfoType models.InfoTypeTableName) (*models.InfoTypeTableName, error) {
    result := DB.Create(&tableInfoType)

    if result.Error != nil {
        fmt.Println("Error when creating new table info type", result.Error)
		return nil, result.Error

    }
	fmt.Println("New table info type was succesfully created", tableInfoType.ID)
	return &tableInfoType, nil
}