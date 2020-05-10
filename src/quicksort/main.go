package quicksort

// Parse ... ordenar un array
func Parse(arr []int) []int {
	if len(arr) == 0 {
		return make([]int, 0)
	}

	var left, right []int
	pivot := arr[0]

	for i := 1; i < len(arr); i++ {
		if arr[i] < pivot {
			left = append(left, arr[i])
		} else {
			right = append(right, arr[i])
		}
	}

	left = Parse(left)
	right = Parse(right)

	left = append(left, pivot)
	left = append(left, right...)

	return left
}
