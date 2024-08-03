package pricing

import (
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

// AddItem will add an item to the catalogue if it already doesnt exist
func (c *Catalogue) AddItem(SKU string, normalPrice int, specialQty int, specialPrice int) error {
	modifiedSpecialPrice, err := validateInputs(SKU, normalPrice, specialQty, specialPrice); 
	if err != nil {
		return err
	}

	if _, ok := c.Prices[SKU]; ok {	
		return fmt.Errorf("item already exists: %s", SKU)	
	}

	c.Prices[SKU] = ItemPricing{
		NormalPrice: normalPrice,
		SpecialQty: specialQty,
		SpecialPrice: modifiedSpecialPrice,
	}

	return nil
}

// UpdateItem will update an existing item in the prices map
func (c *Catalogue) UpdateItem(SKU string, normalPrice int, specialQty int, specialPrice int) error {
	modifiedSpecialPrice, err := validateInputs(SKU, normalPrice, specialQty, specialPrice); 
	if err != nil {
		return err
	}

	if _, ok := c.Prices[SKU]; !ok {	
		return fmt.Errorf("item doesnt exist: %s", SKU)	
	}

	c.Prices[SKU] = ItemPricing{
		NormalPrice: normalPrice,
		SpecialQty: specialQty,
		SpecialPrice: modifiedSpecialPrice,
	}

	return nil
}

// DeleteItem will delete an item from the prices map if it exists
func (c *Catalogue) DeleteItem(SKU string) error {
	if len(SKU) != 1 {
		return fmt.Errorf("invalid SKU: %s", SKU)
	}
	charSKU := SKU[0]
	rSKU := rune(charSKU)
	if !unicode.IsUpper(rSKU) || !unicode.IsLetter(rSKU) {
		return fmt.Errorf("invalid SKU: %s", SKU)
	}

	if _, ok := c.Prices[SKU]; !ok {	
		return fmt.Errorf("item doesnt exist: %s", SKU)	
	}

	delete(c.Prices, SKU)
	return nil
}

func validateInputs(SKU string, normalPrice int, specialQty int, specialPrice int) (int, error) {
	if len(SKU) != 1 {
		return specialPrice, fmt.Errorf("invalid SKU: %s", SKU)
	}
	charSKU := SKU[0]
	rSKU := rune(charSKU)
	if !unicode.IsUpper(rSKU) || !unicode.IsLetter(rSKU) {
		return specialPrice, fmt.Errorf("invalid SKU: %s", SKU)
	}

	if normalPrice <= 0 {
		return specialPrice, fmt.Errorf("invalid normal price: %d", normalPrice)
	}

	if specialQty == 0 {
		specialPrice = 0
	}
	
	return specialPrice, nil
}