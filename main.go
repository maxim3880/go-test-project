package main

import (
	"fmt"

	"./acronym"
	"./bookstore"
)

func main() {
	fmt.Println("Task 1 - Acronym - get acronym from 'portal network graphic' string")
	acronym.Execute()
	fmt.Println()
	fmt.Println("Task 2 - Bookshop - calculate total amount for different book collection")
	array := bookstore.Array{
		15, 15, 10, 5, 5,
	}
	bookstore.Cost(array)
	fmt.Println()
}
