package bookstore



//OrderAmountCalculator represent type of calculation logic of order
type OrderAmountCalculator struct {
	TotalAmount         float32
	TotalDiscountAmount float32
	BookCalculatorLogic BookCalculator
}

//CalcOrder calculate sum and returm
func (oc *OrderAmountCalculator) CalcOrder(order Order) float32 {
	var sum float32 = 0.0
	for _, value := range order.Books {
		sum += oc.BookCalculatorLogic.Cost(value)
	}
	return sum
}