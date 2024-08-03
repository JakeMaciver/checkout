package pricing_test

import (
	"testing"

	"github.com/JakeMaciver/checkout/pricing"
)

// Test for adding an item to the catalogue
func TestAddItem(t *testing.T) {
	prices := map[string]pricing.ItemPricing{
		"A": {NormalPrice: 50, SpecialQty: 3, SpecialPrice: 130},
	}

	catalogue := pricing.NewCatalogue(prices)

	
}

// Test for updating an item in the catalogue
// Test for deleting an item in the catalogue
