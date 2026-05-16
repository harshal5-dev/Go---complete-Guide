package main

import "fmt"

func main() {

	revenue := getUserInput("Enter revenue: ")
	expenses := getUserInput("Enter expenses: ")
	taxRate := getUserInput("Enter tax rate: ")

	fmt.Println("Profit Calculator")

	earningsBeforeTax, profit, ratio := calculateValues(revenue, expenses, taxRate)

	fmt.Printf("Earnings before tax: %.1f\n", earningsBeforeTax)
	fmt.Printf("Profit: %.1f\n", profit)
	fmt.Printf("Ratio:%.2f\n ", ratio)
}

func calculateValues(revenue, expenses, taxRate float64) (float64, float64, float64) {
	earningsBeforeTax := revenue - expenses
	profit := earningsBeforeTax * (1 - taxRate/100)
	ratio := profit / revenue
	return earningsBeforeTax, profit, ratio
}

func getUserInput(text string) float64 {
	var userInput float64
	fmt.Print(text)
	fmt.Scan(&userInput)
	return userInput
}
