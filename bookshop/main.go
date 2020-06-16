package bookshop

import (
	"fmt"

	"./bookstore"
)


//Execute start of presentation task 2
func Execute() {
 
	orders := getOrders()
	orderCalculator := getOrderCalculator()
	for _, value := range orders {
		fmt.Printf("Total amount of order = %0.2f ", orderCalculator.CalcOrder(value))
		fmt.Println("for example = ", value.Name)
	}
}

func getOrderCalculator() *bookstore.OrderAmountCalculator {
	return &bookstore.OrderAmountCalculator{
		BookCalculatorLogic: &bookstore.BookTotalCalculator{
			Price: 8,
		},
	}
}

func getOrders() []bookstore.Order {
	return []bookstore.Order{
		{
			Books: []bookstore.Books{
				bookstore.Books{1, 1, 1, 1, 1},
				bookstore.Books{1, 1, 1, 0, 0},
			},
			Name: "1 group of 5 --> 25% discount (1st,2nd,3rd,4th,5th) + 1 group of 3 --> 10% discount (1st,2nd,3rd)",
		},
		{
			Books: []bookstore.Books{
				bookstore.Books{1, 1, 1, 1, 0},
				bookstore.Books{1, 1, 1, 0, 1},
			},
			Name: "1 group of 4 books --> 20% discount (1st,2nd,3rd,4th) + 1 group of 4 books --> 20% discount (1st,2nd,3rd,5th)",
		},
	}
}
