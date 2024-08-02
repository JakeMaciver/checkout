package checkout_test

import (
	"testing"

	"github.com/JakeMaciver/checkout/pricing"
)

// test Scan
func TestScan(t *testing.T) {

	prices := map[string]pricing.ItemPricing{
		"A": {UnitPrice: 50, SpecialQty: 3, SpecialPrice: 130},
		"B": {UnitPrice: 30, SpecialQty: 2, SpecialPrice: 45},
		"C": {UnitPrice: 20, SpecialQty: 0, SpecialPrice: 0},
		"D": {UnitPrice: 15, SpecialQty: 0, SpecialPrice: 0},
	} 
}
// test GetTotal