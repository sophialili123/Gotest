package main

import (
	"fmt"
	"os"
	"bufio"

)

//defer栈，先进后出
//调试关键点
func tryDefer()  {
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
	return
	panic("error occurred")
	fmt.Println(4)
}

func writeFile(filename string)  {
	file,err := os.Create(filename)
	if err != nil{
		panic(err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	f:=fib.Fibonacci()
	for i:=0;i<20 ;i++  {
		fmt.Fprintln(writer,f())
	}
}

func main() {
	tryDefer()
	writeFile("fib.txt")
}
