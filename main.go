package main

import (
	"Paking_gingo/cmd/api/controllers"
	"Paking_gingo/cmd/api/database"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	database.ConnectDatabase()

	r.GET("/vehicles", controllers.FindVehicles)
	r.GET("/vehicles/:license_plate", controllers.FindVehicleByLicensePlate)
	r.POST("/vehicles", controllers.CreateVehicle)
	r.POST("/vehicles/checkout", controllers.CheckOutVehicle)

	r.Run() // por defecto, corre en localhost:8080
}
