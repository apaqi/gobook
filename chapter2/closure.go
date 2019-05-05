package main

import "fmt"

/**
闭包
 */
func main() {
	var j int = 5i
	a := func() (func()) {
		var i int = 10i
		return func() {
			fmt.Printf("i, j: %d, %d\n", i, j);
		}
	}()

	a()

	j *= 2

	a()
}
