package main

import (
	"fmt"

	"example.com/bank/fileOps"
)

const accountBalanceFile = "balance.txt"

func main() {
	accountBalance, err := fileOps.GetFloatFromFile(accountBalanceFile)

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("-----------------")
		// panic("Can't continue Sorry!")
		return
	}

	fmt.Println("Welcome to Go Banking Application")

	for {

		presentOptions()
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
			fileOps.WriteFloatToFile(accountBalance, accountBalanceFile)
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
			fileOps.WriteFloatToFile(accountBalance, accountBalanceFile)
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
