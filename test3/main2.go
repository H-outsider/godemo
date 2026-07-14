package main

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("record not found")

func queryDB() error {
	// 模拟底层数据库查询失败，并用 %w 包装了原始错误
	return fmt.Errorf("query failed: %w", ErrNotFound)
}

func main() {
	err := queryDB()
	if err != nil {
		// 【填空处】：这里应该用什么代码，才能准确判断出底层错误是 ErrNotFound？
		if errors.Is(err, ErrNotFound) {
		}
	}
}
