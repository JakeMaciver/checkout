package checkout

import (
	"errors"

	"github.com/JakeMaciver/checkout/pricing"
)

// implement interface
type ICheckout interface {
	Scan(SKU string) (err error)
	GetTotalPrice() (totalPrice int, err error)
}

// struct to satisfy interface
type Checkout struct {
	Items map[string]int
	Catalogue pricing.Catalogue
}

// new struct so we're able to mock in testing?
func NewCheckout(catalogue pricing.Catalogue) *Checkout {
	return &Checkout{
		Items: make(map[string]int),
		Catalogue: catalogue,
	}
}

// Scan method
func (c *Checkout) Scan(SKU string) error {
	if len(SKU) == 0 {
		return errors.New("no item to scan")	
	}

	c.Items[SKU]++
	return nil
}
// GetTotal method