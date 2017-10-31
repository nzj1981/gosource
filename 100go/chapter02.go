//chapter02
/*
题目：企业发放的奖金根据利润提成。利润(I)低于或等于10万元时，奖金可提成10%；利润高于10万元，低于20万元，低于10万元的部分按10%提成，高于10万元的部分，可提成7.5%；20万到40万之间时，高于20万元的部分，可提成5%；40万到60万之间时高于40万元的部分，可提成3%；60万到100万之间时，高于60万元的部分，可提成1.5%，高于100万元时，超过100万元的部分按1%提成，从键盘输入当月利润I，求应发放奖金总数？
*/

package main

import (
	"fmt"
)

func main() {
	var I float32 = 0.0
	var bonus float32 = 0.0 // 定义奖金变量保留一位小数
	fmt.Print("输入利润:")
	fmt.Scanf("%f\n", &I)
	//switch用法
	switch {
	case I > 1000000:
		bonus = (I - 1000000) * 0.01
		I = 1000000
		fallthrough
	case I > 600000:
		bonus += (I - 600000) * 0.015
		I = 600000
		fallthrough
	case I > 400000:
		bonus += (I - 400000) * 0.03
		I = 400000
		fallthrough
	case I > 200000:
		bonus += (I - 200000) * 0.05
		I = 200000
		fallthrough
	case I > 100000:
		bonus += (I - 100000) * 0.075
		I = 100000
		fallthrough
	default:
		bonus += I * 0.1
	}
	fmt.Printf("提成总计:%f\n", bonus)
}
