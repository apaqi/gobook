package main

import (
	"fmt"
	"reflect"
)

func main() {
	arrayfunc()
}

func arrayfunc() {
	/**
	数组定义及初始化
	 */
	var array1 [10]int
	fmt.Println("数组#array1 type =", reflect.TypeOf(array1))
	fmt.Println("###########################")

	var array2 = [5]int{1, 2, 3, 4, 5}   // 这种方式，既初始化变量，又带了初始化值，数组长度已经定义好
	var array3 = [...]int{1, 2, 3, 4, 5} // 这种方式，既初始化变量，也是带了初始值，数组长度，根据初始值的个数而定，也就是五个
	array4 := [5]int{1, 2, 3, 4, 5}      // 这种方式，省去 var 关键词，将初始化变量和赋值，放在一起操作,这种方式简单，明了。
	array5 := [...]int{1, 2, 3, 4, 5}    // 这种方式，既初始化变量，也是带了初始值，数组长度，根据初始值的个数而定，也就是五个

	/**
	数组遍历
	 */
	for i := 0; i < len(array2); i++ {
		fmt.Println("数组#array2 value =", array2[i])
	}
	fmt.Println("###########################")
	for i := 0; i < len(array2); i++ {
		fmt.Println("数组#array3 value =", array3[i])
	}
	/*
	数组支持范围遍历
	 */
	fmt.Println("###########################")
	for i, v := range array4 {
		fmt.Println("数组#array4 index =", i, "value=", v)
	}
	fmt.Println("###########################")
	/**
	如果返回值不需要，可用"_" 表示
	 */
	for _, v := range array5 {
		fmt.Println("数组#array5 value=", v)
	}
	return
}
