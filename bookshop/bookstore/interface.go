package bookstore

//OrderCalculator represents  of method for calculation order
type OrderCalculator interface {
	CalcOrder(Order) float32
}

//BookCalculator represents  of method for calculation book collections with discount
type BookCalculator interface {
	Cost(Books) float32
	getBookCount(Books) (int, int)
	getCollectionDiscount(int) (float32) 
}
