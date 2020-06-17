package bookstore

import (
	"fmt"
)

//Cost start of presentation task 2
func Cost(basket Array) {
	orderCalculator := getOrderCalculator()
	field, ok := orderCalculator.GetAmount(basket, 8).(Order)
	if ok {
		fmt.Printf("Min Array amount = $%0.2f ", field.TotalAmount)
		fmt.Println("Best group combination", field.BookGroup)
	}

}

func getOrderCalculator() Calculator {
	return &OrderBasketCalculator{
		Calculator: &BookGroupCalculator{},
		Generator:  &BookGroupGenerator{},
		MaxCount:   5,
	}
}
