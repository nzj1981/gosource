// chapter07.go
/*
题目：输出特殊图案，请在 c 环境中运行，看一看，Very Beautiful!
1.程序分析：字符共有 256 个。不同字符，图形不一样。
*/
package main

import (
	"fmt"
)

func main() {
	var a, b = 176, 219
	fmt.Printf("%c%c%c%c%c \n", b, a, a, a, b)
	fmt.Printf("%c%c%c%c%c \n", a, b, a, b, a)
	fmt.Printf("%c%c%c%c%c \n", a, a, b, a, a)
	fmt.Printf("%c%c%c%c%c \n", a, b, a, b, a)
	fmt.Printf("%c%c%c%c%c \n", b, a, a, a, b)
}
