//interface_method.go
/*
接口的学习及接口方法的实现
*/
package gosource

import (
	"fmt"
	"io"
)

//定义接口，接口可以包含0到多个方法,接口的命名后面跟一个er
type Exchanger interface {
	Exchange()
}

//定义一个结构体，对接口方法的实现
type StringPair struct{ first, second string }

//字符串类型接口方法实现
func (pair *StringPair) Exchange() {
	pair.first, pair.second = pair.second, pair.first
}

//定义一个数字切片，对接口方法的实现
type Point [2]int

//数字类型接口方法实现
func (point *Point) Exchange() {
	point[0], point[1] = point[1], point[0]
}

//string方法 StringPair.String()方法可以满足fmt.Stringer接口的实现
func (pair StringPair) String() string {
	return fmt.Sprintf("%q + %q", pair.first, pair.second)
}

//定义StringPair.Read方法可以满足io.Reader接口的实现
func (pair *StringPair) Read(data []byte) (n int, err error) {
	if pair.first == "" && pair.second == "" {
		return 0, io.EOF
	}
	if pair.first != "" {
		n = copy(data, pair.first)
		pair.first = pair.first[n:]
	}
	if n < len(data) && pair.second != "" {
		m := copy(data[n:], pair.second)
		pair.second = pair.second[m:]
		n += m
	}
	return n, nil
}

// ToBytes()函数接受一个io.Reader（即任何包含签名为Read([]byte)(int,error)的方法的值,例如*os.File值）和一个大小限制，同时返回一个包含所读数据[]byte切片和一个error值.
func ToBytes(reader io.Reader, size int) ([]byte, error) {
	data := make([]byte, size)
	n, err := reader.Read(data)
	if err != nil {
		return data, err
	}
	return data[:n], nil
}

// 便利接口实现各种类型的方法(鸭子类型)
func exchangeThese(exchangers ...Exchanger) {
	for _, exchanger := range exchangers {
		exchanger.Exchange()
	}
}

func main() {
	jekyll := StringPair{"Henry", "Jekyll"}
	hyde := StringPair{"Edward", "Hyde"}
	point := Point{5, -3}
	fmt.Println("Before: ", jekyll, hyde, point)
	jekyll.Exchange()   //当做：(&jekyll).Exchange()
	hyde.Exchange()     //当做：(&hyde).Exchange()
	(&point).Exchange() //当做: point.Exchange()
	fmt.Println("After #1:", jekyll, hyde, point)
	//必须以显式地传入值的地址
	exchangeThese(&jekyll, &hyde, &point)
	fmt.Println("After #2:", jekyll, hyde, point)
	fmt.Println("-----Read方法使用------")
	const size = 16
	robert := &StringPair{"Robert L.", "Stevenson"}
	david := StringPair{"David", "Balfour"}
	for _, reader := range []io.Reader{robert, &david} {
		raw, err := ToBytes(reader, size)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%q\n", raw)
	}
}
