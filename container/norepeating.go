package main

import "fmt"

/**
	寻找最长字符串的子串
 */

func lengthOfNonRepeatingSubStr(s string) int {
	lastOccurred := make(map[rune]int)//byte
	start := 0
	maxLength := 0

	for i, ch := range []rune(s) {//[]byte(s)
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}

		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}

func main() {
	fmt.Println(lengthOfNonRepeatingSubStr("abcabacdeb"))
	fmt.Println(lengthOfNonRepeatingSubStr("这里对的萨菲")) //不正确
	fmt.Println(lengthOfNonRepeatingSubStr("dsafdsaf这里对")) //一个汉字算一个字节
}
