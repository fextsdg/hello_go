package main

//网络编程实践--服务器端
import (
	"fmt"
	"net"
	"sync"
)

var wg sync.WaitGroup

// 用于处理服务器接收的数据
func process(conn net.Conn) {
	defer func() {
		err := conn.Close()
		if err != nil {
			fmt.Printf("关闭连接失败！err:%v\n", err)
			return
		}
		wg.Done()
	}()
	buf := make([]byte, 1024) //建立一个byte切片用于接收数据
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Printf("服务器监接收数据失败！err:%v\n", err)
		return
	}
	fmt.Printf("接收了%d字节的数据，信息为:%v \n!", n, string(buf[:n]))
}

func main() {
	//建立服务器监听
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Printf("服务器监听建立失败！err:%v\n", err)
		return
	}
	for {
		fmt.Println("等待接收请求....")
		//接收请求连接
		conn, err1 := listener.Accept()
		if err1 != nil {
			fmt.Printf("服务器接收请求失败！err:", err1)
			return
		}
		fmt.Printf("服务器接收的客户端信息:%v!\n", conn.RemoteAddr().String())
		wg.Add(1)
		go process(conn)
	}

}
