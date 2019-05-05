package main

/**
  Author by wpx
 */
import (
	"fmt"
)

// hello world
func main() {
	//noReturn();

	var amount = GetMoney()
	fmt.Println("money: ", amount)
	/**
	匿名函数，匿名函数后面跟"()"表示立即执行
	 */
	f := func(a, b int, z float64) bool {
		return a*b < int(z)
	}
	fmt.Println("f result=", f(1, 2, 5))
	fs := func(a, b int, z float64) bool {
		return a*b < int(z)
	}(1, 5, 2) //表示立即执行
	fmt.Println("fs result=", fs)
}

/**
函数没有返回值
 */
func noReturn() {
	var age int = 10; //声明一个变量并初始化
	var a = "RUNOOB"  //根据值自行判定变量类型
	fmt.Println("Hello, World!");
	fmt.Println(a);
	fmt.Println("age: ", age, "岁")
}

func GetMoney0() {
	fmt.Println("money")
	return
}

/**
 函数返回值有变量名
 */
func GetMoney() (_amount int) {
	_amount = 88
	fmt.Println("money: ", _amount)
	return;
}

/**
退出执行，指定返回值
 */
func GetMoney2() (_amount int) {
	fmt.Println("money: ", _amount)
	return 88
}

/**
退出执行，指定返回值和指定默认值
 */
func GetMoney3() (_amount int) {
	_amount = 99 //如果return后面没有指定返回值，就用赋给“返回值变量”的值
	fmt.Println("money: ", _amount)
	return 88 //如果return后面跟有返回值，就使用return后面的返回值
}
