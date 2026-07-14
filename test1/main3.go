package main

import "fmt"

func main() {
	s := make([]int, 2, 4)
	s = append(s, 10)
	fmt.Println(len(s), cap(s))
}
