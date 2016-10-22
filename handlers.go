package main

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Charge handler will do the charging by adding Server Key into header
func Charge(c *gin.Context) {
	// Encode server key using base 64 string
	authorization := base64.StdEncoding.EncodeToString([]byte(VTServerKey + ":"))

	// HTTP client
	client := http.DefaultClient
	var URL = SnapURL
	if EnableProduction {
		URL = SnapURLProduction
	}
	request, err := http.NewRequest("POST", URL+"/transactions", c.Request.Body)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status_code": "400", "status_message": "Bad Request"})
	} else {
		request.Header.Add("Content-Type", "application/json")
		request.Header.Add("Accept", "application/json")
		request.Header.Add("Authorization", "Basic "+authorization)
		response, _ := client.Do(request)

		responseBody, _ := ioutil.ReadAll(response.Body)
		var respObj interface{}
		json.Unmarshal(responseBody, &respObj)
		c.JSON(http.StatusOK, respObj)
	}
}
