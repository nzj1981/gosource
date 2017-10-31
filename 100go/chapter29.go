// chapter29.go
/*
计算数字

题目：给一个不多于5位的正整数，要求：一、求它是几位数，二、逆序打印出各位数字。
*/

package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var x int
	fmt.Printf("请输入一个数字")
	fmt.Scanf("%d\n", &x)
	l := len(strconv.Itoa(x))
	// for i := 5; i > 0; i-- {
	// 	r := x / int(math.Pow10(i-1))
	// 	if r > 0 {
	// 		fmt.Printf("%d是一个%d位数字.\n", x, i)
	// 		out(i, x)
	// 		fmt.Printf("\n")
	// 	}
	// }
	r := x / int(math.Pow10(l-1))
	if r > 0 {
		fmt.Printf("%d是一个%d位数字.\n", x, l)
		out(l, x)
		fmt.Printf("\n")
	}
}

func out(n, x int) {
	if n > 1 {
		out(n-1, x)
	}
	r := x % int(math.Pow10(n)) / int(math.Pow10(n-1))
	fmt.Printf("%d", r)
}
