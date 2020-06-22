package bookstore

//OrderBasketCalculator represent type of calculation logic of order
type OrderBasketCalculator struct {
	BookGroupCalculator
	Generator
	MaxCount int
}

var currentPrice int = 0

//GetAmount calculate sum and returm
func (oc OrderBasketCalculator) GetAmount(basket Array, price int) interface{} {
	currentPrice = price
	combinations := oc.Generator.Generate(oc.MaxCount)
	minCostCombination, amount := oc.generateOrdersByCombination(basket, combinations)
	return Order{
		BookGroup:   minCostCombination,
		TotalAmount: amount,
	}
}

func (oc OrderBasketCalculator) generateOrdersByCombination(basket Array, combinations []Array) (result Array, minAmount int) {
	minAmount = basket.sum() * currentPrice
	oc.findAllCombinationByMask(combinations, Array{}, basket, &minAmount, &result)
	return result, minAmount
}

func (oc OrderBasketCalculator) calculateOrderAmount(bookGroup []int, price int) int {
	return oc.BookGroupCalculator.GetAmount(bookGroup, price).(int)
}

func (oc OrderBasketCalculator) findAllCombinationByMask(combinations []Array, currentCombinations Array, basket Array, minAmount *int, group *Array) {
	if basket.sum() == 0 {
		(*minAmount) = oc.calculateOrderAmount(currentCombinations, currentPrice)
		(*group) = currentCombinations
		return
	}
	for i := 0; i < len(combinations); i++ {
		if basket.isArrayAcceptByMask(combinations[i]) {
			bookGroups := Array{}
			bookGroups = append(bookGroups, currentCombinations...)
			bookGroups = append(bookGroups, combinations[i].sum())
			if oc.calculateOrderAmount(bookGroups, currentPrice) <= (*minAmount) {
				basketMask := Array{}
				basketMask = append(basketMask, basket...)
				basketMask.subtract(combinations[i])
				oc.findAllCombinationByMask(combinations, bookGroups, basketMask, minAmount, group)
			} else {
				return
			}
		}
	}
}
