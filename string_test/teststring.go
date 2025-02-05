package main

//测试字符串相关函数
import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	str := "hello world--你好世界"
	fmt.Printf("the length of str is %d\n", len(str))

	//遍历字符串方式Ⅰ
	for i, v := range str {
		fmt.Printf("the index of %c is %d\n", v, i)

	}
	fmt.Println("--------------------------------------------------------------------------------------------")
	//切片方式遍历字符串
	str1 := []rune(str)
	for i := 0; i < len(str1); i++ {
		fmt.Printf("the index of %c is %d\n", str1[i], i)
	}

	fmt.Println("--------------------------------------------------------------------------------------------")

	//字符串转整数
	var str2 string = "123"
	var a, err = strconv.Atoi(str2)
	fmt.Println(a, err)

	//整数转字符串
	fmt.Println(strconv.Itoa(a))

	//不区分大小写比较字符串
	fmt.Printf("不区分大小写比较字符串Hello ,hello is %t\n", strings.EqualFold("Hello", "hello"))
	fmt.Printf("区分大小写比较字符串Hello ,hello is %t\n", "Hello" == "hello")

	fmt.Println("------------------------------------------------------------------------------------------")

	//计算字符串中第一个出现的子串位置:
	fmt.Println(strings.Index("hello world", "lo"))
	//计算子串在字符串中出现的次数:
	fmt.Println(strings.Count("hello world", "lo"))

	fmt.Println("----------------------------------------------------------------------------------------")

	//字符串替换--最后一个参数代表替换几次，-1代表全替换
	fmt.Println(strings.Replace("goandjavagogo", "go", "golang", -1))

	fmt.Println("----------------------------------------------------------------------------------------")
	//去掉字符串两边的空格
	fmt.Println(strings.TrimSpace("  hello world  "))
	fmt.Println("----------------------------------------------------------------------------------------")
	//去掉字符串两边指定字符
	fmt.Println(strings.Trim("~hello world~", "~"))

	//去掉字符串左边的字符
	fmt.Println(strings.TrimLeft("~hello world~", "~"))

	//去掉字符串右边的字符
	fmt.Println(strings.TrimRight("~hello world~", "~"))
	fmt.Println("----------------------------------------------------------------------------------------")
	//判断字符串头部是否包含某个字符
	fmt.Println(strings.HasPrefix("demo212ww.go", "demo"))
	//判断字符串尾部是否包含某个字符串
	fmt.Println(strings.HasSuffix("demo.go", "go"))

	//转成大写
	fmt.Println(strings.ToUpper("hello woRld"))
	//转成小写
	fmt.Println(strings.ToLower("HELLO WorLD"))
}
