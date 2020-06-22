package bookstore

//Calculator represents  of method for calculation order
type Calculator interface {
	GetAmount(Array, int) interface{}
}

//Generator represents  of method for calculation book collections with discount
type Generator interface {
	Generate(int) []Array
}
