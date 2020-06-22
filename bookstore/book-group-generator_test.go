package bookstore

import (
	"testing"
)

func TestBookGroupGenerator(t *testing.T) {
	t.Run("check is generate correct combination", func(t *testing.T) {
		bookGenerator := BookGroupGenerator{}
		size := 3
		got := bookGenerator.Generate(size)

		want := []Array{
			Array{1, 1, 1},
			Array{1, 1, 0},
			Array{1, 0, 1},
			Array{0, 1, 1},
			Array{1, 0, 0},
			Array{0, 1, 0},
			Array{0, 0, 1},
		}

		if isArrayEqual(want, got) {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("check is generate correct number of combination", func(t *testing.T) {
		bookGenerator := BookGroupGenerator{}
		size := 5
		got := len(bookGenerator.Generate(size))

		want := 31

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

}

func isArrayEqual(a, b []Array) (result bool) {
	result = len(a) == len(b)
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			result = result && a[i][j] != b[i][j]
		}
	}
	return result
}
