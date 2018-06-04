package main

import (
	"io"
	"fmt"
	"strings"
	"bufio"
)

/**
	斐波那契数
	1,  1,  2,  3,  5,  8,  14, ...
        a,  b
            a,  b
 */
func Fibonacci() intGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

//只要是类型就能实现接口
type intGen func() int

//(g intGen)可放前面，也可放后面，主要影响调用方式
/**
为什么有的函数名前面有输入参数，而一些却没有，它们是否有差别？
确实有差别，没有输入参数，是一般的函数；有输入参数，是结构的方法，输入参数叫做“方法接收者”！
GO语言没有类，方法都定义在结构上了！！
 */
func (g intGen) Read(p []byte) (n int, err error) {
	next := g()
	if next > 10000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)
	return strings.NewReader(s).Read(p)
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	f := Fibonacci()

	//fmt.Println(f())//1
	//fmt.Println(f())//1
	//fmt.Println(f())//1
	//fmt.Println(f())//1
	//fmt.Println(f())//1
	//fmt.Println(f())//1
	//fmt.Println(f())//1
	//fmt.Println(f())//1

	printFileContents(f)

}
