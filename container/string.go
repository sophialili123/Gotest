package container

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Yes我爱吴若涵"
	fmt.Printf("%X\n",[]byte(s))
	for _,b:=range []byte(s) {
		fmt.Printf("%X ",b)//UTF-8
	}
	fmt.Println()

	for i,ch:=range s {//ch is a rune
		fmt.Printf("(%d %X) ",i,ch)//unicode
	}
	fmt.Println()
	fmt.Println("Rune count: ",utf8.RuneCountInString(s))

	bytes := []byte(s)
	for len(bytes) > 0{
		ch,size:=utf8.DecodeRune(bytes)
		bytes= bytes[size:]
		fmt.Printf("%c-",ch)
	}

	fmt.Println()
	for i,ch:=range  []rune(s){
		fmt.Printf("(%d %c)",i,ch)
	}
	fmt.Println()
}
