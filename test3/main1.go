package main

import (
	"fmt"
	"strconv"
)

func parseAndAdd(a, b string) int {
	num1, _ := strconv.Atoi(a)
	num2, err := strconv.Atoi(b)

	// 糟糕！这里忘记写 if err != nil 检查了，直接使用了 num2
	fmt.Println(err)
	return num1 + num2
}

func main() {
	result := parseAndAdd("10", "abc")
	fmt.Println(result)
}
