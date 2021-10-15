package 第二章数据类型

import (
	"fmt"
	"strconv"
	_ "time" // 引用但不使用，前面加_ 可以编译通过
)

func main1() {
	fmt.Println("hello world!")

	// 字符串和字相互转化
	var s string = "123"
	sToInt, _ := strconv.Atoi(s)
	fmt.Println(sToInt)
	intToS := strconv.Itoa(123)
	fmt.Printf("type:%T v:%v\n", intToS, intToS)

	//遍历字符串
	var s1 string = "go语言 42章经"
	fmt.Println(s1)
	s2 := "go语言 42章经  111"
	fmt.Printf("%T %v\n", s2, s2)
	for i, r := range s2 {
		fmt.Printf("%d : %v \n", i, string(r))
	}
}
