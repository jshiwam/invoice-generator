package invoice

const (
	TAX_TYPE_PERCENT string = "percent"
	TAX_TYPE_AMOUNT  string = "amount"
)

type Tax struct {
	percent float64
	amount  float64
}

func NewTaxPercent(percent float64) Tax {
	return Tax{percent: percent}
}
func NewTaxAmount(amount float64) Tax {
	return Tax{amount: amount}
}

func (t Tax) getTax() (string, float64) {
	if t.percent > 0 {
		return TAX_TYPE_PERCENT, t.percent
	}
	return TAX_TYPE_AMOUNT, t.amount
}
