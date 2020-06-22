package bookstore

//Order represent state of purchases
type Order struct {
	BookGroup         Array
	TotalAmount       int
	PriceWithDiscount float32
}

//Array represent []int
type Array []int

func (arr *Array) sum() (sum int) {
	for i := 0; i < len(*arr); i++ {
		sum += (*arr)[i]
	}
	return sum
}

func (arr *Array) subtract(model Array) {
	if len(*arr) == len(model) {
		for key, value := range *arr {
			(*arr)[key] = value - model[key]
		}
	}
}

func (arr Array) isArrayAcceptByMask(mask Array) (result bool) {
	if len(arr) == len(mask) {
		result = true
		for i := 0; i < len(arr); i++ {
			result = result && arr[i] >= mask[i]
		}
	}
	return result
}

func (arr *Array) fillBySource(source Array) {
	for _, x := range source {
		(*arr) = append((*arr), x)
	}
}

