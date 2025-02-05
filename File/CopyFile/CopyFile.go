package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

//复制文件示例

func copyFile1(y string, c string) {
	//y  源文件
	//c  复制的文件
	contend, err := ioutil.ReadFile(y)
	if err != nil {
		println("读取失败！")
	} else {
		err = ioutil.WriteFile(c, contend, 0666)
		if err != nil {
			println("复制失败！")
		} else {
			println("复制成功！")
		}
	}
}

// 逐行复制
func copyFile2(y string, c string) {
	//y  源文件
	//c  复制的文件
	file1, err := os.OpenFile(y, os.O_RDONLY, 0666)
	file2, err1 := os.OpenFile(c, os.O_RDWR|os.O_CREATE, 0666)
	defer func() {
		file1.Close()
		file2.Close()
	}()
	if err != nil || err1 != nil {
		fmt.Println("文件打开失败！")
	} else {
		reader := bufio.NewReader(file1)
		writer := bufio.NewWriter(file2)
		for {
			s, err2 := reader.ReadString('\n')
			writer.WriteString(s)
			if err2 == io.EOF {
				break
			}
		}
		writer.Flush()

	}
}

func main() {
	y := "./File/write.txt"
	c := "./File/copy.txt"
	c1 := "./File/copy1.txt"
	copyFile1(y, c)
	copyFile2(y, c1)
}
