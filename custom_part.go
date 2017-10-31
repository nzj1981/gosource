//custom_struct.go
/*
基于结构体的自定义类型，
类型的方法集是指可以被该类型的值调用的所有方法集合
方法表达式
方法表达式是一种高级特性
*/
package gosource

import (
	"fmt"
	"strings"
)

// 自定义结构体
type Part struct {
	Id   int    //聚合
	Name string //聚合
}

//基于自定义类型结构体的方法，其接收者为指针
func (part *Part) LowerCase() {
	part.Name = strings.ToLower(part.Name)
}

func (part *Part) UpperCase() {
	part.Name = strings.ToUpper(part.Name)
}

//基于自定义类型结构体的方法，其接收者为值
func (part Part) String() string {
	return fmt.Sprintf("<<%d  %q>>", part.Id, part.Name)
}

func (part Part) HasPrefix(prefix string) bool {
	return strings.HasPrefix(part.Name, prefix)
}

func main() {
	part := Part{5, "wrench"}

	part.UpperCase()

	part.Id += 11

	fmt.Println(part, part.HasPrefix("w"))

	//方法表达式展现
	fmt.Println("------方法表达式展现--------")
	asStringV := Part.String //有效签名: func(Part) string
	//asStringV()接受一个Part值作为其唯一的参数
	sv := asStringV(part)
	//hasPrefix()接受一个part值作为其第一个参数以及一个字符串作为其第二个参数
	hasPrefix := Part.HasPrefix //有效签名：func(Part,string) bool
	//asStringP()和lower()都接受一个*Part值作为其唯一参数
	asStringP := (*Part).String //有效签名：func(*Part) string
	ap := asStringP(&part)
	lower := (*Part).LowerCase
	lower(&part)
	fmt.Println(sv, ap, hasPrefix(part, "w"), part)
}
