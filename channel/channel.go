package main

//管道实践案例
//类似队列 先进先出

import (
	"fmt"
)

func main() {
	var intchan chan int         //声明
	intchan = make(chan int, 10) //初始化   10是管道长度，只能存放10个数据
	intchan <- 10                //赋值
	num := 20
	intchan <- num //赋值
	intchan <- 2
	fmt.Printf("intchan:%v,the length of intchan is %v,the captity of intchan is %v\n", intchan, len(intchan), cap(intchan))

	num1 := <-intchan //取值
	fmt.Println(num1)
	fmt.Println(<-intchan)
	fmt.Printf("intchan:%v,the length of intchan is %v\n", intchan, len(intchan))

	//关闭管道--关闭后不能写入但能读出数据
	close(intchan)
	fmt.Println(<-intchan)
	fmt.Printf("intchan:%v,the length of intchan is %v\n", intchan, len(intchan))

}
