package 第十三章字典

import "fmt"

func TestMap() {
	maps := make(map[string]string, 5) //可不指定大小
	maps["name"] = "cjw"
	maps["age"] = "20"
	if _, ok := maps["age"]; ok { //判断key是否存在
		fmt.Println(maps["age"])
		delete(maps, "age")
	} else {
		fmt.Println("not in")
	}
}
func TestRange() {
	testSlice := []int{1, 2, 3}
	fmt.Println(testSlice)
	for i, date := range testSlice {
		fmt.Println(i, date)
		date *= 10 //date 只是一份拷贝，改变不会影响原始值
	}
	fmt.Println(testSlice)
	for i, date := range testSlice {
		fmt.Println(i, date)
		testSlice[i] *= 10
	}
	fmt.Println(testSlice)
}
