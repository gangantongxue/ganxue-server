package main

import "fmt"

func UserCode() {
	var a, b int
	fmt.Scan(&a, &b)
	sub, isOddNum := add(a, b)
	fmt.Println("a + b =", sub, "isOddNum:", isOddNum)
}

/*请在下方定义函数add()*/
func add(a, b int) (int, bool) {
	c := a + b
	if c%2 == 0 {
		return c, false
	} else {
		return c, true
	}
}
