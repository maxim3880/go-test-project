package bookstore

//Calculator represents  of method for calculation order
type Calculator interface {
	GetAmount(Array, float32) interface{}
}

//Generator represents  of method for calculation book collections with discount
type Generator interface {
	Generate(combinationModel) []Array
}
