package gotest

import "fmt"

func HelloWorld() error{
	for i:=0;i<1000000;i++{
		fmt.Sprintf("%d",i)
	}
	fmt.Println("hello world")
	return nil
}

func AddNum(a,b int)int{
	return a+b
}

func AddArray(arr []int) {
	n := 0
	for i := range arr {
		n += arr[i]
	}
}

