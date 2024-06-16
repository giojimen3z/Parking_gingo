package models

type Vehicle struct {
	gorm.Model
	LicensePlate string `json:"license_plate"`
	Brand        string `json:"brand"`
	Model        string `json:"model"`
	Color        string `json:"color"`
}
