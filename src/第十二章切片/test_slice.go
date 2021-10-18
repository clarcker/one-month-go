package 第十二章切片

import "fmt"

func get_slice(l int) []int {
	var date []int = []int{1, 2, 3, 4, 5, 6, 7}
	var ret = make([]int, l)
	copy(ret, date) //切片复制
	return ret
}

func TestSlice() {
	var arr = [5]int{11, 22, 33, 44, 55}
	s := arr[1:3:5] // 用a[low:high:max] 这种方式定义的切片，长度=high-low, 即3-1 = 2， 容量为max-low，即5-1=4
	fmt.Printf("s:%v len:%v cap:%v\n", s, len(s), cap(s))
	s1 := get_slice(3)
	fmt.Printf("s1:%v len:%v cap:%v\n", s1, len(s1), cap(s1))
	s2 := append(s1, 1, 2)
	fmt.Printf("s2:%v len:%v cap:%v\n", s2, len(s2), cap(s2))
	s3 := append(s2, s...)
	fmt.Printf("s3:%v len:%v cap:%v\n", s3, len(s3), cap(s3))

}
