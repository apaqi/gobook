package main

import "fmt"

//1. nil 是不能比较的
func nil_not_eq() {
	//fmt.Println(nil == nil) //invalid operation: nil == nil (operator == not defined on nil)
}

//2. 默认 nil 是 typed 的
func default_typed() {
	fmt.Printf("%T", nil) //<nil>
	//print(nil)            //use of untyped nil
}

//3. 不同类型 nil 的 address 是一样的
func address() {
	var m map[int]string
	var ptr *int
	fmt.Printf("%p", m)//0x0
	fmt.Printf("%p", ptr)//0x0
}
//4. 不同类型的 nil 是不能比较的
func dif_type_cannot_eq()  {
	var ptr *int
	var ptr2 *int
	if ptr == ptr2 { //类型相同，可以比较
		fmt.Printf("ptr == ptr2")
	}
	//var m map[int]string
	//fmt.Printf(m == ptr)//invalid operation: m == ptr (mismatched types map[int]string and *int)
}
//5.nil 是 map，slice，pointer，channel，func，interface 的零值
//zero value 是 go 中变量在声明之后但是未初始化被赋予的该类型的一个默认值。
func zero_value()  {
	var m map[int]string
	var ptr *int
	var c chan int
	var sl []int
	var f func()
	var i interface{}
	fmt.Printf("%#v\n", m)
	fmt.Printf("%#v\n", ptr)
	fmt.Printf("%#v\n", c)
	fmt.Printf("%#v\n", sl)
	fmt.Printf("%#v\n", f)
	fmt.Printf("%#v\n", i)
}


func main() {
	zero_value()
}
