package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(accountBalanceFile)

	if err != nil {
		return 1000, errors.New("Failed to find balance file.")
	}

	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)

	if err != nil {
		return 1000, errors.New("Failed to parse stored balance value.")
	}

	return balance, nil
}

func writeBalanceToFile(accountBalance float64) {
	balanceText := fmt.Sprint(accountBalance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func main() {
	accountBalance, err := getBalanceFromFile()

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("-----------------")
		// panic("Can't continue Sorry!")
		return
	}

	fmt.Println("Welcome to Go Banking Application")

	for {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		var choice int

		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your account balance is", accountBalance)
		case 2:
			var deposit float64
			fmt.Print("Your Deposit: ")
			fmt.Scan(&deposit)

			if deposit <= 0 {
				fmt.Println("Invalid Amount. Must be greater than 0.")
				// return
				continue
			}

			accountBalance += deposit
			fmt.Println("Account Balance updated, new balance is", accountBalance)
			writeBalanceToFile(accountBalance)
		case 3:
			var withdraw float64
			fmt.Print("Your Withdraw: ")
			fmt.Scan(&withdraw)

			if withdraw <= 0 {
				fmt.Println("Invalid Amount. Must be greater than 0.")
				// return
				continue
			}

			if withdraw > accountBalance {
				fmt.Println("Invalid Amount. you can't withdraw more than you have.")
				// return
				continue
			}

			accountBalance -= withdraw
			fmt.Println("Account Balance updated, new balance is", accountBalance)
			writeBalanceToFile(accountBalance)
		default:
			fmt.Println("Goodbye..!")
			fmt.Println("Thank you for choosing our bank")
			return
			// break
		}

		// if choice == 1 {
		// 	fmt.Println("Your account balance is", accountBalance)
		// } else if choice == 2 {
		// 	var deposit float64
		// 	fmt.Print("Your Deposit: ")
		// 	fmt.Scan(&deposit)

		// 	if deposit <= 0 {
		// 		fmt.Println("Invalid Amount. Must be greater than 0.")
		// 		// return
		// 		continue
		// 	}

		// 	accountBalance += deposit
		// 	fmt.Println("Account Balance updated, new balance is", accountBalance)
		// } else if choice == 3 {
		// 	var withdraw float64
		// 	fmt.Print("Your Withdraw: ")
		// 	fmt.Scan(&withdraw)

		// 	if withdraw <= 0 {
		// 		fmt.Println("Invalid Amount. Must be greater than 0.")
		// 		// return
		// 		continue
		// 	}

		// 	if withdraw > accountBalance {
		// 		fmt.Println("Invalid Amount. you can't withdraw more than you have.")
		// 		// return
		// 		continue
		// 	}

		// 	accountBalance -= withdraw
		// 	fmt.Println("Account Balance updated, new balance is", accountBalance)
		// } else {
		// 	fmt.Println("Goodbye..!")
		// 	break
		// }
	}
}
