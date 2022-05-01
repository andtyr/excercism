package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	var Anyvarname int = 0
	switch card {
	case "ace":
		Anyvarname = 11
	case "king", "queen", "jack", "ten":
		Anyvarname = 10
	case "nine":
		Anyvarname = 9
	case "eight":
		Anyvarname = 8
	case "seven":
		Anyvarname = 7
	case "six":
		Anyvarname = 6
	case "five":
		Anyvarname = 5
	case "four":
		Anyvarname = 4
	case "three":
		Anyvarname = 3
	case "two":
		Anyvarname = 2
	case "one":
		Anyvarname = 1
	default:
		return 0
	}
	return Anyvarname
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	playerAction := ""
	playerCardSum := ParseCard(card1) + ParseCard(card2)
	dealerCardSum := ParseCard(dealerCard)

	switch {
	case card1 == "ace" && card2 == "ace":
		playerAction = "P"
	case playerCardSum == 21 && dealerCardSum < 10:
		playerAction = "W"
	case playerCardSum >= 17 && playerCardSum <= 21 || (playerCardSum >= 12 && playerCardSum <= 16 && dealerCardSum < 7):
		playerAction = "S"
	case playerCardSum >= 12 && playerCardSum <= 16 && dealerCardSum >= 7 || playerCardSum <= 11:
		playerAction = "H"
	}
	return playerAction

}
