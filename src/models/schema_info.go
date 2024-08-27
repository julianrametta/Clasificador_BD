package models

type SchemaInfo struct{
    Esquema string `gorm:"column:esquema"`
    Tabla   string `gorm:"column:tabla"`
    Columna string `gorm:"column:columna"`
}