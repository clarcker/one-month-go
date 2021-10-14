package main

import (
	"fmt"
)

func main() {
	x := [3]int{1, 2, 3} //定义数组，x的类型为[3]int
	fmt.Printf("%T, %v\n", x, x)
	func(arr *[3]int) {
		(*arr)[0] = 7
		fmt.Println(*arr)
	}(&x)
	fmt.Println(x)
}
