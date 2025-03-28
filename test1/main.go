package main

import (
	"fmt"
	"math"
)

func main() {
	var investmentAmount float64
	var years float64
	var expectedReturnRate float64

	// Ask the user for inputs
	fmt.Print("Enter the investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Enter the number of years: ")
	fmt.Scan(&years)

	fmt.Print("Enter the expected return rate (in %): ")
	fmt.Scan(&expectedReturnRate)

	// Calculate the future value
	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	fmt.Printf("The future value of the investment is: %.2f\n", futureValue)
}
