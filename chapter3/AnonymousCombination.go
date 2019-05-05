package main

import "fmt"

/**
匿名组合
 */
type Base struct {
	Name string
}

/**
组合Base
 */
type Foo struct {
	Base
	age int
}

func (base *Base) Foo() {
	fmt.Println("Base Foo#name=",base.Name)
}

func (base *Base) Bar() {
	fmt.Println("Base Bar#name=",base.Name)
}
/**
继承并重写Base的Bar方法
 */
func (foo *Foo) Bar() {
	foo.Base.Bar()
	fmt.Println("Foo Bar name=",foo.Name,";age=",foo.age)
}

func main()  {
	x:=Base{"wpx"}
	x.Bar()
	x.Foo()

	y:=Foo{Base{"wpx"},30}
	y.Bar()
	y.Foo()

}
