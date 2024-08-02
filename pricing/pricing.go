package pricing

// Catalogue defines the collection of items and mapping them to their pricing schema
type Catalogue struct {
	Prices map[string]ItemPricing
}

// ItemPricing defines the pricing schema per item in a Catalogue
type ItemPricing struct {
	NormalPrice int
	SpecialPrice int
	SpecialQty int
}

// NewCatalogue initialises and returns a Catalogue instance
func NewCatalogue(itemPrices map[string]ItemPricing) *Catalogue {
	return &Catalogue{
		Prices: itemPrices,
	}
}