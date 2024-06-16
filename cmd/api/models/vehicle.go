package models

import (
	"time"

	"gorm.io/gorm"
)

type Vehicle struct {
	gorm.Model
	LicensePlate string    `json:"license_plate"`
	Brand        string    `json:"brand"`
	VehicleModel string    `json:"vehicle_model"`
	Color        string    `json:"color"`
	CheckInTime  time.Time `json:"check_in_time"`
	CheckOutTime time.Time `json:"check_out_time"`
	Cost         float64   `json:"cost"`
}

func (v *Vehicle) CalculateCost(ratePerHour float64) {
	duration := v.CheckOutTime.Sub(v.CheckInTime)
	hours := duration.Hours()
	v.Cost = hours * ratePerHour
}
