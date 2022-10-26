package invoice

const (
	DISCOUNT_TYPE_PERCENT string = "percent"
	DISCOUNT_TYPE_AMOUNT  string = "amount"
)

type Discount struct {
	percent float64
	amount  float64
}

func NewDiscountPercent(percent float64) Discount {
	return Discount{percent: percent}
}
func NewDiscountAmount(amount float64) Discount {
	return Discount{amount: amount}
}

func (d Discount) getDiscount() (string, float64) {
	if d.percent > 0 {
		return DISCOUNT_TYPE_PERCENT, d.percent
	}
	return DISCOUNT_TYPE_AMOUNT, d.amount
}
