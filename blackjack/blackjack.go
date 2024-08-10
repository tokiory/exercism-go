package blackjack

/*
| card  | value | card    | value |
| :---: | :---: | :-----: | :---: |
|  ace  |  11   | eight   |   8   |
|  two  |   2   | nine    |   9   |
| three |   3   |  ten    |  10   |
| four  |   4   | jack    |  10   |
| five  |   5   | queen   |  10   |
|  six  |   6   | king    |  10   |
| seven |   7   | *other* |   0   |
*/

const (
	ActionStand            = "S"
	ActionHit              = "H"
	ActionSplit            = "P"
	ActionAutomaticallyWin = "W"
)

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
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
	case "ace":
		return 11
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	ownHandSum := ParseCard(card1) + ParseCard(card2)
	dealerSum := ParseCard(dealerCard)

	switch {
	case ownHandSum == 22:
		return ActionSplit
	case ownHandSum == 21:
		if dealerSum == 10 || dealerSum == 11 {
			return ActionStand
		} else {
			return ActionAutomaticallyWin
		}
	case ownHandSum >= 17 && ownHandSum <= 20:
		return ActionStand
	case ownHandSum >= 12 && ownHandSum <= 16:
		if dealerSum >= 7 {
			return ActionHit
		}
		return ActionStand
	default:
		return ActionHit
	}
}
