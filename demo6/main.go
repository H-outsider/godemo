package main

import "fmt"

func main() {
	fib := Fib()
	for i := 0; i < 10; i++ {
		fmt.Println(fib())
	}
}

func Fib() func() int {
	a, b := 1, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}
