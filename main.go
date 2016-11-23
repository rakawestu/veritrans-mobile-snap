package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

// SnapURL is snap endpoint
var SnapURL = "https://app.sandbox.veritrans.co.id/snap/v1"

// SnapURLProduction is snap endpoint in production mode
var SnapURLProduction = "https://app.midtrans.com/snap/v1"

// VTServerKey is server key of Veritrans PAPI
var VTServerKey = ""

// EnableProduction is environment variable to set production mode
var EnableProduction = false

func main() {
	port := os.Getenv("PORT")
	serverKey := os.Getenv("SERVER_KEY")
	enableProduction := os.Getenv("PRODUCTION")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	if serverKey != "" {
		VTServerKey = serverKey
	} else {
		log.Fatal("$SERVER_KEY must be set")
	}

	if enableProduction != "" && enableProduction == "true" {
		EnableProduction = true
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.POST("/charge", Charge)
	router.POST("/installment/charge", ChargeWithInstallment)

	router.Run(":" + port)
}
