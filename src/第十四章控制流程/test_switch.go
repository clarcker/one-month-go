package 第十四章控制流程

import "fmt"

func TestSwitch() {
	a := 3
	switch a {
	case 1:
		fmt.Println("a=1")
		fmt.Println("----1-----")
	case 2:
		fmt.Println("a=2")
		fmt.Println("----2----")
	case 3, 4:
		fmt.Println("a=", a)
		fmt.Println("a", "----", a, "----")
	default:
		fmt.Println("a=", a)
		fmt.Println("----", a, "----")
	}
}
