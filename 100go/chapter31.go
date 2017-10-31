// chapter31.go
/*
检索字符串

题目：请输入星期几的第一个字母来判断一下是星期几，如果第一个字母一样，则继续判断第二个字母。
1.程序分析：用情况语句比较好，如果第一个字母一样，则判断用情况语句或 if 语句判断第
二个字母。
*/
package main

import (
	"fmt"
	"strings"
)

func main() {
	// inweek := []string{"Monday", "Tuesday", "Wednesday", "Thurday", "Friday", "Saturday", "Sunday"}
	inweek := []string{"成人之美", "好吃赖做", "几起几落", "巾帼英雄", "几个风云"}
	instr := ""
	fmt.Printf("input a char:")
	fmt.Scanf("%s\n", &instr)
	outweek(instr, inweek)
}

func outweek(instr string, inweek []string) {
	nextstr := ""
	nextweek := make([]string, 0)
	for _, value := range inweek {
		index := strings.Index(value, instr)
		if index == 0 {
			nextweek = append(nextweek, value)
		}
	}
	fmt.Printf("%s\n", nextweek[0:])
	if len(nextweek) > 1 {
		fmt.Printf("intput next char:")
		fmt.Scanf("%s\n", &nextstr)
		nextstr = instr + nextstr
		outweek(nextstr, nextweek)
	}
}
