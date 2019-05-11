package main

import "fmt"

func Count(ch chan int) {
	fmt.Println("Counting")
	//向channel写入值方式： ch <- value
	ch <- 1
}

func main() {
	//创建缓存大小为10 的channel，在缓存区不被填满之前，都不会堵塞
	chs := make([]chan int, 10)
	for i := 0; i < 10; i++ {
		chs[i] = make(chan int)
		go Count(chs[i])
	}
	for _, ch := range chs {
		//从channel中读取值： value := <-ch
		value :=<-ch
		fmt.Println(value)
	}
}
