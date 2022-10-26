package invoice

import "fmt"

type Invoice struct {
	id            string
	date          string
	company       *Contact
	customer      *Contact
	items         []*Item
	totalDiscount Discount
}

func NewInvoice(id string, date string, company *Contact, customer *Contact, items []*Item, discount Discount) *Invoice {
	return &Invoice{id: id, date: date, company: company, customer: customer, items: items, totalDiscount: discount}
}
func (i *Invoice) TotalWithoutTaxAndWithoutFinalDiscount() float64 {
	var total float64
	for _, item := range i.items {
		total += item.TotalWithoutTaxAndWithDiscount()
	}
	return total
}

func (i *Invoice) TotalWithoutTaxAndWithFinalDiscount() float64 {
	total := i.TotalWithoutTaxAndWithoutFinalDiscount()

	discountType, discountValue := i.totalDiscount.getDiscount()
	if discountType == DISCOUNT_TYPE_PERCENT {
		total -= (discountValue * total / 100)
	} else if discountType == DISCOUNT_TYPE_AMOUNT {
		total -= discountValue
	}

	return total
}

func (i *Invoice) TotalWithTaxAndDiscount() float64 {
	return i.TotalWithoutTaxAndWithFinalDiscount() + i.TotalTax()
}

func (i *Invoice) TotalTax() float64 {
	totalWithoutTax := i.TotalWithoutTaxAndWithoutFinalDiscount()
	discountType, discountValue := i.totalDiscount.getDiscount()

	var totalTax float64

	if discountType == DISCOUNT_TYPE_AMOUNT && discountValue == 0 {
		for _, item := range i.items {
			totalTax += item.TaxOnDiscountedTotal()
		}
	} else {
		if discountType == DISCOUNT_TYPE_AMOUNT {
			if discountValue <= totalWithoutTax {
				discountValue = (discountValue * 100 / totalWithoutTax)
			} else {
				fmt.Println("DiscountValue cant be greater than totalWithoutTax")
			}
		}
		for _, item := range i.items {
			taxType, taxValue := item.GetTax().getTax()
			if taxType == TAX_TYPE_AMOUNT {
				totalTax += taxValue
			} else {
				itemTotal := item.TotalWithoutTaxAndWithDiscount()
				discount := (discountValue * itemTotal / 100)
				itemTotalDiscounted := itemTotal - discount

				itemTaxDiscounted := (taxValue * itemTotalDiscounted / 100)
				fmt.Println(totalWithoutTax, discountValue, discount, itemTotal, itemTaxDiscounted, "*********************")
				totalTax += itemTaxDiscounted
			}
		}
	}
	return totalTax
}
