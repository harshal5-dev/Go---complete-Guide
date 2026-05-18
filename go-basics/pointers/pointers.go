package main

import "fmt"

func main() {
	age := 32

	var agePointer *int
	agePointer = &age

	fmt.Println("Age:", age)
	fmt.Println("Age Pointer address and value", agePointer, *agePointer)

	// adultYears := getAdultYears(agePointer)
	editAgeToAdultYears(agePointer)
	// fmt.Println(adultYears)
	fmt.Println(age)
}

func editAgeToAdultYears(agePointer *int) {
	// return *agePointer - 18
	*agePointer = *agePointer - 18
}
