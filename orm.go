package main

import (
	"fmt"

	"gopkg.in/mgo.v2/bson"
)

const (
	cardCollectionKey = "cards"
)

// GetCards to get cards from database
func GetCards(userID string, limit int, skip int) []JSONCard {
	var cards []JSONCard
	c := MongoDB.C(cardCollectionKey)
	err := c.Find(bson.M{"userid": userID}).Limit(limit).Skip(skip).All(&cards)
	if err != nil {
		fmt.Println("Error when getting cards: " + err.Error())
	}
	return cards
}

// SaveCards to save list of cards into database
func SaveCards(userID string, cards []Card) error {
	var savedCards []Card
	c := MongoDB.C(cardCollectionKey)
	err := c.Find(bson.M{"userid": userID}).All(&savedCards)
	if err != nil {
		fmt.Println("Error when getting cards: " + err.Error())
	}

	if len(savedCards) > 0 {
		_, err1 := c.RemoveAll(bson.M{"userid": userID})
		if err1 != nil {
			return err
		}
	}

	for _, card := range cards {
		card.UserID = userID
		err2 := c.Insert(&card)
		if err2 != nil {
			return err2
		}
	}
	return nil
}
