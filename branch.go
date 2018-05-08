package main

import (
	"io/ioutil"
	"fmt"
)

func grade(score int) string {
	g := ""
	//switch可以没有表达式，直接在case里面加表达式即可，默认每个case后面都有一个break
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("Wrong score:%d",score))//输出错误，并中断程序
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "D"
	case score < 100:
		g = "A"
	}
	return g;
}

func main() {
	const filename = "abcd.txt"
	/**
		if的条件里可以赋值
		if的条件里赋值的变量作用域就在这个if语句里
	 */
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}

	fmt.Println(
		grade(0),grade(59),grade(60),grade(100),grade(101),
	)
}
