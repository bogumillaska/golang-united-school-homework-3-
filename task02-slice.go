package homework

func reverse(input []int64) (result []int64) {
	reversed := make([]int64, len(input))
	for i, j := 0, len(input)-1; j >= 0; i, j = i+1, j-1 {
		reversed[i] = input[j]
	}
	return reversed
}
