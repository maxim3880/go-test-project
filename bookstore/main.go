package bookstore

import (
	"fmt"
)

//Cost start of presentation task 2
func Cost(basket Array) {
	orderCalculator := getOrderCalculator()
	field, ok := orderCalculator.GetAmount(basket, 800).(Order)
	if ok {
		fmt.Printf("Min Array amount = $%d ", field.TotalAmount)
		fmt.Println("Best group combination", field.BookGroup)
	}

}

func getOrderCalculator() Calculator {
	return &OrderBasketCalculator{
		BookGroupCalculator: BookGroupCalculator{},
		Generator:           &BookGroupGenerator{},
		MaxCount:            5,
	}
}
