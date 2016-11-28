package main

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func containsCard(cards []Card, card Card) bool {
	for _, a := range cards {
		if a.SavedToken == card.SavedToken {
			return true
		}
	}

	return false
}
