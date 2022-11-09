package models

type HealthResponse struct {
	Http  string `json:"http"`
	Mongo string `json:"mongo"`
}
