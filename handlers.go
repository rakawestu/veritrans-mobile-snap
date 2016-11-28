package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

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
	if creditCard == nil {
		requestJSON.Set("credit_card", CreditCard{Installment: installment, WhitelistBins: whitelist})
	} else {
		creditCard.Set("installment", installment)
		creditCard.Set("whitelist_bins", whitelist)
		requestJSON.Set("credit_card", creditCard)
	}
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

// GetCardsEndpoint handles get cards action
func GetCardsEndpoint(c *gin.Context) {
	id := c.Param("id")
	limitQuery := c.Query("limit")
	offsetQuery := c.Query("offset")

	var limit int64 = 10
	var offset int64

	if limitQuery != "" {
		limit1, _ := strconv.ParseInt(limitQuery, 10, 32)
		limit = limit1
	}
	if offsetQuery != "" {
		offset1, _ := strconv.ParseInt(offsetQuery, 0, 32)
		offset = offset1
	}

	cards := GetCards(id, int(limit), int(offset))

	if len(cards) > 0 {
		c.JSON(http.StatusOK, cards)
	} else {
		c.String(http.StatusOK, "There's no saved card for that ID")
	}
}

// SaveCardsEndpoint handles save card
func SaveCardsEndpoint(c *gin.Context) {
	id := c.Param("id")
	var arrayCard []Card
	requestBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.String(http.StatusBadRequest, "Wrong data format")
	} else {
		err1 := json.Unmarshal(requestBody, &arrayCard)
		if err1 != nil {
			c.String(http.StatusBadRequest, "Wrong data format")
		} else {
			err2 := SaveCards(id, arrayCard)
			if err2 != nil {
				c.String(http.StatusInternalServerError, err2.Error())
			} else {
				c.String(http.StatusOK, "Card is saved")
			}
		}
	}
}

func getInstallmentData() Installment {
	return Installment{Required: false, Terms: Terms{BNI: []int{3, 6, 12}, Mandiri: []int{3, 6, 12}, BCA: []int{3, 6, 12}, CIMB: []int{3, 6, 12}, Offline: []int{3, 6, 12}}}
}

func getWhitelistBin() []string {
	return []string{"481111", "521111"}
}
