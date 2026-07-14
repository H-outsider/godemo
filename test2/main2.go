package main

import "fmt"

func main() {
	var funcs []func()

	for i := 1; i <= 3; i++ {
		funcs = append(funcs, func() {
			fmt.Println(i)
		})
	}

	for _, f := range funcs {
		f() // 这里会打印什么？
	}
}
