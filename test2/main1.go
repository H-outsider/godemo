package main

import "fmt"

func main() {
	x := 10

	// Defer 1：直接调用常规函数
	defer fmt.Println("A:", x)

	// Defer 2：调用匿名函数，按值传参
	defer func(val int) {
		fmt.Println("B:", val)
	}(x)

	// Defer 3：调用匿名函数，无参数，直接捕获外部变量
	defer func() {
		fmt.Println("C:", x)
	}()

	x = 20
}
