package main

import (
	"fmt"
	"one-month-go/import_test"
	"one-month-go/第十一章数组"
	"one-month-go/第十三章字典"
	"one-month-go/第十二章切片"
	"one-month-go/第十四章控制流程"
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

func test13() {
	第十三章字典.TestMap()
	第十三章字典.TestRange()
}
func test14() {
	第十四章控制流程.TestSwitch()
}
func main() {
	fmt.Println("hello world")
	fmt.Println(import_test.TestA(1, 2))
	test14()
}
