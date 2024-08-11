package interest

/*
- 3.213% for a balance less than `0` dollars (balance gets more negative).
- 0.5% for a balance greater than or equal to `0` dollars, and less than `1000` dollars.
- 1.621% for a balance greater than or equal to `1000` dollars, and less than `5000` dollars.
- 2.475% for a balance greater than or equal to `5000` dollars.
*/

const (
	NegativeInterestMod float32 = 3.213
	NormalInterestMod   float32 = 0.5
	HighInterestMod     float32 = 1.621
	VeryHighInterestMod float32 = 2.475
)

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
	switch {
	case balance < 0:
		return NegativeInterestMod
	case balance < 1000:
		return NormalInterestMod
	case balance < 5000:
		return HighInterestMod
	default:
		return VeryHighInterestMod
	}
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	rate := InterestRate(balance)
	return balance / 100 * float64(rate)
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	return balance + Interest(balance)
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance.
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	futureBalance := balance
	years := 0

	for ; futureBalance < targetBalance; years++ {
		futureBalance = AnnualBalanceUpdate(futureBalance)
	}

	return years
}
