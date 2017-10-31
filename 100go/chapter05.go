// chapter05.go
/*
题目：输入三个整数 x,y,z，请把这三个数由小到大输出。
1.程序分析：我们想办法把最小的数放到 x 上，先将 x 与 y 进行比较，如果 x>y 则将 x 与 y
的值进行交换，然后再用 x 与 z 进行比较，如果 x>z 则将 x 与 z 的值进行交换，这样能使 x 最小。
*/
package main

import (
	"fmt"
)

func main() {
	var x, y, z int = 0, 0, 0
	fmt.Println("(请输入三个数字，以空格隔开:例如 3 6 2):")
	fmt.Scanf("%d%d%d", &x, &y, &z)
	if x > y {
		x, y = y, x
	}
	if x > z {
		x, z = z, x
	}
	if y > z {
		y, z = z, y
	}
	fmt.Printf("%d<%d<%d\n", x, y, z)
}
