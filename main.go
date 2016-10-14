package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

// SnapURL is snap endpoint
var SnapURL = "https://app.sandbox.veritrans.co.id/snap/v1"

// VTServerKey is server key
var VTServerKey = "VT-server-F_UKmzr3AbJv07Lupq_KCLPV"

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.POST("/charge", Charge)

	router.Run(":" + port)
}
