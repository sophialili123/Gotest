package basic

/**
	内建变量类型
	bool,string
	(u)int,(u)int8,(u)int16,(u)int32,(u)int64,uintptr
	byte,run
	float32,float64,complex64,complex128
	Go语言只有强制类型转换没有隐式类型转换
	变量类型写在变量名之后
	编译器可推测变量类型
	没有char,只有rune
	原生支持复数类型
 */
import (
	"fmt"
	"math/cmplx"
	"math"
)

var (
	aa = 3
	ss = "kk"
	bb = true
)

func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

func variableTypeDeduction() {
	var a, b, c, s = 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}

func variableShorter() {
	a, b, c, s := 3, 4, true, "def" //第一次使用是定义，:=只能在方法函数中定义，GO中没有全局变量的作用
	fmt.Println(a, b, c, s)
}

func euler()  {
	c := 3 + 4i
	fmt.Println(cmplx.Abs(c))
	fmt.Println(cmplx.Pow(math.E,1i * math.Pi) + 1)
	fmt.Println(cmplx.Exp(1i*math.Pi) + 1)
}

/**
强制类型转换
不支持隐式转换
 */
func triangle()  {
	var a,b int = 3,4
	var c int
	c= int(math.Sqrt(float64(a*a+b*b)))
	fmt.Println(c)
}

func consts()  {
	//不一定要求全部大写，因为大小写在Go语言当中是有一定含义的
	//常量可以作为各种类型使用，下面就不需要强制类型转换了
	const (
		filename  = "abc.txt"
		a,b  =3,4
	)
	var c int
	c= int(math.Sqrt(a*a+b*b))
	fmt.Println(c)
}

func enums()  {
	const (
		cpp = iota
		_
		python
		golang
		javascript
	)
	fmt.Println(cpp,javascript,python,golang)

	const (
		b = 1 << (10*iota)
		kb
		mb
		gb
		tp
		pb
	)
	fmt.Println(b,kb,mb,gb,tp,pb)
}

func main() {
	fmt.Println("Hello world")
	variableZeroValue()
	variableInitialValue()
	variableShorter()
	variableTypeDeduction()

	euler()
	triangle()
	consts()
	enums();
}
