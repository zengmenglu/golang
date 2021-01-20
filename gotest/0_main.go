package gotest

import "fmt"

func HelloWorld() error{
	fmt.Println("hello world")
	return nil
}

func Add() {
	var n int
	var nums []int
	for i := 0; i < 100; i++ {
		n += i
		nums = append(nums, i)
	}
}

func AddArray(arr []int) {
	n := 0
	for i := range arr {
		n += arr[i]
	}
}

func Parallel() {

}



func PrintWords(s string) string {
	fmt.Println(s)
	return s
}
