//custom_types.go
/*
自定义类型的学习
把一个短语中的标点符号移除掉
*/

package gosource

import (
	"fmt"
	"strings"
	"unicode"
)

//自定义RuneForRuneFunc为func(rune) rune 类型
type RuneForRuneFunc func(rune) rune

func main() {

	//声明一个移除标点符号的变量
	var removePunctuation RuneForRuneFunc
	//定义一个带标点符号的短语
	phrases := []string{"Day; dusk, and night.", "All day long", "大家好，同志们辛苦！"}
	//创建一个RuneForRuneFunc的匿名函数
	removePunctuation = func(char rune) rune {
		//unicode.Is()报告该符文是否在指定范围的表中。
		//unicode.Terminal_Punctuation 为带属性 Terminal_Punctuation 的 Unicode 字符集合。
		if unicode.Is(unicode.Terminal_Punctuation, char) {
			return -1
		}
		return char

	}

	processPhrases(phrases, removePunctuation)
}
func processPhrases(phrases []string, function RuneForRuneFunc) {
	//strings.Map返回字符串s的副本，其所有字符根据映射函数修改。 如果映射返回负值，则从没有替换的字符串中删除字符。
	for _, phrase := range phrases {
		fmt.Println(strings.Map(function, phrase))
	}

}
