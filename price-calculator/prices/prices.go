package prices

import (
	"fmt"

	"example.com/price-calculator/conversion"
	"example.com/price-calculator/iomanager"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64             `json:"tax_rate"`
	InputPrices       []float64           `json:"input_prices"`
	TaxIncludedPrices map[string]string   `json:"tax_included_prices"`
	IOManager         iomanager.IOManager `json:"-"`
	// IOManage          filemanager.FileManager `json:"-"`
}

func (job *TaxIncludedPriceJob) LoadData() error {

	// lines, err := filemanager.ReadLines("prices.txt")
	lines, err := job.IOManager.ReadLines()

	if err != nil {
		// fmt.Println(err.Error())
		return err
	}

	prices, err := conversion.StringToFloat(lines)

	if err != nil {
		// fmt.Println(err.Error())
		return err
	}

	job.InputPrices = prices
	return nil
}

func (job *TaxIncludedPriceJob) Process() error {
	err := job.LoadData()

	if err != nil {
		return err
	}

	result := make(map[string]string)
	key := fmt.Sprintf("%0.2f", job.TaxRate)

	for _, price := range job.InputPrices {
		ans := price * (1 + job.TaxRate)
		result[key] = fmt.Sprintf("%0.2f", ans)
	}

	job.TaxIncludedPrices = result
	// filemanager.WriteJSON(fmt.Sprintf("result_%0.1f.json", job.TaxRate*100), job)
	return job.IOManager.WriteResult(job)
}

func NewTaxIncludedPriceJob(fm iomanager.IOManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		InputPrices: []float64{10, 20, 30, 40},
		TaxRate:     taxRate,
		IOManager:   fm,
	}
}
