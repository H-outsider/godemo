package main

import "fmt"

func main() {
	fmt.Println("--- 场景1：基础多分支匹配 ---")
	basicSwitch()

	fmt.Println("\n--- 场景2：带初始化语句的 switch ---")
	initStatementSwitch()

	fmt.Println("\n--- 场景3：无条件表达式的 switch ---")
	conditionSwitch()

	fmt.Println("\n--- 场景4：使用 fallthrough 强制穿透 ---")
	fallthroughSwitch()
}

// basicSwitch 场景1：基础的字符串匹配
func basicSwitch() {
	// 场景：在微服务开发中，常用于根据 HTTP Method 或不同策略模式路由业务逻辑
	str := "a"

	// Go 的 switch 默认自带 break，匹配到一个 case 执行完毕后自动跳出，不会发生意外穿透
	switch str {
	case "a":
		str += "a"
		str += "c"
	case "b":
		str += "bb"
		str += "aaaa"
	default:
		// 当所有 case 都不匹配后，就会执行 default 分支
		str += "CCCC"
	}
	fmt.Println("场景1执行结果:", str)
}

// initStatementSwitch 场景2：带初始化语句，实现作用域隔离
func initStatementSwitch() {
	// 场景：常用于先调用接口/数据库获取数据，然后立即进行状态判断。
	// num := f() 的作用域仅限于该 switch 代码块内，执行完毕后立即释放内存，避免变量逃逸或污染外层作用域。
	switch num := f(); { // 等价于 switch num := f(); true {
	case num >= 0 && num <= 1:
		num++
	case num > 1:
		num--
		fallthrough // 注意：这里使用了 fallthrough，会无条件执行下一个 case 的代码块
	case num < 0:
		num += num
	}
}

// f 模拟一个获取状态码的底层函数（如从 Redis 读取配置）
func f() int {
	return 1
}

// conditionSwitch 场景3：无目标表达式，直接判断条件（Go 原生思维极力推荐的写法）
func conditionSwitch() {
	num := 2

	// 省略表达式，等价于 switch true。
	// 这是 Go 语言重构深层级/复杂的 if-else if-else 链的标准做法，代码更加扁平化。
	switch {
	case num >= 0 && num <= 1:
		num++
	case num > 1:
		num--
	case num < 0:
		num *= num
	}
	fmt.Println("场景3执行结果:", num)
}

// fallthroughSwitch 场景4：演示 fallthrough 的强行穿透效应
func fallthroughSwitch() {
	num := 2

	switch {
	case num >= 0 && num <= 1:
		num++
	case num > 1:
		num--
		// fallthrough 会打破 Go 的默认防穿透机制。
		// 执行完该分支后，会强行继续执行相连的下一个 case 分支，且【不检查】下一个 case 的条件是否成立。
		fallthrough
	case num < 0:
		num += num
	}
	fmt.Println("场景4执行结果:", num)
}
