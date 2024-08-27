package models

type DBConfig struct {
    ID              uint   `gorm:"primaryKey;autoIncrement"`
    ConnectionString string `gorm:"type:text"`
    Schemas []Schema
}
