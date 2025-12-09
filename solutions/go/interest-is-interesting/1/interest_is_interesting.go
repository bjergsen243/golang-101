package interest

const RATE_0 = 3.213
const RATE_1000 = 0.5
const RATE_5000 = 1.621
const RATE_MORE_5000 = 2.475

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
	if balance < 0 {
		return RATE_0
	}
	if balance < 1000 {
		return RATE_1000
	}
	if balance < 5000 {
		return RATE_5000
	}
	return RATE_MORE_5000
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	return balance * float64(InterestRate(balance)) / 100
}

// AnnualBalanceUpdate calculates the annual balance update.
func AnnualBalanceUpdate(balance float64) float64 {
	return balance + Interest(balance)
}

// YearsBeforeDesiredBalance calculates how many years needed
// until balance >= targetBalance.
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	years := 0
	for balance < targetBalance {
		balance = AnnualBalanceUpdate(balance)
		years++
	}
	return years
}
