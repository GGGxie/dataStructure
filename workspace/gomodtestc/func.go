package gomodtestc

import "fmt"

func PrintStr(str string, num int) string {
	return fmt.Sprintf("project 1C %s_%d", str, num)
}

// TestWork 新增功能。。。
func TestWork() string {
	return "hello workespace"
}
