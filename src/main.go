package main

import (
	"fmt"
	"quickSort/src/quicksort"
)

func main() {
	nums := []int{9, 33, 8, 9, 7, 6, 5, 56, 4, 17, 5, 2, 2}
	fmt.Println(quicksort.Standard(nums))
}
