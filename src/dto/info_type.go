package dto

type InfoType struct {
    Name        string `json:"name" binding:"required"`
    Description string `json:"description"`
    Regex       string `json:"regex" binding:"required"`
}

type InfoTypeResponse struct {
    InfoTypes   []InfoType `json:"info_types"`
}
