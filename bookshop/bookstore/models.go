package bookstore

//Order represent state of purchases
type Order struct {
	Books []Books
	Name string
}

//Books represent group of books
type Books struct {
	First, Second, Third, Fourth, Fifth int
}
