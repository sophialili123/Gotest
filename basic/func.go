package main

import (
	"fmt"
	"reflect"
	"runtime"
	"math"
)

//返回两个值的情景一
//仅用于非常简单的函数
func eval(a,b int,op string) (int, error)  {
	switch op {
	case "+":
		return a+b,nil
	case "-":
		return a-b,nil
	case "*":
		return a*b,nil
	case "/":
		q,_ := div(a,b)
		return q,nil
	default:
		//panic("unsupported operation:"+ op)
		return 0,fmt.Errorf("unsupported operation:%s",op)
	}
}

func div(a,b int)(q,r int)  {
	//return a/b,a%b
	q=a/b
	r=a%b
	return
}

func apply(op func(int,int) int,a,b int) int  {
	p:= reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p)
	fmt.Println("Calling function %s with args " + "(%d,%d)",opName,a,b)
	return op(a,b)
}

func pow(a,b int) int {
	return int(math.Pow(float64(a),float64(b)))
}

func main() {
	fmt.Println(eval(3,4,"*"))
	fmt.Println(eval(3,4,"-"))
	fmt.Println(eval(3,4,"X"))

	q,r:= div(13,3)
	fmt.Println(q,r)
	fmt.Println(apply(pow,3,4))
}
