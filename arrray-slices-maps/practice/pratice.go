package main

import "fmt"

type Product struct {
	Title string
	Id    int
	Price float64
}

func New(id int, title string, price float64) *Product {
	return &Product{
		Id:    id,
		Title: title,
		Price: price,
	}
}

func main() {
	hobbies := [3]string{"reading", "hiking", "coding"}
	fmt.Println("My hobbies:", hobbies)
	fmt.Println("First Element of Array:", hobbies[0])
	fmt.Println("Second and Third Elements of Array:", hobbies[1:3])
	hobbiesSlice := hobbies[:2]
	// hobbiesSlice := hobbies[0:2]
	fmt.Println("hobbiesSlice: ", hobbiesSlice)
	hobbiesSlice = hobbiesSlice[1:3]
	fmt.Println("hobbiesSlice: ", hobbiesSlice)

	fmt.Println("-------------------")

	goals := []string{"learn go", "learn DS"}
	fmt.Println("Goals:", goals)
	goals[1] = "learn javascript"
	goals = append(goals, "learn system design")
	fmt.Println("Goals:", goals)

	fmt.Println("-------------------")
	products := []Product{*New(1, "Computer", 1500.0), {Id: 2, Title: "Mouse", Price: 500.0}}
	fmt.Println("Products:", products)
	newProduct := *New(3, "Keyboard", 800.0)
	products = append(products, newProduct)
	fmt.Println("Products:", products)
}
