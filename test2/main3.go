package main

import "fmt"

// NewCounter 是一个“工厂函数”，它负责组装并返回一个闭包
func NewCounter() func() int {
	// 1. 定义一个局部变量 count，作为初始状态
	count := 0

	// 2. 返回一个匿名函数（这就是闭包）
	return func() int {
		// 闭包内部引用了外层的 count 变量
		count++
		return count
	}
}

// main 函数才是程序真正的入口，用来测试我们的闭包
func main() {
	counter1 := NewCounter()
	fmt.Println(counter1()) // 输出 1
	fmt.Println(counter1()) // 输出 2

	counter2 := NewCounter()
	fmt.Println(counter2()) // 输出 1 (独立的计数器)
}
