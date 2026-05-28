package prices

import (
	"fmt"

	"example.com/price-calculator/conversion"
	"example.com/price-calculator/filemanager"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]string
}

func (job *TaxIncludedPriceJob) LoadData() {
	lines, err := filemanager.ReadLines("prices.txt")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	prices, err := conversion.StringToFloat(lines)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	job.InputPrices = prices
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		InputPrices: []float64{10, 20, 30, 40},
		TaxRate:     taxRate,
	}
}

func (job *TaxIncludedPriceJob) Process() {
	job.LoadData()
	result := make(map[string]string)
	key := fmt.Sprintf("%0.2f", job.TaxRate)

	for _, price := range job.InputPrices {
		ans := price * (1 + job.TaxRate)
		result[key] = fmt.Sprintf("%0.2f", ans)
	}

	job.TaxIncludedPrices = result
	filemanager.WriteJSON(fmt.Sprintf("result_%0.1f.json", job.TaxRate*100), job)
}
