package main

import "fmt"

func main() {
	numbers := []int{1, 10, 15}
	sum := sumUp(1, 4, 6, 7, 4, 5)
	anotherSum := sumUp(11, numbers...)

	fmt.Println(sum)
	fmt.Println(anotherSum)
}

func sumUp(startingValue int, numbers ...int) int {
	sum := 0

	fmt.Println("startingValue:", startingValue)

	for _, val := range numbers {
		sum += val
	}

	return sum
}
