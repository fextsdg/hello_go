package main

//读取文件示例

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// 读取文件内部封装了open和close
// 一次读取全部文件内容--适合读取小文件
func readFile1(f string) {
	fmt.Println("------读取方式1---------")
	content, err := ioutil.ReadFile(f)
	if err != nil {
		fmt.Println("文件打开失败,", err)
	} else {
		fmt.Println(string(content))
	}
}

// 通过缓冲区读取，可以按行读取
func readFile2(s string) {
	fmt.Println("------读取方式2---------")
	//打开文件
	file, err := os.Open(s)
	if err != nil {
		fmt.Println("文件打开失败！", err)
	} else {
		reader := bufio.NewReader(file)
		for i := 1; ; i++ {
			str, err := reader.ReadString('\n') //遇到换行就读取结束

			if err == io.EOF { //如果是文件末尾则退出读取循环
				fmt.Printf("第%d行，内容为：%s", i, str)
				break
			} else if err != nil {
				fmt.Println("读取出错，", err)
			} else {
				fmt.Printf("第%d行，内容为：%s", i, str)
			}
		}
	}
	defer func() {
		err := file.Close()
		if err != nil {
			fmt.Println("文件关闭失败！", err)
		} else {
			fmt.Println("文件关闭成功！")
		}
	}()
}
func main() {
	//方式1

	readFile1("./File/test.txt")

	//方式2
	readFile2("./File/test.txt")

}
