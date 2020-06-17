package bookstore

//OrderBasketCalculator represent type of calculation logic of order
type OrderBasketCalculator struct {
	Calculator
	Generator
	MaxCount int
}

//GetAmount calculate sum and returm
func (oc OrderBasketCalculator) GetAmount(basket Array, price float32) (result interface{}) {
	minPrice := float32(basket.sum()) * price
	combinations := oc.Generator.Generate(oc.generateModel(basket))
	orders := oc.generateOrdersByCombination(basket, combinations)
	for i := 0; i < len(orders); i++ {
		orders[i].TotalAmount = oc.calculateOrderAmount(orders[i].BookGroup, price)
		if orders[i].TotalAmount < minPrice {
			minPrice = orders[i].TotalAmount
			result = orders[i]
		}
	}
	return result
}

func (oc OrderBasketCalculator) generateOrdersByCombination(basket Array, combinations []Array) (result []Order) {
	result = make([]Order, 0)
	for i := 0; i < len(combinations); i++ {
		basketMask := Array{}
		basketMask.fillBySource(basket)
		basketMask.subtract(combinations[i])
		group := make([]Array, 0)
		oc.findAllCombinationByMask(combinations, combinations[i].sum(), basketMask, basketMask.sum(), &group)
		for _, value := range group {
			order := Order{BookGroup: value}
			result = append(result, order)
		}
	}
	return result
}

func (oc OrderBasketCalculator) calculateOrderAmount(bookGroup Array, price float32) (totalAmount float32) {
	totalAmount, _ = oc.Calculator.GetAmount(bookGroup, price).(float32)
	return totalAmount
}

func (oc OrderBasketCalculator) generateModel(combinations Array) combinationModel {
	return combinationModel{
		inputArray: combinations,
		start:      0,
		end:        0,
		size:       oc.MaxCount,
	}
}

var currentCombinations = Array{}

func (oc OrderBasketCalculator) findAllCombinationByMask(combinations []Array, useBooksCount int, basketMask Array, elementCount int, group *[]Array) {
	currentCombinations = append(currentCombinations, useBooksCount)
	if elementCount == 0 {
		(*group) = append(*group, currentCombinations)
		currentCombinations = Array{}
		return
	}
	for i := 0; i < len(combinations); i++ {
		if basketMask.isArrayAcceptByMask(combinations[i]) {
			basketMask.subtract(combinations[i])
			oc.findAllCombinationByMask(combinations, combinations[i].sum(), basketMask, basketMask.sum(), group)
		}
	}
}
