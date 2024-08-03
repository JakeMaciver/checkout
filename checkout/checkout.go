package checkout

import (
	"errors"
	"fmt"

	"github.com/JakeMaciver/checkout/pricing"
)

// ICheckout represents an interface for the checkout
type ICheckout interface {
	// Scan is mehtod that will scan an item to be paid for
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

// Scan will add the SKU into an items list or increment if the item is already in the list, errors occur of rinvlid input and not found
func (c *Checkout) Scan(SKU string) error {
	// invalid input
	if len(SKU) == 0 {
		return errors.New("no item to scan")
	}

	// not found in catalogue
	if _, exists := c.Catalogue.Prices[SKU]; !exists {
		err := fmt.Sprintf("item not found in the catalogue: %s", SKU)
		return errors.New(err)
	}

	c.Items[SKU]++
	return nil
}

// GetTotal method
func (c *Checkout) GetTotalPrice() (int, error) {

	totalPrice := 0
	for SKU, qty := range c.Items {
		if c.Catalogue.Prices[SKU].SpecialQty > 0 && qty >= c.Catalogue.Prices[SKU].SpecialQty {
			totalPrice += (qty/c.Catalogue.Prices[SKU].SpecialQty) * c.Catalogue.Prices[SKU].SpecialPrice
			totalPrice += (qty % c.Catalogue.Prices[SKU].SpecialQty) * c.Catalogue.Prices[SKU].NormalPrice
		} else {
			totalPrice += qty * c.Catalogue.Prices[SKU].NormalPrice
		}
	}

	return totalPrice, nil
}
