package di_4

import "fmt"

func main2() {
	const (
		a = iota
		b
		c
		_
		e
		f = 10
		g
	)
	fmt.Println(a, b, c, e, f, g)
	main1()
}
func main1() {
	if a := 1; false {
	} else if b := 2; false {
	} else if c := 3; false {
	} else {
		println(a, b, c)
	}
}
