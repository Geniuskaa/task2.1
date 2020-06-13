package credit

import "math"

func CalculateCredit(sumOfCredit, annualInterestRate, durationInYears float64) (int64, int64, int64) {
	durationInMonths := durationInYears * 12.0
	monthlyInterestRate := annualInterestRate / 12.0 / 100.0
	a := (1.0 + monthlyInterestRate)
	pow := math.Pow(a, durationInMonths)
	rate := (monthlyInterestRate * pow) / (pow - 1.0)
	monthlyPayment := rate * sumOfCredit / 100.0
	monthlyPaymentInRub := int64(math.Round(monthlyPayment))
	totalPayment := int64(durationInMonths) * monthlyPaymentInRub
	interestPayment := totalPayment - int64(sumOfCredit)
	return monthlyPaymentInRub, interestPayment, totalPayment
}
