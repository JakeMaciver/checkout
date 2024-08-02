package checkout

import "github.com/JakeMaciver/checkout/pricing"

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

// Scan method

// GetTotal method