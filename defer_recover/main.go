package main

import (
	"fmt"
)

//go语言捕获处理错误 defer+匿名函数+recover

func main() {
	r := test()
	fmt.Printf("%v %T\n", nil, nil)
	fmt.Println("结果是", r)
}

func test() int {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("test()函数出现错误!")
			fmt.Println("err=", err)

		}
	}()
	num1 := 10
	num2 := 0
	result := num1 / num2
	fmt.Println("result=", result)
	return result
}
