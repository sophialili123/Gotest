package basic

/**
	数字十进制进行转换
 */

import (
	"fmt"
	"strconv"
	"os"
	"bufio"
)

func convertToBin(n int) string {
	result := ""
	for ; n > 0;n /=2 {
		lsb := n%2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func printFile(filename string)  {
	file,err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	//相当于while,Go语言中没有while
	for scanner.Scan(){
		fmt.Println(scanner.Text())
	}
}

//死循环
func forever()  {
	for{
		fmt.Println("abc")
	}
}
func main() {
	fmt.Println(
		convertToBin(5),//101
		convertToBin(13),//1011 -> 1101
	)
	printFile("abcd.txt")
}
