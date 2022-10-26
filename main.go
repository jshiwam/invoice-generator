package main

import (
	"fmt"
	"time"

	inv "github.com/jshiwam/invoice-generator/invoice"
)

func main() {
	fmt.Println("Hello World")
	item1 := inv.NewItem("Coke", 2, 10, inv.NewTaxPercent(20), inv.NewDiscountPercent(10))

	item2 := inv.NewItem("Fanta", 2, 10, inv.NewTaxPercent(20), inv.NewDiscountPercent(10))
	company := inv.NewContact("Google", inv.NewAddress("Mountain View", "123456", "California", "USA"), inv.NewPhoneNumber("+1", "994395071"))
	customer := inv.NewContact("Shiwam", inv.NewAddress("Santa Barbara", "678921", "Texas", "USA"), inv.NewPhoneNumber("+1", "9943995072"))
	invoice := inv.NewInvoice("1", time.Now().Format("2006-02-01"), company, customer, []*inv.Item{item1, item2}, inv.NewDiscountAmount(10))

	fmt.Printf("%#v\n", invoice)
	fmt.Printf("Total Without tax and Without Final discount %v \n", invoice.TotalWithoutTaxAndWithoutFinalDiscount())
	fmt.Printf("Total Without tax and With Final discount %v \n", invoice.TotalWithoutTaxAndWithFinalDiscount())
	fmt.Printf("Total Tax %v \n", invoice.TotalTax())
	fmt.Printf("Total With tax and With Final discount %v \n", invoice.TotalWithTaxAndDiscount())
}
