package checkout_test

import (
	"testing"

	"github.com/JakeMaciver/checkout/checkout"
	"github.com/JakeMaciver/checkout/pricing"
)

// test Scan
func TestScan(t *testing.T) {

	prices := map[string]pricing.ItemPricing{
		"A": {NormalPrice: 50, SpecialQty: 3, SpecialPrice: 130},
		"B": {NormalPrice: 30, SpecialQty: 2, SpecialPrice: 45},
		"C": {NormalPrice: 20, SpecialQty: 0, SpecialPrice: 0},
		"D": {NormalPrice: 15, SpecialQty: 0, SpecialPrice: 0},
	}

	catalogue := pricing.NewCatalogue(prices)
	checkout := checkout.NewCheckout(catalogue)

}
// test GetTotal