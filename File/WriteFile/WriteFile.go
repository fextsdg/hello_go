package main

//写文件示例
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/*
*
os.O_RDONLY: 以只读模式打开文件。
os.O_WRONLY: 以只写模式打开文件。
os.O_RDWR: 以读写模式打开文件。
os.O_CREATE: 如果文件不存在，则创建文件。
os.O_TRUNC: 如果文件存在，打开文件时将文件长度截断为 0。
os.O_APPEND: 以追加模式打开文件，即写操作会追加到文件末尾。
os.O_EXCL: 在创建文件时，如果文件已存在，则返回错误。
os.O_SYNC: 同步文件的读写操作，即写操作会等到磁盘实际写入完成。
os.O_DIRECT: 直接 I/O，不经过系统缓存。
*/
func writeFile(f string) {
	file, err := os.OpenFile(f, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("文件打开失败！", err)
	} else {
		writer := bufio.NewWriter(file)
		for i := 0; i < 10; i++ {
			writer.WriteString("写入" + strconv.Itoa(i) + "\n")
		}
		writer.Flush()
		fmt.Println("文件写入完成！")
	}
	defer func() {
		err = file.Close()
		if err != nil {
			fmt.Println("文件关闭失败！", err)
		} else {
			fmt.Println("文件关闭成功！")
		}
	}()
}

func main() {
	writeFile("./File/write.txt")
}
