package checkout

import (
	"errors"
	"fmt"
	"sync"

	"github.com/JakeMaciver/checkout/pricing"
)

// ICheckout represents an interface for the checkout
type ICheckout interface {
	// Scan is method that will scan an item to be paid for
	Scan(SKU string) (err error)
	// GetTotalPrice is a method that will get the total price of all the items that have been scanned
	GetTotalPrice() (totalPrice int, err error)
}

// Checkout represents an implementation of the ICheckout interface
type Checkout struct {
	// Items is a store that contains the items that have been scanned and the amount of that item that has been scanned
	Items map[string]int
	// Catalogue is an index of items, a normal price for that item, a special price for that item and a special quantity for that item
	Catalogue pricing.Catalogue
}

// NewCheckout initialises and returns an instance of Checkout
func NewCheckout(catalogue pricing.Catalogue) *Checkout {
	return &Checkout{
		Items:     make(map[string]int),
		Catalogue: catalogue,
	}
}

// Scan will add the SKU into an items list or increment if the item is already in the list, errors occur of invalid input and not found
func (c *Checkout) Scan(SKU string) error {
	// invalid input
	if err := pricing.ValidateSKU(SKU); err != nil {
		return err
	}

	// not found in catalogue
	if _, exists := c.Catalogue.Prices[SKU]; !exists {
		return fmt.Errorf("item not found in the catalogue: %s", SKU)
	}

	c.Items[SKU]++
	return nil
}

// GetTotal will add up the price of all the items in the c.Items map, takes into account if an item reaches the special price requirement. errors occur if not items available
func (c *Checkout) GetTotalPrice() (int, error) {
	if len(c.Items) == 0 {
		return 0, errors.New("you have not scanned any items yet")
	}

	var wg sync.WaitGroup
	totalChan := make(chan int, len(c.Items))

	for SKU, qty := range c.Items {
		wg.Add(1)
		go func(SKU string, qty int) {
			defer wg.Done()
			cost := c.Catalogue.Prices[SKU]
			tPrice := 0
			// each item check if the user has ordered enough items to meet the special price
			if cost.SpecialQty > 0 && qty >= cost.SpecialQty {
				// apply special price to all the items it can
				tPrice += (qty / cost.SpecialQty) * cost.SpecialPrice
				// apply normal price to remainder of the items
				tPrice += (qty % cost.SpecialQty) * cost.NormalPrice
			} else {
				tPrice += qty * cost.NormalPrice
			}
			totalChan <- tPrice
		}(SKU, qty)
	}

	wg.Wait()
	close(totalChan)

	totalPrice := 0
	for price := range totalChan {
		totalPrice += price
	}

	return totalPrice, nil
}
