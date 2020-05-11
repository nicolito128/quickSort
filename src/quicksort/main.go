package quicksort

type callback func(arr []int) []int

func parse(pivot int, left []int, right []int, handler callback) []int {
	left = handler(left)
	right = handler(right)

	left = append(left, pivot)
	left = append(left, right...)

	return left
}

// Standard ... ordenar un array
func Standard(arr []int) (result []int) {
	len := len(arr)
	if len <= 1 {
		return arr
	}

	var left, right []int
	pivot := arr[0]

	for i := 1; i < len; i++ {
		if arr[i] < pivot {
			left = append(left, arr[i])
		} else {
			right = append(right, arr[i])
		}
	}

	result = parse(pivot, left, right, Standard)
	return
}
