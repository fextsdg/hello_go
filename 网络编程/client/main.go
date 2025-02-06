package main

//网络编程实践--客户端

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

// 客户端将终端获取的数据发送给服务器端
func portSomething(conn net.Conn) {
	reader := bufio.NewReader(os.Stdin) //通过标准输入获取终端数据
	fmt.Println("请输入要发送的数据：")
	str, err := reader.ReadString('\n') //遇到换行结束
	if err != nil {
		fmt.Printf("获取终端输入失败！err:%v\n", err)
		return
	}
	n, err1 := conn.Write([]byte(str)) //向服务器端发送数据，n代表发送数据的字节数
	if err1 != nil {
		fmt.Printf("向服务器端发送数据失败！err:%v\n", err1)
		return
	}
	fmt.Printf("发送数据成功！共%d字节数据！\n", n)

}
func main() {
	//客户端发送建立链接请求
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Printf("客户端发送建立链接请求！err:%v\n", err)
		return
	}
	fmt.Println("请求成功！")
	fmt.Printf("服务器端信息：%v\n", conn.RemoteAddr().String())
	portSomething(conn)

}
