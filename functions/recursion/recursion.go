package main

import "fmt"

func main() {

	fmt.Println(calculateFactorial(5))

}

func calculateFactorial(num int) int {
	if num == 1 {
		return num
	}

	return num * calculateFactorial(num-1)
}

// func calculateFactorial(num int) int {
// 	result := 1

// 	for val := 1; val <= num; val++ {
// 		result *= val
// 	}

// 	return result
// }
