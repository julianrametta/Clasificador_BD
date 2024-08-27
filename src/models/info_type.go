package models

type InfoTypeColName struct {
    ID          uint   `gorm:"primaryKey"`
    Name        string
    Description string
    Regex       string
    Columns      []Column
}

type InfoTypeTableName struct {
    ID          uint   `gorm:"primaryKey"`
    Name        string
    Description string
    Regex       string
    Tables      []Table
}