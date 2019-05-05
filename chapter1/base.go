package main

/**
  Author by wpx
 */
import (
	"fmt"
	"reflect"
)

/**
	以大写字母开头的可以在其他包中访问，而小写的是私有的，不可以在其他包中访问
	iota 可以被编译期修改的常量，在每一个const出现时被重置为0，然后在下一个const出现之前，
	没出现一次iota,其代表的数字就会增加1
	 */
const (
	Sunday       = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	numberOfDays
)

func main() {
	declareParam()
}

/**
变量声明
 */
func declareParam() {
	var a float32 = 35
	/**
	自动推到类型为float64
	 */
	var b = 35.0
	fmt.Println("b type before convert=", reflect.TypeOf(b))
	/*
	用:=  必须保证变量没被声明过
	 */
	bconvert := float32(b)
	c := 10
	fmt.Println("a=", a, "; b=", b, "; c=", c)
	fmt.Println("a type=", reflect.TypeOf(a), "; b after convert type=", reflect.TypeOf(bconvert), "; c type=", reflect.TypeOf(c))

	/**
	多个变量可以一起声明
	 */
	var (
		v1 int    = 10
		v2 string = "v2 String"
	)
	fmt.Println("多个变量可以一起声明#v1 =", v1, "; v2 =", v2)
	fmt.Println("Thursday is ", Thursday)

	return
}
