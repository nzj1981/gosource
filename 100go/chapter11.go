// chapter11.go
/*
题目：古典问题：有一对兔子，从出生后第3个月起每个月都生一对兔子，小兔子长到第三个月后每个月又生一对兔子，假如兔子都不死，问每个月的兔子总数为多少？
1.程序分析： 兔子的规律为数列 1,1,2,3,5,8,13,21....
*/

package main

import (
	"fmt"
)

func main() {
	var m1, m2 int = 1, 1
	k := make([]int, 0)
	for i := 1; i < 12; i++ {
		if len(k) >= 12 {
			break
		}
		k = append(k, m1)
		k = append(k, m2)

		m1 += m2
		m2 += m1
	}
	fmt.Printf("每个月是:%v\n", k)

}
