package tests

import "fmt"

func HelloWorld() {
	fmt.Println("hello world")
	if false {
		fmt.Println("hi world")
	}
}

func PrintWords(s string) string {
	fmt.Println(s)
	return s
}
