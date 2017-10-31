// chapter32.go
/*
改变控制台颜色

题目：Press any key to change color, do you want to try it. Please hurry up!
1.程序分析： 使用Win API GetStdHandle和SetConsoleTextAttribute
*/

package main

import (
	"fmt"
	"syscall"
)

type (
	HANDLE uintptr
	WORD   uint16
	DWORD  uint32
)

const (
	STD_OUTPUT_HANDLE    = 0xFFFFFFF5
	FOREGROUND_BLUE      = 0x01
	FOREGROUND_GREEN     = 0x02
	FOREGROUND_RED       = 0x04
	FOREGROUND_INTENSITY = 0x08
	BACKGROUND_BLUE      = 0x10
	BACKGROUND_GREEN     = 0x20
	BACKGROUND_RED       = 0x40
	BACKGROUND_INTENSITY = 0x80
)

var (
	modkernel32 = syscall
)
