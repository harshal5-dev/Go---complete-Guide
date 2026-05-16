package main

import (
	"fmt"
	"math"
)

const inflationRate = 6.5

func main() {

	var investmentAmount float64
	var years float64
	expectedReturnRate := 5.5

	outputText("Investment Amount: ")
	fmt.Scan(&investmentAmount)
	outputText("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)
	outputText("Years: ")
	fmt.Scan(&years)

	// futureValue := investmentAmount * math.Pow((1+expectedReturnRate/100), years)
	// futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)
	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)

	formattedFv := fmt.Sprintf("Future Value: %.1f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future Real Value: %.1f\n", futureRealValue)

	// fmt.Println("Future Value: ", futureValue)
	// fmt.Println("Future Real Value: ", futureRealValue)
	// fmt.Printf("Future Value: %v\nFuture Real Value: %v\n", futureValue, futureRealValue)
	// fmt.Printf("Future Value: %.1f\nFuture Real Value: %.1f\n", futureValue, futureRealValue)
	// fmt.Print(formattedFv, formattedRFV)
	// fmt.Printf(`Future Value: %.1f
	// 	Future Real Value: %.1f`, futureValue, futureRealValue)
	output(formattedFv, formattedRFV)
}

func outputText(text string) {
	fmt.Print(text)
}

func output(formattedFv string, formattedRFV string) {
	fmt.Print(formattedFv, formattedRFV)
}

func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (float64, float64) {
	futureValue := investmentAmount * math.Pow((1+expectedReturnRate/100), years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)
	return futureValue, futureRealValue
}

// func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (futureValue float64, futureRealValue float64) {
// 	futureValue = investmentAmount * math.Pow((1+expectedReturnRate/100), years)
// 	futureRealValue = futureValue / math.Pow(1+inflationRate/100, years)
// 	return
// }
