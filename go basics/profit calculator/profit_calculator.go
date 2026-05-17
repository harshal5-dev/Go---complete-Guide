package main

import (
	"errors"
	"fmt"
	"os"
)

func writeResultInFile(earningsBeforeTax, profit, ratio float64) {
	outputText := fmt.Sprintf("Earnings before tax: %.1f\n", earningsBeforeTax)
	outputText += fmt.Sprintf("Profit: %.1f\n", profit)
	outputText += fmt.Sprintf("Ratio:%.2f\n ", ratio)
	os.WriteFile("result.txt", []byte(outputText), 0644)
}

func main() {

	fmt.Println("Profit Calculator")

	revenue, err := getUserInput("Enter revenue: ")

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		return
	}

	expenses, err := getUserInput("Enter expenses: ")

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		return
	}

	taxRate, err := getUserInput("Enter tax rate: ")

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		return
	}

	earningsBeforeTax, profit, ratio := calculateValues(revenue, expenses, taxRate)

	fmt.Printf("Earnings before tax: %.1f\n", earningsBeforeTax)
	fmt.Printf("Profit: %.1f\n", profit)
	fmt.Printf("Ratio:%.2f\n ", ratio)
	writeResultInFile(earningsBeforeTax, profit, ratio)
}

func calculateValues(revenue, expenses, taxRate float64) (float64, float64, float64) {
	earningsBeforeTax := revenue - expenses
	profit := earningsBeforeTax * (1 - taxRate/100)
	ratio := profit / revenue
	return earningsBeforeTax, profit, ratio
}

func getUserInput(text string) (float64, error) {
	var userInput float64
	fmt.Print(text)
	fmt.Scan(&userInput)

	if userInput <= 0 {
		outputError := fmt.Sprintf("Invalid Input %.1f, please enter value greater than 0", userInput)
		return 0, errors.New(outputError)
	}

	return userInput, nil
}
