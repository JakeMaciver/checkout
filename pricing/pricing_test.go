package pricing_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/JakeMaciver/checkout/pricing"
)

// Helper function to create a new catalogue with prices
func newCatalogue() *pricing.Catalogue {
	return pricing.NewCatalogue(map[string]pricing.ItemPricing{
		"A": {NormalPrice: 50, SpecialQty: 3, SpecialPrice: 130},
	})
}

// Test for adding an item to the catalogue
func TestAddItem(t *testing.T) {
	catalogue := newCatalogue()

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
	catalogue := newCatalogue()

	// testing for a positive run through the method resulting in no errors and expected behaviour of updating an item
	t.Run("positive case", func(t *testing.T) {
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
	})

	// testing the validation of the SKU input
	t.Run("error case, invalid SKU", func(t *testing.T) {
		SKUtoUpdate := "4"
		newItem := pricing.ItemPricing{
			NormalPrice:  20,
			SpecialQty:   3,
			SpecialPrice: 40,
		}

		got := catalogue.UpdateItem(SKUtoUpdate, newItem.NormalPrice, newItem.SpecialQty, newItem.SpecialPrice)
		want := errors.New("invalid SKU: 4")

		if got.Error() != want.Error() {
			t.Errorf("got: %v, want: %v", got, want)
		}

		SKUtoUpdate = ""
		gotNoInput := catalogue.UpdateItem(SKUtoUpdate, newItem.NormalPrice, newItem.SpecialQty, newItem.SpecialPrice)
		wantNoInput := errors.New("invalid SKU: ")

		if gotNoInput.Error() != wantNoInput.Error() {
			t.Errorf("got: %v, want: %v", gotNoInput, wantNoInput)
		}
	})

	// testing the validation of the normalPrice input
	t.Run("error case, invalid normal price", func(t *testing.T) {
		SKUtoUpdate := "C"
		newItem := pricing.ItemPricing{
			NormalPrice:  0,
			SpecialQty:   3,
			SpecialPrice: 40,
		}

		got := catalogue.UpdateItem(SKUtoUpdate, newItem.NormalPrice, newItem.SpecialQty, newItem.SpecialPrice)
		want := errors.New("invalid normal price: 0")

		if got.Error() != want.Error() {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	// testing is the user enters 0 in the specialQty
	t.Run("positive case, switching SpecialPrice based on SpecialQty", func(t *testing.T) {
		SKUtoUpdate := "A"
		newItem := pricing.ItemPricing{
			NormalPrice:  15,
			SpecialQty:   0,
			SpecialPrice: 40,
		}

		catalogue.UpdateItem(SKUtoUpdate, newItem.NormalPrice, newItem.SpecialQty, newItem.SpecialPrice)

		got := catalogue.Prices[SKUtoUpdate]
		want := pricing.ItemPricing{
			NormalPrice:  15,
			SpecialQty:   0,
			SpecialPrice: 0,
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	// testing if the item doesnt exist in the prices map
	t.Run("error case, not found", func(t *testing.T) {
		SKUtoUpdate := "B"
		newItem := pricing.ItemPricing{
			NormalPrice:  15,
			SpecialQty:   0,
			SpecialPrice: 40,
		}

		got := catalogue.UpdateItem(SKUtoUpdate, newItem.NormalPrice, newItem.SpecialQty, newItem.SpecialPrice)
		want := errors.New("item doesnt exist: B")

		if got.Error() != want.Error() {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

// Test for deleting an item in the catalogue
func TestDeleteItem(t *testing.T) {
	catalogue := newCatalogue()

	// test the deleting of an item in the prices map, positive pass
	t.Run("positive case", func(t *testing.T) {
		SKUtoDelete := "A"

		catalogue.DeleteItem(SKUtoDelete)

		_, got := catalogue.Prices[SKUtoDelete]
		want := false

		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	// testing the validation of the SKU input
	t.Run("error case, invalid SKU", func(t *testing.T) {
		// no input
		SKUtoDelete := ""

		got := catalogue.DeleteItem(SKUtoDelete)
		want := errors.New("invalid SKU: ")

		if got.Error() != want.Error() {
			t.Errorf("got: %v, want: %v", got, want)
		}

		// not an uppercase letter
		SKUtoDelete = "6"

		gotletter := catalogue.DeleteItem(SKUtoDelete)
		wantletter := errors.New("invalid SKU: ")

		if got.Error() != want.Error() {
			t.Errorf("got: %v, want: %v", gotletter, wantletter)
		}
	})

	// testing if the item does exist in the prices map
	t.Run("error case, item doesnt exist", func(t *testing.T) {
		SKUtoDelete := "B"

		got := catalogue.DeleteItem(SKUtoDelete)
		want := errors.New("item doesnt exist: B")

		if got.Error() != want.Error() {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}
