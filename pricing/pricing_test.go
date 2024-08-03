package pricing_test

import (
	"reflect"
	"testing"

	"github.com/JakeMaciver/checkout/pricing"
)

// Test for adding an item to the catalogue
func TestAddItem(t *testing.T) {
	prices := map[string]pricing.ItemPricing{
		"A": {NormalPrice: 50, SpecialQty: 3, SpecialPrice: 130},
	}

	catalogue := pricing.NewCatalogue(prices)

	SKUtoAdd := "B"
	newItem := pricing.ItemPricing{
		NormalPrice: 20,
		SpecialQty: 3,
		SpecialPrice: 40,
	}

	catalogue.AddItem(SKUtoAdd, newItem.NormalPrice, newItem.SpecialQty, newItem.SpecialPrice)

	_, got := catalogue.Prices[SKUtoAdd]
	want := true

	if got != want {
		t.Errorf("got: %v, want: %v", got, want)
	}

	gotItem := catalogue.Prices[SKUtoAdd]
	wantItem := newItem

	if !reflect.DeepEqual(gotItem, wantItem) {
		t.Errorf("got: %v, want: %v", gotItem, wantItem)
	}
}

// Test for updating an item in the catalogue
// Test for deleting an item in the catalogue
