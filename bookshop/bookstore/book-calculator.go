package bookstore

import "reflect"

//BookTotalCalculator represent type of calculation logic of book collection
type BookTotalCalculator struct {
	Price float32
}

// Cost implements the book store exercise.
func (bc BookTotalCalculator) Cost(books Books) float32 {
	totalBookCount, differentPartCount := bc.getBookCount(books)
	discount := bc.getCollectionDiscount(differentPartCount)
	return float32(totalBookCount) * (bc.Price - (bc.Price * float32(discount)))
}

func (bc *BookTotalCalculator) getCollectionDiscount(differentBookCount int) (discount float32) {

	switch differentBookCount {
	case 2:
		discount = 0.05
	case 3:
		discount = 0.1
	case 4:
		discount = 0.2
	case 5:
		discount = 0.25
	}
	return discount
}

func (bc *BookTotalCalculator) getBookCount(books Books) (totalBookCount int, differentPartCount int) {
	bookValues := reflect.ValueOf(books)
	for i := 0; i < bookValues.NumField(); i++ {
		value := int(bookValues.Field(i).Int())
		totalBookCount += value
		if value > 0 {
			differentPartCount++
		}
	}
	return totalBookCount, differentPartCount
}
