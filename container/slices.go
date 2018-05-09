package container

import "fmt"

func updateSlice(s []int) {
	s[0] = 100
}

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	//半开半闭，从第2个到第6个
	s := arr[2:6]

	fmt.Println("arr[2:6] = ", s)
	fmt.Println("arr[:6] = ", arr[:6])
	fmt.Println("arr[2:] = ", arr[2:])
	fmt.Println("arr[:] = ", arr[:])

	updateSlice(s)
	fmt.Println(s)
	fmt.Println("arr[:] = ", arr[:])

	fmt.Println("After Reslice:")
	fmt.Println(s)
	s = s[:5]
	fmt.Println(s)
	s = s[2:]
	fmt.Println(s)

	fmt.Println("Test:")
	arr1 := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := arr1[2:6]
	fmt.Println(s1)
	s2 := s1[3:5]//[s1[3],s1[4]]
	fmt.Println(s2)

	s3:=append(s2,10);
	s4:=append(s3,11);
	s5:=append(s4,12);
	fmt.Println("s3,s4,s5",s3,s4,s5)
	fmt.Println("arr",arr)
}
