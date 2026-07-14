package main

import "fmt"

func main() {
	dest := make([]int, 0)
	src := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(src, dest)
	fmt.Println(copy(dest, src))
	fmt.Println(src, dest)

	var nums [5][5]int
	for _, num := range nums {
		fmt.Println(num)
	}
	fmt.Println()
	slices := make([][]int, 5)
	for _, slice := range slices {
		fmt.Println(slice)
	}

}
