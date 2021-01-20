package gotest

import (
	"testing"
)

// get started: 一个最简单的单元测试用例
func TestHelloWorld(t *testing.T) {
	if err := HelloWorld(); err != nil{
		t.Fail()
		t.Errorf("test fail, err:%s\n", err)
	}else{
		t.Log("test success")
	}
	t.Log("test finish")
}

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

