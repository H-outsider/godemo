package main

import "fmt"

func main() {
	s := []int{1, 2, 3}
	func(p []int) {
		p[0] = 10
		p = append(p, 4)
		p[0] = 100
	}(s)
	fmt.Println(s)
}
