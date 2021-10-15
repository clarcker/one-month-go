package main

import (
	"fmt"
	"one-month-go/import_test"
	"one-month-go/第十一章数组"
	"one-month-go/第十二章切片"
	"one-month-go/第十章string"
)

func test10() {
	第十章string.TestString()
	第十章string.StingAdd()
}
func test11() {
	第十一章数组.TestArray()
}

func test12() {
	第十二章切片.TestSlice()
}

func main() {
	fmt.Println("hello world")
	fmt.Println(import_test.TestA(1, 2))
	test11()
}
