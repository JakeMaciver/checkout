package pricing_test

import (
	"errors"
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

	// positive run through AddItem 
	t.Run("positive case", func(t *testing.T) {

		SKUtoAdd := "B"
		newItem := pricing.ItemPricing{
			NormalPrice:  20,
			SpecialQty:   3,
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
	})

	// testing the validation of the SKU input
	t.Run("error case, invalid SKU", func(t *testing.T) {

		SKUtoAdd := "4"
		newItem := pricing.ItemPricing{
			NormalPrice:  20,
			SpecialQty:   3,
			SpecialPrice: 40,
		}

		got := catalogue.AddItem(SKUtoAdd, newItem.NormalPrice, newItem.SpecialQty, newItem.SpecialPrice)
		want := errors.New("invalid SKU: 4")

		if got.Error() != want.Error() {
			t.Errorf("got: %v, want: %v", got, want)
		}

		SKUtoAdd = ""
		gotNoInput := catalogue.AddItem(SKUtoAdd, newItem.NormalPrice, newItem.SpecialQty, newItem.SpecialPrice)
		wantNoInput := errors.New("invalid SKU: ")

		if gotNoInput.Error() != wantNoInput.Error() {
			t.Errorf("got: %v, want: %v", gotNoInput, wantNoInput)
		}
	})

	// testing the validation of the normalPrice input
	t.Run("error case, invalid normal price", func(t *testing.T) {

		SKUtoAdd := "C"
		newItem := pricing.ItemPricing{
			NormalPrice:  0,
			SpecialQty:   3,
			SpecialPrice: 40,
		}

		got := catalogue.AddItem(SKUtoAdd, newItem.NormalPrice, newItem.SpecialQty, newItem.SpecialPrice)
		want := errors.New("invalid normal price: 0")

		if got.Error() != want.Error() {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	// testing is the user enters 0 in the specialQty
	t.Run("positive case, switching SpecialPrice based on SpecialQty", func(t *testing.T) {

		SKUtoAdd := "C"
		newItem := pricing.ItemPricing{
			NormalPrice:  15,
			SpecialQty:   0,
			SpecialPrice: 40,
		}

		catalogue.AddItem(SKUtoAdd, newItem.NormalPrice, newItem.SpecialQty, newItem.SpecialPrice)

		got := catalogue.Prices[SKUtoAdd]
		want := pricing.ItemPricing{
			NormalPrice:  15,
			SpecialQty:   0,
			SpecialPrice: 0,
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	// testing if the item the user is trying to add already exists in the catalogue
	t.Run("error case, already exists", func(t *testing.T) {
		SKUtoAdd := "A"
		newItem := pricing.ItemPricing{
			NormalPrice:  20,
			SpecialQty:   3,
			SpecialPrice: 40,
		}

		got := catalogue.AddItem(SKUtoAdd, newItem.NormalPrice, newItem.SpecialQty, newItem.SpecialPrice)
		want := errors.New("item already exists: A")

		if got.Error() != want.Error() {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

// Test for updating an item in the catalogue
func TestUpdateItem(t *testing.T) {
	prices := map[string]pricing.ItemPricing{
		"A": {NormalPrice: 50, SpecialQty: 3, SpecialPrice: 130},
	}

	catalogue := pricing.NewCatalogue(prices)

	SKUtoUpdate := "A"
	newItem := pricing.ItemPricing{
		NormalPrice:  30,
		SpecialQty:   3,
		SpecialPrice: 40,
	}	

	catalogue.UpdateItem(SKUtoUpdate, newItem.NormalPrice, newItem.SpecialQty, newItem.SpecialPrice)

	got := catalogue.Prices[SKUtoUpdate]
	want := pricing.ItemPricing{
		NormalPrice:  30,
		SpecialQty:   3,
		SpecialPrice: 40,
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}

	// errors to handle
	// all input validation

	// testing the validation of the SKU input
	t.Run("error case, invalid SKU", func(t *testing.T) {

		SKUtoAdd := "4"
		newItem := pricing.ItemPricing{
			NormalPrice:  20,
			SpecialQty:   3,
			SpecialPrice: 40,
		}

		got := catalogue.UpdateItem(SKUtoAdd, newItem.NormalPrice, newItem.SpecialQty, newItem.SpecialPrice)
		want := errors.New("invalid SKU: 4")

		if got.Error() != want.Error() {
			t.Errorf("got: %v, want: %v", got, want)
		}

		SKUtoAdd = ""
		gotNoInput := catalogue.UpdateItem(SKUtoAdd, newItem.NormalPrice, newItem.SpecialQty, newItem.SpecialPrice)
		wantNoInput := errors.New("invalid SKU: ")

		if gotNoInput.Error() != wantNoInput.Error() {
			t.Errorf("got: %v, want: %v", gotNoInput, wantNoInput)
		}
	})

	// testing the validation of the normalPrice input
	t.Run("error case, invalid normal price", func(t *testing.T) {

		SKUtoAdd := "C"
		newItem := pricing.ItemPricing{
			NormalPrice:  0,
			SpecialQty:   3,
			SpecialPrice: 40,
		}

		got := catalogue.UpdateItem(SKUtoAdd, newItem.NormalPrice, newItem.SpecialQty, newItem.SpecialPrice)
		want := errors.New("invalid normal price: 0")

		if got.Error() != want.Error() {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	// testing is the user enters 0 in the specialQty
	t.Run("positive case, switching SpecialPrice based on SpecialQty", func(t *testing.T) {

		SKUtoAdd := "A"
		newItem := pricing.ItemPricing{
			NormalPrice:  15,
			SpecialQty:   0,
			SpecialPrice: 40,
		}

		catalogue.UpdateItem(SKUtoAdd, newItem.NormalPrice, newItem.SpecialQty, newItem.SpecialPrice)

		got := catalogue.Prices[SKUtoAdd]
		want := pricing.ItemPricing{
			NormalPrice:  15,
			SpecialQty:   0,
			SpecialPrice: 0,
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}
// Test for deleting an item in the catalogue
