// chapter10.go
/*
题目：打印楼梯，同时在楼梯上方打印笑脸。
1.程序分析：用 i 控制行，j 来控制列，j 根据 i 的变化来控制输出黑方格的个数。
*/
package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 8; i++ {
		fmt.Printf("%s\n", " ☺")
		for j := 0; j <= i; j++ {
			fmt.Printf("■■")
		}
		fmt.Println("")
	}
}
