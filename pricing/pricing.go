package pricing

import (
	"errors"
	"fmt"
	"unicode"
)

// Catalogue defines the collection of items and mapping them to their pricing schema
type Catalogue struct {
	Prices map[string]ItemPricing
}

// ItemPricing defines the pricing schema per item in a Catalogue
type ItemPricing struct {
	// NormalPrice represent the usual price of an item
	NormalPrice int
	// SpecialPrice represents the special price of an item when a quantity has been reached
	SpecialPrice int
	// SpecialQty represents the number needed to be reached in order for the SpecialPrice to come in effect
	SpecialQty int
}

// NewCatalogue initialises and returns a Catalogue instance
func NewCatalogue(itemPrices map[string]ItemPricing) *Catalogue {
	return &Catalogue{
		Prices: itemPrices,
	}
}

func (c *Catalogue) AddItem(SKU string, normalPrice int, specialQty int, specialPrice int) error {
	if len(SKU) != 1 {
		err := fmt.Sprintf("invalid SKU: %s", SKU)
		return errors.New(err)
	}
	charSKU := SKU[0]
	rSKU := rune(charSKU)
	if !unicode.IsUpper(rSKU) || !unicode.IsLetter(rSKU) {
		err := fmt.Sprintf("invalid SKU: %s", SKU)
		return errors.New(err)		
	}

	c.Prices[SKU] = ItemPricing{
		NormalPrice: normalPrice,
		SpecialQty: specialQty,
		SpecialPrice: specialPrice,
	}

	return nil
}