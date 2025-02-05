package main

import "fmt"

//数组初始化示例

func main() {
	var a [3]int = [3]int{1, 2, 3}
	fmt.Println(a)
	var b = [3]int{1, 4, 7}
	fmt.Println(b)
	var c = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(c)
	var d = [...]int{1: 3, 0: 2, 5: 6, 3: 22, 4: 33}
	fmt.Println(d)
}
