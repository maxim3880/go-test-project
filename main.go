package main

import (
	"fmt"

	"./acronym"
	"./bookshop"
)

func main() {
	fmt.Println("Task 1 - Acronym - get acronym from 'portal network graphic' string")
	acronym.Execute()
	fmt.Println()
	fmt.Println("Task 2 - Bookshop - calculate total amount for different book collection")
	bookshop.Execute()
	fmt.Println()
}
