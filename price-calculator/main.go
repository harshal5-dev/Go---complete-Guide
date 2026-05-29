package main

import (
	"fmt"

	// "example.com/price-calculator/cmdmanage"
	"example.com/price-calculator/filemanager"
	"example.com/price-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRates {
		outputFilePath := fmt.Sprintf("result_%0.1f.json", taxRate*100)
		fm := filemanager.New("prices.txt", outputFilePath)
		// cmd := cmdmanage.New()
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		// priceJob1 := prices.NewTaxIncludedPriceJob(cmd, taxRate)
		err := priceJob.Process()
		// priceJob1.Process()

		if err != nil {
			fmt.Println("Could not process job")
			fmt.Println(err)
		}

	}

}
