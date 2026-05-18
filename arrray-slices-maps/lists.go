package main

import "fmt"

// type Product struct {
// 	title string
// 	id    string
// 	price float64
// }

func dynamicSlices() {
	prices := []float64{10.99, 8.99}
	fmt.Println(prices[0:1])

	prices = append(prices, 5.99)
	fmt.Println(prices)
}

func main() {
	var productNames [4]string = [4]string{"A Book"}
	prices := [4]float64{10.99, 9.99, 45.99, 20.0}
	productNames[2] = "A Carpet"

	fmt.Println("prices:", prices)
	fmt.Println("product names:", productNames)

	fmt.Println(prices[2])

	fmt.Println("--------------")

	// featuredPrices := prices[1:3]
	// featuredPrices := prices[:3]
	featuredPrices := prices[1:]
	featuredPrices[0] = 199.99
	fmt.Println("featured prices:", featuredPrices)
	fmt.Println("len, cap:", len(featuredPrices), cap(featuredPrices))
	fmt.Println("prices:", prices)
	fmt.Println("-----------")

	highlightedPrices := featuredPrices[:1]
	fmt.Println("highlighted prices:", highlightedPrices)
	fmt.Println("len, cap:", len(highlightedPrices), cap(highlightedPrices))
	fmt.Println("-------------")

	highlightedPrices = highlightedPrices[:3]
	fmt.Println("highlighted prices:", highlightedPrices)
	fmt.Println("len, cap:", len(highlightedPrices), cap(highlightedPrices))
	fmt.Println("-------------------------------------------")
	dynamicSlices()
}
