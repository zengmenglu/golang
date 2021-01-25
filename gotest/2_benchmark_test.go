package gotest

import (
	"testing"
)

// 基准测试
func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld() // call tested func
	}
}

// 重置定时器
func BenchmarkAddArray(b *testing.B) {
	// 测试前的数据准备
	n := 10000
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
	}

	// reset timer
	b.ResetTimer()

	for j := 0; j < b.N; j++ {
		AddArray(arr)
	}
}

// table-driven-test 构造子测试
func BenchmarkAddArray2(b *testing.B) {
	var benchCases = []struct {
		name string
		arr  []int
	}{
		{"one", []int{1}},
		{"two", []int{1, 2}},
		{"three", []int{1, 2, 3}},
	}
	for _, arr := range benchCases {
		b.Run(arr.name, func(b *testing.B) { // run sub benchmark tests
			for i := 0; i < b.N; i++ {
				AddArray(arr.arr)
			}
		})
	}
}

// 并行测试
func BenchmarkAddArray4(b *testing.B) {
	var benchCases = []struct {
		name string
		arr  []int
	}{
		{"one", []int{1, 2, 3}},
		{"two", []int{1, 2, 3}},
		{"three", []int{1, 2, 3}},
	}
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