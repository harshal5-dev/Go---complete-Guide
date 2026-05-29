package main

import (
	"fmt"

	// "example.com/price-calculator/cmdmanage"
	"example.com/price-calculator/filemanager"
	"example.com/price-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}
	doneChans := make([]chan bool, len(taxRates))
	errChnas := make([]chan error, len(taxRates))

	for index, taxRate := range taxRates {
		doneChans[index] = make(chan bool)
		errChnas[index] = make(chan error)
		outputFilePath := fmt.Sprintf("result_%0.1f.json", taxRate*100)
		fm := filemanager.New("prices.txt", outputFilePath)
		// cmd := cmdmanage.New()
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		// priceJob1 := prices.NewTaxIncludedPriceJob(cmd, taxRate)
		// err := priceJob.Process()
		// priceJob1.Process()
		go priceJob.Process(doneChans[index], errChnas[index])

		// if err != nil {
		// 	fmt.Println("Could not process job")
		// 	fmt.Println(err)
		// }

	}

	for index, _ := range taxRates {
		select {
		case err := <-errChnas[index]:
			if err != nil {
				fmt.Println(err)
			}

		case <-doneChans[index]:
			fmt.Println("Done!")
		}
	}

	// for _, errorChan := range errChnas {
	// 	<-errorChan
	// }

	// for _, doneChan := range doneChans {
	// 	<-doneChan
	// }
}
