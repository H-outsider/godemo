package main

import "fmt"

func main() {
	mp := map[string]int{
		"a": 0,
		"b": 1,
		"c": 2,
		"d": 3,
	}
	fmt.Println(mp["a"])
	fmt.Println(mp["b"])
	fmt.Println(mp["d"])
	fmt.Println(mp["f"])
}
