package main

import "fmt"

func UserCode() {
	var name string
	var age int

	fmt.Println("Please input your name:")

	/*############*/
	fmt.Scanf("%s", &name)

	fmt.Println("Please input your age:")

	/*############*/
	fmt.Scan(age)

	/*############*/
	fmt.Printf("Hello, my name is %s, I am %d years old.\n", name, age)
}
