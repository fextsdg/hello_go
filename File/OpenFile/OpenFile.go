package main

//打开-关闭文件示例

import (
	"fmt"
	"os"
)

func main() {
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("获取当前工作目录失败，", err)
	} else {
		fmt.Println("当前工作目录是：", currentDir)
	}
	file, err := os.Open("./File/test.txt")
	if err != nil {
		fmt.Println("文件打开失败，", err)
	} else {
		fmt.Println("文件打开成功！")
		fmt.Println(*file)
	}

	defer func() {
		err = file.Close()
		if err != nil {
			fmt.Println("文件关闭失败，", err)
		} else {
			fmt.Println("文件关闭成功！")

		}
	}()

}
