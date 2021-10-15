package 第十一章数组

import (
	"fmt"
)

func TestArray() {
	var arr1 = new([5]int)
	arr2 := arr1 //指针传递
	arr1[2] = 100
	fmt.Printf("arr1:%v arr1:%T arr2:%v arr2:%T\n", arr1[2], arr1, arr2[2], arr2)

	var arr3 = [...]int{1, 2, 3, 4}
	arr4 := arr3 //值传递
	arr3[2] = 100
	fmt.Printf("arr3:%v arr3T:%T arr4:%v arr4T:%T\n", arr3[2], arr3, arr4[2], arr4)

	//多维数组遍历
	var nArray = [...][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Printf("nAarray len:%d\n", len(nArray))
	for i, arr := range nArray {
		fmt.Println(i, arr)
		fmt.Println("-----")
		for j, array := range arr {
			fmt.Println(j, array)
		}
	}
}
