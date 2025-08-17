package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
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
    return 0
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	score := ParseCard(card1) + ParseCard(card2)
    dealerScore := ParseCard(dealerCard)

    if score == 22 {
        return "P"
    }

    if score == 21 {
        if dealerScore < 10 {
            return "W"
        } else {
        	return "S"
        }
    }

    if score >= 17 {
        return "S"
    }

    if score >= 12 {
        if dealerScore >= 7 {
            return "H"
        } else {
        	return "S"
        }
    }

    return "H"
}
