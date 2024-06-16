package controllers

import (
	"net/http"
	"time"

	"Paking_gingo/cmd/api/database"
	"Paking_gingo/cmd/api/models"
	"github.com/gin-gonic/gin"
)

const RatePerHour = 5.0 // Define the rate per hour for parking

func FindVehicles(c *gin.Context) {
	var vehicles []models.Vehicle
	database.DB.Find(&vehicles)
	c.JSON(http.StatusOK, gin.H{"data": vehicles})
}

func FindVehicleByLicensePlate(c *gin.Context) {
	var vehicle models.Vehicle
	licensePlate := c.Param("license_plate")
	if err := database.DB.Where("license_plate = ?", licensePlate).First(&vehicle).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Vehicle not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": vehicle})
}

func CreateVehicle(c *gin.Context) {
	var input models.Vehicle
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	vehicle := models.Vehicle{
		LicensePlate: input.LicensePlate,
		Brand:        input.Brand,
		VehicleModel: input.VehicleModel,
		Color:        input.Color,
		CheckInTime:  time.Now(),
	}
	database.DB.Create(&vehicle)
	c.JSON(http.StatusOK, gin.H{"data": vehicle})
}

func CheckOutVehicle(c *gin.Context) {
	var input struct {
		LicensePlate string `json:"license_plate"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var vehicle models.Vehicle
	if err := database.DB.Where("license_plate = ?", input.LicensePlate).First(&vehicle).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Vehicle not found"})
		return
	}

	vehicle.CheckOutTime = time.Now()
	vehicle.CalculateCost(RatePerHour)
	database.DB.Save(&vehicle)

	c.JSON(http.StatusOK, gin.H{"data": vehicle})
}
