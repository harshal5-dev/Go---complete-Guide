package main

import "fmt"

func main() {
	var accountBalance = 1000.0

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

		if choice == 1 {
			fmt.Println("Your account balance is", accountBalance)
		} else if choice == 2 {
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
		} else if choice == 3 {
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
		} else {
			fmt.Println("Goodbye..!")
			break
		}
	}

	fmt.Println("Thank you for choosing our bank")
}
