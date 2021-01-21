package gotest

import (
	"testing"
)


// get started: 一个最简单的测试函数
//要点：
// 1-测试文件以 _test.go 结尾
// 2-测试函数形如 TestXxx(), Xxx 部分不可以小写开头
// 3-测试函数入参 testing.T， 这个数据结构用于管理测试状态，并提供格式化日志功能（t.Log， t.Errorf等）
func TestHelloWorld(t *testing.T) {
	if err := HelloWorld(); err != nil{
		t.Errorf("test fail, err:%s\n", err)
	}else{
		t.Log("test success")
	}
}

/*
运行测试函数有2种模式：
1. 本地目录模式

*/


var testCases = []struct {
	name string
	s    string
}{
	{"single", "a"},
	{"double", "aa"},
	{"triple", "aaa"},
}



// Run() - subtests
func TestPrintWords(t *testing.T) {
	for _, cases := range testCases {
		t.Run(cases.name, func(t *testing.T) {
			if PrintWords(cases.s) == "a" {
				t.Fatal("output is \"a\"")
			}
		})
	}
}

// Run() - subtests
func TestPrintWords4(t *testing.T) {
	for _, cases := range testCases {
		t.Run(cases.name, func(t *testing.T) {
			if PrintWords(cases.s) == "a" {
				t.Skip("skip s==a'")
			}
			if testing.Short()&&PrintWords(cases.s)=="aa"{
				t.Skip("skip short aa")
			}
		})
	}
}

// without t.Run() cannot form sub-tests
func TestPrintWords2(t *testing.T) {
	for _, cases := range testCases {
		if PrintWords(cases.s) == "a" {
			t.Fatal("output is \"a\"")
		}
	}
}

// 并行执行
func TestPrintWords3(t *testing.T) {
	for _, c := range testCases {
		cases := c // 不重新赋值会导致每个goroutine的cases取值相同
		t.Run(cases.name, func(t *testing.T) {
			t.Parallel()
			PrintWords(cases.s)
			//if PrintWords(cases.s) == "a"{
			//	t.Fatal("s is \"a\"")
			//}
		})
	}
}

// 不重新赋值会导致每个goroutine的cases取值相同
//func TestGorouteArrayIndex(t *testing.T) {
//	a := []int{1, 2, 3}
//	var c = make(chan interface{})
//	for i := range a {
//		go func() {
//			fmt.Println(i)
//			if i == 2 {
//				c <- 0
//			}
//		}()
//	}
//	<-c
//}

