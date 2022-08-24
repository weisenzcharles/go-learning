package main

import "fmt"

func main() {
	var a [5]int = [5]int{1, 2, 3, 4, 5}

	fmt.Println("数组 a 的值为：", a)

	var slice []int
	fmt.Println("切片 slice 的值为：", slice)

	slice = a[1:4]
	fmt.Println("切片 slice 的值为：", slice)
	fmt.Println("切片 slice 的长度为：", len(slice))
	fmt.Println("切片 slice 的容量为：", cap(slice))

	var b []int = make([]int, 5)
	fmt.Println("切片 b 的长度为：", len(b))
	fmt.Println("切片 b 的容量为：", cap(b))
	// 声明容量
	var c []int = make([]int, 5, 12)
	fmt.Println("切片 b 的长度为：", len(c))
	fmt.Println("切片 b 的容量为：", cap(c))

}
