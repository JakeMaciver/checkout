package checkout_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/JakeMaciver/checkout/checkout"
	"github.com/JakeMaciver/checkout/pricing"
)

// test Scan method on the checkout struct
func TestScan(t *testing.T) {

	prices := map[string]pricing.ItemPricing{
		"A": {NormalPrice: 50, SpecialQty: 3, SpecialPrice: 130},
		"B": {NormalPrice: 30, SpecialQty: 2, SpecialPrice: 45},
		"C": {NormalPrice: 20, SpecialQty: 0, SpecialPrice: 0},
		"D": {NormalPrice: 15, SpecialQty: 0, SpecialPrice: 0},
	}

	catalogue := pricing.NewCatalogue(prices)
	checkout := checkout.NewCheckout(*catalogue)

	// Runs through the Scan func with no errors
	t.Run("postive case", func(t *testing.T) {
		itemsToScan := []string{"A", "B", "C", "A", "B", "D"}

		for _, item := range itemsToScan {
			checkout.Scan(item)
		}

		got := checkout.Items
		want := map[string]int{
			"A": 2,
			"B": 2,
			"C": 1,
			"D": 1,
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	// Runs through scan and tests if we recieve an error when the input is equal to nothing
	t.Run("error case, input nonexistant", func(t *testing.T) {
		itemsToScan := []string{""}

		var err error
		for _, item := range itemsToScan {
			err = checkout.Scan(item)
		}

		got := err
		want := errors.New("no item to scan")

		if got.Error() != want.Error() {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	// runs through Scan and tests is we recieve an error when the input is not found in the catalogue
	t.Run("error case, not found", func(t *testing.T) {
		itemsToScan := []string{"F"}

		var err error
		for _, item := range itemsToScan {
			err = checkout.Scan(item)
		}

		got := err
		want := errors.New("item not found in the catalogue: F")

		if got.Error() != want.Error() {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

// test GetTotalPrice method on the Checkout struct
func TestGetTotalPrice(t *testing.T) {

	prices := map[string]pricing.ItemPricing{
		"A": {NormalPrice: 50, SpecialQty: 3, SpecialPrice: 130},
		"B": {NormalPrice: 30, SpecialQty: 2, SpecialPrice: 45},
		"C": {NormalPrice: 20, SpecialQty: 0, SpecialPrice: 0},
		"D": {NormalPrice: 15, SpecialQty: 0, SpecialPrice: 0},
	}

	// Runs through the GetTotalPrice method with no errors	
	t.Run("positive case", func(t *testing.T) {	
		catalogue := pricing.NewCatalogue(prices)
		checkout := checkout.NewCheckout(*catalogue)
	
		itemsToScan := []string{"A", "B", "C", "D", "A", "B"}
	
		for _, item := range itemsToScan {
			_ = checkout.Scan(item)
		}
	
		got, _ := checkout.GetTotalPrice()
		want := 180
	
		if got != want {
			t.Errorf("got: %v, wanted: %v", got, want)
		}
	})

	// Runs throught the GetTotalPrice method with no items in Checkout to check for error
	t.Run("error case, invlid parameter", func(t *testing.T) {	
		catalogue := pricing.NewCatalogue(prices)
		checkout := checkout.NewCheckout(*catalogue)	
		
		_, got := checkout.GetTotalPrice()
		want := errors.New("you have not scanned any items yet")

		if got.Error() != want.Error() {
			t.Errorf("got: %v, want: %v", got , want)
		}
	})
}
