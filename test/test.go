package test

import "fmt"

var A, _ int = cal(2)
var b = cal

func cal(num int) (int, int) {
	fmt.Println("cal!")
	return num, num
}

func init() {
	fmt.Println("test init!")
	fmt.Printf("b=%T\n", b)
	fmt.Println(b)

}
