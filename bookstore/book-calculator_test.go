package bookstore

import (
	"testing"
)

//TestBookGroupCalculator run test of correct amount
func TestBookGroupCalculator(t *testing.T) {
	t.Run("check tootal amount for all discount", func(t *testing.T) {
		bookCalculator := BookGroupCalculator{}
		bookGroup := []int{5, 4, 3, 2, 1}
		price := 800
		got := bookCalculator.GetAmount(bookGroup, price).(int)

		want := 10040

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("check discount if group of book repeat", func(t *testing.T) {
		bookCalculator := BookGroupCalculator{}
		bookGroup := []int{4, 4, 4, 2, 2, 2}
		price := 800
		got := bookCalculator.GetAmount(bookGroup, price).(int)

		want := 12240

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

}
