package services

import (
	"clasificador_bd/models"
	"regexp"
)

type ClassificationService struct{}

func (s ClassificationService) ClassifyDBTablesCols(dbInfo []models.SchemaInfo, dbConfig models.DBConfig, columnInfoTypes, tableInfoTypes map[uint]string) ([]models.Schema, error) {
	var schemas []models.Schema
	var currentSchema *models.Schema
	var currentTable *models.Table

	for _, info := range dbInfo {
		if currentSchema == nil || currentSchema.Name != info.Esquema {
			newSchema := models.Schema{
				Name:       info.Esquema,
				DBConfigID: dbConfig.ID,
				Tables:     []models.Table{},
			}
			schemas = append(schemas, newSchema)
			currentSchema = &schemas[len(schemas)-1]
			currentTable = nil
		}

		if currentTable == nil || currentTable.Name != info.Tabla {
			infoTypeTableID, err := s.ClassifyByInfoType(info.Tabla, tableInfoTypes)
			if err != nil {
				return nil, err
			}
			newTable := models.Table{
				Name:               info.Tabla,
				InfoTypeTableNameID: infoTypeTableID,
				Columns:            []models.Column{},
			}
			currentSchema.Tables = append(currentSchema.Tables, newTable)
			currentTable = &currentSchema.Tables[len(currentSchema.Tables)-1]
		}

		infoTypeColID, err := s.ClassifyByInfoType(info.Columna, columnInfoTypes)
		if err != nil {
			return nil, err
		}
		column := models.Column{
			Name:               info.Columna,
			InfoTypeColNameID:  infoTypeColID,
		}
		currentTable.Columns = append(currentTable.Columns, column)
	}

	return schemas, nil
}

func (s ClassificationService) ClassifyByInfoType(name string, infoTypes map[uint]string) (uint, error) {
	var defaultInfoTypeID uint = 1
	for infoType, regex := range infoTypes {
		isThisInfoType, err := s.IsThisInfoType(name, regex)
		if err != nil {
			return 0, err
		}
		if isThisInfoType {
			return infoType, nil
		}
	}
	return defaultInfoTypeID, nil
}

func (s ClassificationService) IsThisInfoType(name string, types string) (bool, error) {
	if types == ""{
		return false, nil
	}
	regexNoDoubleSlash := regexp.MustCompile(`\\+`)
	expression := regexNoDoubleSlash.ReplaceAllString(types, `\`)

	re, err := regexp.Compile(expression)
	if err != nil {
		return false, err
	}

	return re.MatchString(name), nil
}
