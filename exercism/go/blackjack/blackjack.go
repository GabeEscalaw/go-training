package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
		default: 
			return 0
		case "ace": 
			return 11
		case "two": 
			return 2
		case "three": 
			return 3
		case "four": 
			return 4
		case "five": 
			return 5
		case "six": 
			return 6
		case "seven": 
			return 7
		case "eight": 
			return 8
		case "nine": 
			return 9
		case "ten", "jack", "queen", "king": 
			return 10
		}
}

// IsBlackjack returns true if the player has a blackjack, false otherwise.
func IsBlackjack(card1, card2 string) bool {
	return ParseCard(card1) + ParseCard(card2) == 21
}

// LargeHand implements the decision tree for hand scores larger than 20 points.
func LargeHand(isBlackjack bool, dealerScore int) string {
	if isBlackjack {
		if dealerScore != 10 && dealerScore != 11 {
			return "W"
		} else {
			return "S"
		}
	} else {
		return "P"
	}
}

// SmallHand implements the decision tree for hand scores with less than 21 points.
func SmallHand(handScore, dealerScore int) string {
	if handScore < 21 {
		if handScore >= 17 {
			return "S"
		} else if handScore <= 11 {
			return "H"
		} else if handScore > 11 && handScore < 17 && dealerScore <= 6 {
			return "S"
		} else if handScore > 11 && handScore < 17 && dealerScore > 6 {
			return "H"
		} else {
			return "S"
		}
	}

	return ""
}
