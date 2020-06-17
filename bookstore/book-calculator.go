package bookstore

//BookGroupCalculator represent type of calculation logic of book collection
type BookGroupCalculator struct {
}

var discountPercent = map[int]float32{
	0: 0,
	1: 0,
	2: 0.05,
	3: 0.10,
	4: 0.20,
	5: 0.25,
}

//GetAmount implements the book store exercise.
func (bc BookGroupCalculator) GetAmount(bookGroup Array, price float32) interface{} {
	var sum float32 = 0.0
	for _, value := range bookGroup {
		sum += float32(value) * (price - (price * discountPercent[value]))
	}
	return sum
}
