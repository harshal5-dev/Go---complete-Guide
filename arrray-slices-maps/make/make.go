package main

import "fmt"

type flatMap map[string]float64

func (cr flatMap) display() {
	fmt.Println(cr)
}

func main() {
	userNames := make([]string, 2, 5)

	userNames[0] = "Harshal"

	userNames = append(userNames, "shraddha")
	userNames = append(userNames, "ganbote")

	fmt.Println(userNames)

	courseRatings := make(flatMap, 3)

	courseRatings["go"] = 4.6
	courseRatings["react"] = 4.7
	courseRatings["angular"] = 4.5

	courseRatings.display()

	for index, value := range userNames {
		fmt.Println("Index:", index)
		fmt.Println("Value:", value)
	}

	for key, value := range courseRatings {
		fmt.Println("Key:", key)
		fmt.Println("Value:", value)
	}
}
