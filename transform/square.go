package transform

// SquareSlices return suqare of a slice
func SquareSlices(s []int) []int {
	squares := make([]int, len(s))

	for index, value := range s {
		squares[index] = value * value
	}

	return squares
}
