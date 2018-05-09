package container

import "fmt"

/**
	[]int 切片
	[5]int 数组
	数组是值类型
	加*定义为指针
 */
func printArray(arr *[5]int)()  {
	//(*arr)[0] = 100
	arr[0] = 100
	for i,v := range arr{
		fmt.Println(i,v)
	}

}

func printArray1(arr []int)()  {
	//(*arr)[0] = 100
	arr[0] = 100
	for i,v := range arr{
		fmt.Println(i,v)
	}

}

func main() {
	var arr1 [5]int
	arr2 := [3]int{1,3,5}//此方式需要初始值
	arr3 := [...]int{2,4,6,8,10}//让程序帮我们计算个数
	var grid [4][5]int//多维数组，4行五列

	//方式一
	for i := 0;i<len(arr3);i++{
		fmt.Println(arr3[i])
	}

	//range,方式二
	for i,v := range arr3{
		fmt.Println(i,v)
	}

	fmt.Println(arr1,arr2,arr3,grid)
	printArray1(arr1[:])//调用func(arr [10]in)会拷贝数组
	//printArray(arr2)
	printArray(&arr3)

	fmt.Println(arr1)
}
