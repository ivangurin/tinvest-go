package utils

import (
	contractv1 "tinvest-go/internal/pb"
)

const nano = 1000000000

func QuotationToFloat64(quotation *contractv1.Quotation) float64 {
	return float64(quotation.GetUnits()) + float64(quotation.GetNano())/nano
}

func MoneyToFloat64(money *contractv1.MoneyValue) float64 {
	return float64(money.GetUnits()) + float64(money.GetNano())/nano
}
