package models

type Schema struct {
    ID       uint `gorm:"primaryKey;autoIncrement"`
    Name     string
    DBConfigID uint
    Tables     []Table
}

type Table struct {
    ID           uint `gorm:"primaryKey;autoIncrement"`
    Name         string
    SchemaID     uint
    InfoTypeTableNameID uint
    InfoTypeTableName  InfoTypeTableName `gorm:"foreignKey:InfoTypeTableNameID"`
    Columns       []Column
}

type Column struct {
    ID           uint `gorm:"primaryKey;autoIncrement"`
    Name         string
    InfoTypeColNameID uint
    InfoTypeColName  InfoTypeColName `gorm:"foreignKey:InfoTypeColNameID;references:ID"`
    TableID      uint
}