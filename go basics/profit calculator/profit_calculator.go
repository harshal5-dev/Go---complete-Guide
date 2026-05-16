package main

import "fmt"

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Print("Enter revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Enter expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Enter tax rate: ")
	fmt.Scan(&taxRate)

	fmt.Println("Profit Calculator")

	earningsBeforeTax := revenue - expenses
	profit := earningsBeforeTax * (1 - taxRate/100)
	ratio := profit / revenue
	fmt.Println("Earnings before tax: ", earningsBeforeTax)
	fmt.Println("Profit: ", profit)
	fmt.Println("Ratio: ", ratio)

}
