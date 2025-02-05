package main

import (
	"fmt"
	"strings"
)

// 字符串翻转1--遍历
func reverseString1(s string) string {
	var runes []rune = []rune(s) //专门处理中文字符串的一种类型

	for i := 0; i <= len(runes)/2; i++ {
		var r rune = runes[i]
		runes[i] = runes[len(runes)-1-i]
		runes[len(runes)-1-i] = r
	}
	return string(runes)

}

// 字符串翻转2--使用strings.Builder
func reverseString2(s string) string {
	var sb strings.Builder
	runes := []rune(s)
	for i := len(runes) - 1; i >= 0; i-- {
		sb.WriteRune(runes[i])

	}
	return sb.String()
}

func assert(a string, b string) {
	if a != b {
		fmt.Println("翻转失败！")
	} else {
		fmt.Println("翻转成功！")
	}
}
func main() {
	s := "aajjduyjgdyuqwiugauas你好！"
	fmt.Println("翻转前:", s)
	s1 := reverseString1(s)
	fmt.Println("函数1翻转后：", s1)
	assert(s, reverseString1(s1))

	s2 := reverseString2(s)
	fmt.Println("函数2翻转后：", s2)
	assert(s, reverseString2(s2))

}
