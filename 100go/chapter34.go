// chapter34.go
/*
函数调用

题目：练习函数调用
*/

package main

import (
	"fmt"
)

func main() {
	noresult()
	ret := oneResult()
	fmt.Println("存在1个返回值的函数", ret)
	_, b := twoResult()
	fmt.Println("存在2个返回值的函数", b)
	_, _, _ = threeResult()
}

func noresult() {
	fmt.Println("没有返回值的函数")
}
func oneResult() int {
	fmt.Println("存在1个返回值的函数")
	return 1
}

func twoResult() (int, bool) {
	fmt.Println("存在2个返回值的函数")
	return 2, true
}
func threeResult() (int, bool, string) {
	fmt.Println("存在3个返回值的函数")
	return 3, false, "错误"
}
