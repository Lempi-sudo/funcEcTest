package ecofunc

// PresentValue calculates the present value of a future amount
func PresentValue(futureValue float64, rate float64, periods int) float64 {
	return futureValue / (1 + rate)
}

// PresentValue calculates the present value of a future amount
func Version() string {
	return "3.0"
}
