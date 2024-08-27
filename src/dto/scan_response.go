package dto

type ColumnResponse struct {
    Name               string `json:"name"`
    InfoType    string `json:"info_type"`
}

type TableResponse struct {
    Name               string           `json:"name"`
    InfoType  string           `json:"info_type"`
    Columns            []ColumnResponse `json:"columns"`
}

type SchemaResponse struct {
    Name   string         `json:"name"`
    Tables []TableResponse `json:"tables"`
}

type ScanResponse struct {
    Schemas []SchemaResponse `json:"schemas"`
}
