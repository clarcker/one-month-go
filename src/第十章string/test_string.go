package 第十章string

import (
	"bytes"
	"fmt"
	"strings"
	"unicode/utf8"
)

func TestString() {
	s := "go 42章经"
	fmt.Printf("len:%v, utf-8.RunCountInString:%v", len(s), utf8.RuneCountInString(s))
}
func StingAdd() {
	//1)
	s1 := "hello " + "world!"
	fmt.Println(s1)
	//2)
	s2 := fmt.Sprintf("%d:%s", 2021, "year")
	fmt.Println(s2)
	//3)
	s3 := strings.Join([]string{"hello", "world"}, ", ")
	fmt.Println(s3)
	//4)
	var buffer bytes.Buffer
	buffer.WriteString("hello")
	buffer.WriteString(", ")
	buffer.WriteString("world!")
	fmt.Println(buffer.String())
	//5)
	var s5 strings.Builder
	s5.WriteString("hello")
	s5.WriteString(", ")
	s5.WriteString("world!")
	fmt.Println(s5.String())
}
