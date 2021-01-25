package gotest

import "fmt"

func ExampleHelloWorld() {
	HelloWorld()
	// Output: hello world
}

func ExampleAddNum() {
	a,b:=1,2
	fmt.Println(AddNum(a,b))
	// Output:3
}

func ExampleAddArray() {
	arr:=[]int{1,2,3}
	AddArray(arr)
}