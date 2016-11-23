package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/bitly/go-simplejson"
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

// ChargeWithInstallment will do the charging with added installment
func ChargeWithInstallment(c *gin.Context) {
	// Encode server key using base 64 string
	authorization := base64.StdEncoding.EncodeToString([]byte(VTServerKey + ":"))

	// HTTP client
	client := http.DefaultClient
	var URL = SnapURL
	if EnableProduction {
		URL = SnapURLProduction
	}

	installment := getInstallmentData()
	whitelist := getWhitelistBin()
	requestBody, _ := ioutil.ReadAll(c.Request.Body)
	requestJSON, _ := simplejson.NewJson(requestBody)
	creditCard := requestJSON.Get("credit_card")
	if creditCard != nil {
		creditCard.Set("installment", CreditCard{Installment: installment, WhitelistBin: whitelist})
		requestJSON.Set("credit_card", creditCard)
	}

	requestJSON.Set("credit_card", CreditCard{Installment: installment})
	requestJSONMarshaled, _ := requestJSON.MarshalJSON()
	requestObj := bytes.NewReader(requestJSONMarshaled)
	request, err := http.NewRequest("POST", URL+"/transactions", requestObj)

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

func getInstallmentData() Installment {
	return Installment{Required: false, Terms: Terms{BNI: []int{3, 6, 12}, Mandiri: []int{3, 6, 12}, BCA: []int{3, 6, 12}, CIMB: []int{3, 6, 12}, Offline: []int{3, 6, 12}}}
}

func getWhitelistBin() []string {
	return []string{"481111", "521111"}
}
