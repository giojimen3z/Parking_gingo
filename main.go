package main

import (
	"time"

	"Paking_gingo/cmd/api/controllers"
	"Paking_gingo/cmd/api/database"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	database.ConnectDatabase()

	// Configurar CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.GET("/vehicles", controllers.FindVehicles)
	r.GET("/vehicles/:license_plate", controllers.FindVehicleByLicensePlate)
	r.POST("/vehicles", controllers.CreateVehicle)
	r.POST("/vehicles/checkout", controllers.CheckOutVehicle)

	r.Run() // por defecto, corre en localhost:8080
}
