package model

type Customer struct {
	Id      string `json:"id" db:"id"`
	ZipCode string `json:"zip_code" db:"zip_code"`
	City    string `json:"city" db:"city"`
	State   string `json:"state" db:"state"`
}
