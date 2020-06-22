package bookstore

//BookGroupCalculator represent type of calculation logic of book collection
type BookGroupCalculator struct {
}

var discountPercent = map[int]int{
	0: 0,
	1: 0,
	2: 5,
	3: 10,
	4: 20,
	5: 25,
}

//GetAmount implements the book store exercise.
func (bc BookGroupCalculator) GetAmount(bookGroup []int, price int) ( interface{}) {
	sum := 0
	for _, value := range bookGroup {
		sum += value * (price - ((price * discountPercent[value]) / 100))
	}
	return sum
}
