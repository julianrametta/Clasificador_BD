package dto

type DBConfigRequest struct {
    Host     string `json:"host" binding:"required"`
    Port     string `json:"port" binding:"required"`
    Username string `json:"username" binding:"required"`
    Password string `json:"password" binding:"required"`
}