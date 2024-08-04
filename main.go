package main

import (
	"fmt"
	"time"

	"github.com/JakeMaciver/checkout/checkout"
	"github.com/JakeMaciver/checkout/pricing"
)

func main() {
	catalogue := pricing.NewCatalogue(make(map[string]pricing.ItemPricing))
	checkout := checkout.NewCheckout(*catalogue)

	itemsToAdd := map[string]pricing.ItemPricing{
		"A": {
			NormalPrice: 20,
			SpecialQty: 3,
			SpecialPrice: 45,
		},
		"B": {
			NormalPrice: 15,
			SpecialQty: 2,
			SpecialPrice: 20,
		},
		"C": {
			NormalPrice: 30,
			SpecialQty: 0,
			SpecialPrice: 0,
		},
	}

	fmt.Println("adding 3 items to the catalogue: A, B and C...")
	for SKU, itemPrice := range itemsToAdd {
		if err := catalogue.AddItem(SKU, itemPrice.NormalPrice, itemPrice.SpecialQty, itemPrice.SpecialPrice); err != nil {
			fmt.Println("error adding item to cataloguge: ", err)
		} else {
			fmt.Printf("successfully added item to catalogue: %s, %+v\n", SKU, itemPrice)
		}
	}

	fmt.Println("\nupdating item A in the catalogue...")
	if err := catalogue.UpdateItem("A", 15, 3, 30); err != nil {
		fmt.Println("error updating item in the catalogue: ", err)
	} else {
		fmt.Println("successfully updated item in the catalogue: A, {NormalPrice:15 SpecialPrice:30 SpecialQty:3}", )
	}

	fmt.Println("\ndeleting item B in the catalogue...")
	if err := catalogue.DeleteItem("B"); err != nil {
		fmt.Println("error deleting item in the catalogue: ", err)
	} else {
		fmt.Println("successfully deleted item from the catalogue")
	}

	itemsToScan := []string{"A", "A", "C", "C", "A", "A"}

	fmt.Println("\nScanning items...")
	for _, scanItem := range itemsToScan {
		fmt.Println("scanning ", scanItem)
		if err := checkout.Scan(scanItem); err != nil {
			fmt.Println("error scanning item: ", err)
		}
	}

	fmt.Println("\nCalculating total price...")
	start := time.Now()
	if total, err := checkout.GetTotalPrice(); err != nil {
		fmt.Println("error getting total price: ", err)
	} else {
		fmt.Println("Total is: ", total)
	}
	elapsed := time.Since(start)
	fmt.Printf("Time taken without concurrency: %+v\n", elapsed.Nanoseconds())
}