package main

import "fmt"

type transFormFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4}

	transformValue := transformNumbers(&numbers, func(val int) int {
		return val * 2
	})

	fmt.Println("transformValue:", transformValue)

	doubleFunc := createTransformFunc(2)
	tripleFunc := createTransformFunc(3)

	double := transformNumbers(&numbers, doubleFunc)
	triple := transformNumbers(&numbers, tripleFunc)

	fmt.Println("double:", double)
	fmt.Println("triple:", triple)
}

func transformNumbers(numbers *[]int, transform transFormFn) []int {
	tNumber := []int{}

	for _, val := range *numbers {
		tNumber = append(tNumber, transform(val))
	}

	return tNumber
}

func createTransformFunc(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}
