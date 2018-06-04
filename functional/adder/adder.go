package main

import "fmt"

func adder() func(int) int {
	sum := 0//自由变量
	return func(v int) int {//函数体-》闭包（包括变量的引用）
		fmt.Println(v)
		sum += v//局部变量
		return sum
	}
}

//正统式函数编程
type iAdder func(int) (int, iAdder)

func adder2(base int) iAdder {
	fmt.Println(base)
	return func(v int) (int, iAdder){
		fmt.Println(v)
		return base + v, adder2(base + v)
	}
}

func main() {
	//a := adder()//返回的是一个函数体
	//fmt.Println(a);
	//for i := 0; i < 10; i++ {
	//	fmt.Printf("0 + 1 + ... + %d = %d\n", i, a(i))
	//}

	b := adder2(0)
	for i := 1; i < 10; i++ {
		var s int
		s, b = b(i)
		fmt.Printf("0 + 1 + ... + %d = %d\n", i, s)
	}
}
