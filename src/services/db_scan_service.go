package services

import (
	"clasificador_bd/repositories"
	"clasificador_bd/models"
	"clasificador_bd/dto"
)

type DBScanService struct{}

func (s DBScanService) ExecuteScan(db_id string) error {
	dbConfig, err := repositories.DBConfigRepository{}.GetByID(db_id)
	if err != nil {
		return err
	}

	dbInfo, err := repositories.SchemaInfoRepository{}.GetInfoByConfig(*dbConfig)
	if err != nil {
		return err
	}


	columnInfoTypes, err := repositories.InfoTypeRepository{}.GetMapColumnInfoTypeIDToRegex()
	if err != nil {
		return err
	}

	tableInfoTypes, err := repositories.InfoTypeRepository{}.GetMapTableInfoTypeIDToRegex()
	if err != nil {
		return err
	}

	scanResults, err := ClassificationService{}.ClassifyDBTablesCols(dbInfo, *dbConfig, columnInfoTypes, tableInfoTypes)
	if err != nil {
		return err
	}

	return repositories.ScanResultsRepository{}.Update(dbConfig.ID, &scanResults)
}

func (s DBScanService) GetScanResults(db_id string) (*dto.ScanResponse, error) {
	scanResults, err := repositories.ScanResultsRepository{}.GetByDbConfigID(db_id)
	if err != nil {
		return nil, err
	}
	scanResponse := s.GetScanResponseFromSchemas(*scanResults)
	return scanResponse, nil
}

func (s DBScanService) GetScanResponseFromSchemas(schemas []models.Schema) *dto.ScanResponse{
	var response dto.ScanResponse

	for _, schema := range schemas {
		schemaResp := dto.SchemaResponse{
			Name:   schema.Name,
			Tables: []dto.TableResponse{},
		}

		for _, table := range schema.Tables {
			tableResp := dto.TableResponse{
				Name:              table.Name,
				InfoType: table.InfoTypeTableName.Name,
				Columns:           []dto.ColumnResponse{},
			}

			for _, column := range table.Columns {
				columnResp := dto.ColumnResponse{
					Name:            column.Name,
					InfoType: column.InfoTypeColName.Name,
				}
				tableResp.Columns = append(tableResp.Columns, columnResp)
			}

			schemaResp.Tables = append(schemaResp.Tables, tableResp)
		}

		response.Schemas = append(response.Schemas, schemaResp)
	}

	return &response
}
