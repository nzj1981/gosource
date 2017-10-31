// chapter37.go
/*
求最大数

题目：对 10 个数进行排序
1.程序分析：可以利用选择法，即从后9个比较过程中，选择一个最小的与第一个元素交换，下次类推，即用第二个元素与后 8 个进行比较，并进行交换。
*/

package main

import (
	"fmt"
)

func main() {
	var array [10]int
	for i := 0; i < 10; i++ {
		fmt.Printf("array[%d]=", i)
		fmt.Scanf("%d\n", &array[i])
	}
	fmt.Println()
	for i := 0; i < 9; i++ {
		for j := i + 1; j < 9; j++ {
			if array[i] > array[j] {
				array[i], array[j] = array[j], array[i]
			}
			if array[i] > array[j+1] {
				array[i], array[j+1] = array[j+1], array[i]
			}
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	fmt.Println(array)
}
