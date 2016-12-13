package main

import (
	"log"
	"os"

	"gopkg.in/mgo.v2"

	"github.com/gin-gonic/gin"
)

// MongoDBUrl url of mongo db
var MongoDBUrl string

// MongoDB database connection
var MongoDB *mgo.Database

// SnapURL is snap endpoint
var SnapURL = "https://app.sandbox.midtrans.com/snap/v1"

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
	MongoDBUrl := os.Getenv("MONGODB_URL")
	MongoDBName := os.Getenv("MONGODB_NAME")

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

	if MongoDBUrl == "" || MongoDBName == "" {
		log.Fatal("$MONGODB_URL and $MONGODB_NAME must be set")
	}

	MongoSession, err := mgo.Dial(MongoDBUrl)
	if err != nil {
		panic(err)
	}
	defer MongoSession.Close()
	MongoSession.SetMode(mgo.Monotonic, true)
	MongoDB = MongoSession.DB(MongoDBName)

	router := gin.New()
	router.Use(gin.Logger())
	router.POST("/installment/charge", ChargeWithInstallment)
	router.GET("/installment/users/:id/tokens", GetCardsEndpoint)
	router.POST("/installment/users/:id/tokens", SaveCardsEndpoint)
	router.POST("/charge", Charge)
	router.GET("/users/:id/tokens", GetCardsEndpoint)
	router.POST("/users/:id/tokens", SaveCardsEndpoint)

	router.Run(":" + port)
}
