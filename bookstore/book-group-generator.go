package bookstore

//BookGroupGenerator represent book grouping logic in order
type BookGroupGenerator struct {
}

var buffer = Array{}
var count = 0

//Generate return different book groups
func (bg BookGroupGenerator) Generate(model combinationModel) (result []Array) {
	result = make([]Array, 0)
	for size := model.size; size > 0; size-- {
		buffer = make(Array, size)
		bg.combinationUtil(model.end, model.start, size, model.size, &result)
	}
	return result
}

func (bg BookGroupGenerator) combinationUtil(end int, begin int, k int, n int, result *[]Array) {
	for i := begin; i < n; i++ {
		buffer[end] = i
		if end+1 < k {
			bg.combinationUtil(end+1, buffer[end]+1, k, n, result)
		} else {
			group := make([]int, n)
			for _, value := range buffer {
				group[value] = 1
			}
			(*result) = append((*result), group)
		}
	}
}
