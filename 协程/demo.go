package main

//协程简单案例
//通过函数创建协程

//主死从随--主函数消亡，协程也立即消亡（不管是否执行完成）

import (
	"fmt"
)

import "sync"

var wg sync.WaitGroup //主函数等待所有协程执行完之后，再消亡的一个控制
func test() {
	defer wg.Done()
	fmt.Println("这是一个普通的协程！")
}
func main() {
	wg.Add(2)
	//匿名协程
	go func() {
		defer wg.Done()
		fmt.Println("这是一个匿名函数创建的协程！")
	}()

	go test()
	//time.Sleep(time.Second)
	fmt.Println("主函数！")
	wg.Wait() //等待协程执行完之后再退出主函数

}
