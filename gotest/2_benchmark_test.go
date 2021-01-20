package gotest

import (
	"testing"
)

// 基准测试
func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add() // call tested func
	}
}

// 重置定时器
func BenchmarkAddArray(b *testing.B) {
	n := 100
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
	}

	b.ResetTimer() // reset timer

	for j := 0; j < b.N; j++ {
		AddArray(arr)
	}
}

var benchCases = []struct {
	name string
	arr  []int
}{
	{"one", []int{1, 2, 3}},
	{"two", []int{1, 2, 3}},
	{"three", []int{1, 2, 3}},
}

// Run() - subtests
func BenchmarkAddArray2(b *testing.B) {
	for _, arr := range benchCases {
		b.Run(arr.name, func(b *testing.B) { // run sub benchmark tests
			for i := 0; i < b.N; i++ {
				AddArray(arr.arr)
			}
		})
	}
}

// without Run() cannot make sub-benchmark test
func BenchmarkAddArray3(b *testing.B) {
	for _, arr := range benchCases {
		for i := 0; i < b.N; i++ {
			AddArray(arr.arr)
		}
	}
}

// 并行测试
func BenchmarkAddArray4(b *testing.B) {
	for _, c := range benchCases {
		arr := c
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				AddArray(arr.arr)
			}
		})
	}
}

//func TestCPU(t *testing.T){
//	fmt.Println(runtime.NumCPU())
//	fmt.Println(runtime.GOMAXPROCS(-1))
//}