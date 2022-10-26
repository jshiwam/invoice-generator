package invoice

type Item struct {
	name      string
	quantity  int
	unitPrice float64
	tax       Tax
	discount  Discount
}

func NewItem(name string, quantity int, unitPrice float64, tax Tax, discount Discount) *Item {
	return &Item{
		name:      name,
		quantity:  quantity,
		unitPrice: unitPrice,
		tax:       tax,
		discount:  discount,
	}
}

func (i *Item) GetName() string {
	return i.name
}

func (i *Item) GetQuantity() int {
	return i.quantity
}

func (i *Item) GetUnitPrice() float64 {
	return i.unitPrice
}

func (i *Item) GetTax() Tax {
	return i.tax
}

func (i *Item) TotalWithoutTaxAndWithoutDiscount() float64 {
	return i.unitPrice * float64(i.quantity)
}

func (i *Item) TotalWithoutTaxAndWithDiscount() float64 {

	total := i.TotalWithoutTaxAndWithoutDiscount()

	discountType, discountValue := i.discount.getDiscount()
	if discountType == DISCOUNT_TYPE_PERCENT {
		total -= (discountValue * total / 100)
	} else if discountType == DISCOUNT_TYPE_AMOUNT {
		total -= discountValue
	}
	return total
}

func (i *Item) TotalWithTaxAndDiscount() float64 {
	return i.TotalWithoutTaxAndWithDiscount() + i.TaxOnDiscountedTotal()
}

func (i *Item) TaxOnDiscountedTotal() float64 {
	total := i.TotalWithoutTaxAndWithDiscount()
	var result float64

	taxType, taxValue := i.tax.getTax()
	if taxType == TAX_TYPE_PERCENT {
		result = (taxValue * total / 100)
	} else if taxType == TAX_TYPE_AMOUNT {
		result = taxValue
	}
	return result
}
