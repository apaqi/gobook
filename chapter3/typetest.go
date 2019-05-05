package main

import "fmt"

/**
定义Integer类
 */
type Integer int

/**
定义矩形的结构体
 */
type Rect struct {
	x, y          float64
	width, height float64
}
/**
Area 首字母大写，则对其他包也可见，如果首字母小写，则只能在该类所在的包内访问
 */
func (r *Rect) Area() float64 {
	return r.width * r.height
}

/**
go 语言用NewXXX来命名，来表示“构造函数”
 */
func NewRect(x, y, width, height float64) *Rect {
	return &Rect{x, y, width, height}
}

/**
给类型添加方法
 */
func (a Integer) Less(b Integer) bool {
	return a < b
}

/**
只有需要修改对象的时候，才必须用指针，go类型都是基于值传递的，要想修改变量的值，必须传递指针
 */
func (a *Integer) Add(b Integer) {
	*a += b
}

func main() {
	fmt.Println("##############Integer##############")

	var a Integer = 1
	if a.Less(2) {
		fmt.Println("a=", a, ",a Less than 2")
	}
	a.Add(2)
	fmt.Println("a=", 2)
	fmt.Println("##############Rect##############")
	/**
	类型初始化方式
	1.ra := new (Rect)
	2.ra := &Rect{}
	3.ra := &Rect{0,0,100,200}
	4.ra := &Rect{width:100,height:200}
	 */
	ra := &Rect{width: 100, height: 200}
	fmt.Println("Rect area=", ra.Area())
	//采用构造函数方式初始化
	rb := NewRect(0,0,10,30)
	fmt.Println("Rect area=", rb.Area())
}
