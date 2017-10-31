// chapter08.go
/*
题目：输出 9*9 口诀。
1.程序分析：分行与列考虑，共 9 行 9 列，i 控制行，j 控制列。
*/
package main

import (
	"fmt"
)

func main() {
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%-3d", j, i, i*j) //-3d表示左对齐,占3位
		}
		fmt.Println("")
	}
}
