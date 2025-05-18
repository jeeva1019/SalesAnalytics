package models

import "gorm.io/gorm"

type SalesModel struct {
	DB *gorm.DB
}

func NewSalesModel(db *gorm.DB) *SalesModel {
	return &SalesModel{db}
}

type FinalResponse struct {
	Status     string `json:"status"`
	StatusCode string `json:"statuscode,omitempty"`
	Message    string `json:"message,omitempty"`
	Result     any    `json:"result,omitempty"`
}
