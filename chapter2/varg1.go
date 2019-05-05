package main

import (
	"fmt"
)

/**
可变参数
 */
func myFunc(args ...int) {
	for _, arg := range args {
		fmt.Println(arg)
	}
}

/**
任意类型的不定参数
采用 interface{}方式
 */
func MyPrintf(args ...interface{}) {
	for _, arg := range args {
		switch arg.(type) {
		case int:
			fmt.Println(arg, "is an int value.")
		case string:
			fmt.Println(arg, "is a string value.")
		case int64:
			fmt.Println(arg, "is an int64 value.")
		default:
			fmt.Println(arg, "is an unknown type.")
		}
	}
}

func switchTest(i int) {
	switch i {
	case 0:
		fmt.Println("switchTest#0")
	case 1:
		fmt.Println("switchTest#1")
	case 2, 3, 4:
		fmt.Println("switchTest#2,3,4")
	case 5:
		fallthrough
	case 6:
		fmt.Println("5,6")
	default:
		fmt.Println("default")
	}
}

/**
switch后面的表达式可省略
 */
func switchTest2(i int) {
	switch {
	case 0 <= i && i <= 3:
		fmt.Println("switchTest2#0-3")
	case 4 <= i && i <= 6:
		fmt.Println("switchTest2#4-6")
	case i >= 7:
		fmt.Println("switchTest2#greater than 7")
	}
}

func main() {
	var (
		v1 int     = 1
		v2 int64   = 234
		v3 string  = "hello"
		v4 float32 = 1.234
	)
	MyPrintf(v1, v2, v3, v4)

	fmt.Println("#############switchTest################")
	switchTest(5)

	fmt.Println("#############switchTest2################")
	switchTest2(2)

	fmt.Println("#############myFunc################")
	myFunc(1, 2, 3, 4, 5, 6)
}
