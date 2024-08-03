package main

import (
	"github.com/JakeMaciver/checkout/checkout"
	"github.com/JakeMaciver/checkout/pricing"
)

func main() {
	catalogue := pricing.NewCatalogue(make(map[string]pricing.ItemPricing))
	checkout := checkout.NewCheckout(*catalogue)
	
}