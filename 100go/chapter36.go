// chapter36.go
/*
求素数

题目：求 100 之内的素数
*/

package main

import (
	"fmt"
)

func main() {
	for i := 3; i < 100; i++ {
		var j = 2
		for j = 2; j < i; j++ {
			if i%j == 0 {
				break
			}
		}
		if i == j {
			fmt.Print(i, " ")
		}
	}
}
