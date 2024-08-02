package checkout

// implement interface
type ICheckout interface {
	Scan(SKU string) (err error)
	GetTotalPrice() (totalPrice int, err error)
}

// struct to satisfy interface

// new struct so we're able to mock in testing?

// Scan method

// GetTotal method