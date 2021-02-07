package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"unsafe"
)

func str2bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

func bytes2str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func main() {
	buf := make([]byte, 6)
	context := make([][]byte, 2024)
	f, err := os.Open("README.md")
	defer f.Close()
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(f)

	for {
		if _, err := reader.Read(buf); err == io.EOF {
			break
		}
		fmt.Println(context, buf)
	}

	bytes.Join(context, []byte(""))

	wirte := bufio.NewWriter(f)
	wirte.WriteString()


	//读
	//context := "这是一个测试文本文件"
	//conByte := []byte(context)
	//fmt.Println(conByte)
	//if err := ioutil.WriteFile("test.txt",conByte,0755);err != nil {
	//	fmt.Println(err)
	//	return
	//}

}
