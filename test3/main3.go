package main

import (
	"fmt"
)

// 注意这里的返回值写法：(res int, err error)
// 我们给返回值起了名字，这在拦截 panic 时非常关键
func SafeDivide(a, b int) (res int, err error) {
	// 1. 必须在可能发生 panic 的代码之前，提前挂载 defer
	defer func() {
		// 2. 调用 recover() 尝试捕获 panic
		// 如果程序正常运行，r 会是 nil；如果发生了 panic，r 就是 panic 的具体内容
		if r := recover(); r != nil {
			// 3. 将捕获到的 panic 转换为普通的 error，赋值给命名返回值 err
			err = fmt.Errorf("division by zero intercepted: %v", r)
		}
	}()

	// 4. 执行真正的业务逻辑
	// 如果 b 为 0，这一行会立刻触发 panic，程序执行流会瞬间跳转到上面的 defer 中
	res = a / b

	return res, err
}

func main() {
	// 正常测试
	res1, err1 := SafeDivide(10, 2)
	fmt.Printf("正常情况 -> 结果: %d, 错误: %v\n", res1, err1)

	// 触发 panic 的测试
	res2, err2 := SafeDivide(10, 0)
	fmt.Printf("异常情况 -> 结果: %d, 错误: %v\n", res2, err2)
}
