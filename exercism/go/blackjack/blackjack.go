package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "one":
		return 1
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
	case "ten":
		return 10
	case "jack":
		return 10
	case "queen":
		return 10
	case "king":
		return 10
	case "ace":
		return 11
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	c1 := ParseCard(card1)
	c2 := ParseCard(card2)
	cs := c1 + c2
	d := ParseCard(dealerCard)

	switch {
	case cs == 22:
		return "P"
	case cs == 21 && d < 10:
		return "W"
	case 17 <= cs && cs <= 20:
		return "S"
	case 12 <= cs && cs <= 16 && d < 7:
		return "S"
	case 12 <= cs && cs <= 16 && d >= 7:
		return "H"
	case cs <= 11:
		return "H"
	default:
		return "S"
	}
}
