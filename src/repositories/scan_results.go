package repositories

import (
	"clasificador_bd/models"
	"gorm.io/gorm"
	"fmt"
)

type ScanResultsRepository struct{}

func (s ScanResultsRepository) Update(dbConfigID uint, schemas *[]models.Schema) error {
    tx := DB.Begin()
    if tx.Error != nil {
        return tx.Error
    }
    if err := s.Delete(tx, dbConfigID); err != nil {
        tx.Rollback()
        return err
    }
    if err := s.Create(tx, schemas); err != nil {
        tx.Rollback()
        return err
    }
    if err := tx.Commit().Error; err != nil {
        return err
    }
    return nil
}

func (s ScanResultsRepository) Delete(tx *gorm.DB, dbConfigID uint) error {
    if err := tx.Where("db_config_id = ?", dbConfigID).Delete(&models.Schema{}).Error; err != err {
		fmt.Println("Error when deleting past DB scan result: ", err.Error)
        return err
    }
    return nil
}

func (s ScanResultsRepository) Create(tx *gorm.DB, schemas *[]models.Schema) error {
    result := tx.Create(schemas)
    if result.Error != nil {
        fmt.Println("Error when inserting new DB scan result: ", result.Error)
        return result.Error
    }
    fmt.Println("New DB scan result was successfully inserted")
    return nil
}

func (s ScanResultsRepository) GetByDbConfigID(dbConfigID string) (*[]models.Schema, error) {
    var schemas []models.Schema

    err := DB.Where("db_config_id = ?", dbConfigID).
    Preload("Tables.InfoTypeTableName").
    Preload("Tables.Columns.InfoTypeColName").
    Find(&schemas).Error

    if err != nil {
        fmt.Println("Error retrieving scan results from db:", err)
        return nil, err
    }

    return &schemas, nil

}