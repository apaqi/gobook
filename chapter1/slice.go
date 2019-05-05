package main

import (
	"fmt"
)

/**
slice 的创建有两种方法
 */

func createSlice() {
	//直接创建
	test := []int{2, 3}
	fmt.Println("test values=", test)
	//基于已有数组创建, [first:last] 操作
	var array1 = [3]int{4, 5, 6}
	var slice1 = array1[:2] //表示前2个元素
	fmt.Println("slice1 values=", slice1)

	//给slice添加新值 append
	test2 := append(slice1, 1, 2, 3) //4, 5,1, 2, 3
	fmt.Println("after apend test2 values=", test2)
	test3 := append(test2, slice1...)
	fmt.Println("after apend slice1 to test2 values=", test3)

	//使用make，而且通常我们使用 make创建的情况比较多
	myslice := make([]int, 5, 10)
	fmt.Println("myslice values=", myslice)
	fmt.Println("myslice len=", len(myslice), ", myslice cap=", cap(myslice))

}
func remove(arr []int, i int) []int {
	return append(arr[:i], arr[i+1:]...)
}

func testRemove() {
	a := []int{1, 2, 3, 4, 5}

	a = remove(a, 0)
	fmt.Println(a) // [2, 3, 4, 5]

	a = remove(a, 3)
	fmt.Println(a) // [2, 3, 4]

	a = remove(a, 1)
	fmt.Println(a) // [2, 4]
}

func main() {
	createSlice()
	fmt.Println("################testRemove###################3")
	testRemove()
}
