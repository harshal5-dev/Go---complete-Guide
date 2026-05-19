package main

import "fmt"

type transFormFn func(int) int

// type anotherFn func(int, string, []int) int

func main() {
	numbers := []int{1, 2, 3, 4}
	otherNumbers := []int{3, 1, 5}

	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)

	fmt.Println(doubled)
	fmt.Println(tripled)

	fmt.Println("------------")

	tranFormNumberFunc1 := getTransformerFunction(&numbers)
	tranFormNumberFunc2 := getTransformerFunction(&otherNumbers)

	doubled1 := transformNumbers(&numbers, tranFormNumberFunc1)
	tripled2 := transformNumbers(&numbers, tranFormNumberFunc2)

	fmt.Println("doubled1:", doubled1)
	fmt.Println("tripled2:", tripled2)
}

func transformNumbers(numbers *[]int, transform transFormFn) []int {
	tNumber := []int{}

	for _, val := range *numbers {
		tNumber = append(tNumber, transform(val))
	}

	return tNumber
}

func getTransformerFunction(numbers *[]int) transFormFn {
	if (*numbers)[0] == 1 {
		return double
	} else {
		return triple
	}
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}
