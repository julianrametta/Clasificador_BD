package services

import (
	"clasificador_bd/dto"
	"clasificador_bd/models"
	"clasificador_bd/repositories"
	"regexp"
	"strings"
)

type InfoTypeService struct{}

func (s InfoTypeService) ValidateRegex(req_regex string) bool {
	regexPattern := `^\\b\([A-Za-z]+\)\\b(\|\\b\([A-Za-z]+\)\\b)*$`
	regex, _ := regexp.Compile(regexPattern)
	return regex.MatchString(req_regex)
}

func (s InfoTypeService) CreateColumnInfoType(info_type_req dto.InfoType) (*models.InfoTypeColName, error) {
	// Replace '\' by '\\'
	regex := strings.ReplaceAll(info_type_req.Regex, "\\", "\\\\")

	infoColumnType := models.InfoTypeColName{Name: info_type_req.Name, Description: info_type_req.Description, Regex: regex}

	newColumnInfoType, err := repositories.InfoTypeRepository{}.InsertColumnInfoType(infoColumnType)
	if err != nil {
		return nil, err
	}

	return newColumnInfoType, nil
}

func (s InfoTypeService) GetColumnInfoTypes() (*dto.InfoTypeResponse, error) {
	columnInfoTypeResults, err := repositories.InfoTypeRepository{}.GetAllColumnInfoTypes()
	if err != nil {
		return nil, err
	}

	var response dto.InfoTypeResponse
	for _, columnInfoTypeResult := range columnInfoTypeResults{
		columnInfoType := dto.InfoType{
			Name:   columnInfoTypeResult.Name,
			Description: columnInfoTypeResult.Description,
			Regex: columnInfoTypeResult.Regex,
		}
		response.InfoTypes = append(response.InfoTypes, columnInfoType)
	}
	return &response, nil
}

func (s InfoTypeService) CreateTableInfoType(info_type_req dto.InfoType) (*models.InfoTypeTableName, error) {
	// Replace '\' by '\\'
	regex := strings.ReplaceAll(info_type_req.Regex, "\\", "\\\\")

	infoTableType := models.InfoTypeTableName{Name: info_type_req.Name, Description: info_type_req.Description, Regex: regex}

	newTableInfoType, err := repositories.InfoTypeRepository{}.InsertTableInfoType(infoTableType)
	if err != nil {
		return nil, err
	}

	return newTableInfoType, nil
}

func (s InfoTypeService) GetTableInfoTypes() (*dto.InfoTypeResponse, error) {
	tableInfoTypeResults, err := repositories.InfoTypeRepository{}.GetAllTableInfoTypes()
	if err != nil {
		return nil, err
	}

	var response dto.InfoTypeResponse
	for _, tableInfoTypeResult := range tableInfoTypeResults{
		tableInfoType := dto.InfoType{
			Name:   tableInfoTypeResult.Name,
			Description: tableInfoTypeResult.Description,
			Regex: tableInfoTypeResult.Regex,
		}
		response.InfoTypes = append(response.InfoTypes, tableInfoType)
	}
	return &response, nil
}

