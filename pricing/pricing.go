package pricing

// pricing struct?
type Catalogue struct {
	Prices map[string]ItemPricing
}

// individual item pricing breakdown, in struct?
type ItemPricing struct {
	NormalPrice int
	SpecialPrice int
	SpecialQty int
}

// constructor function for struct
