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
	//t.Fatal("fatal")
	if err := HelloWorld(); err != nil{
		t.Fatalf("test fail, err:%s\n", err)
	}else{
		t.Log("test success")
	}
	t.Log("finish")
}

// table-driven test
func TestAddNum(t *testing.T) {
	testdatas:=[]struct{
		name string
		a int
		b int
		out int
	}{
		{"1+1",1,1,2},
		{"1-1",1,-1,0},
		{"0-0",0,0,1},
	}
	for _, data:=range testdatas{
		t.Run(data.name, func(t *testing.T) {
			if data.name == "0-0"{
				t.Skip("skip test") // 跳过测试
			}
			if data.name=="1-1"{
				//t.Error("test fail")
			}
			if res:=AddNum(data.a,data.b);res!=data.out{
				t.Errorf("got %d, want %d", res, data.out)
			}
		})
	}
}

// parallel test
// go test -run=TestLogParallel -parallel 8
func TestLogParallel1(t *testing.T) {
	t.Parallel() // 标记本测试函数并行
	testdatas:=[]struct{
		name string
	}{
		{"A 1"},
		{"A 2"},
		{"A 3"},
		{"A 4"},
		{"A 5"},
	}
	for _, data:=range testdatas{
		d:=data
		t.Run(d.name, func(t *testing.T) {
			t.Log(d.name)
			HelloWorld()
		})
	}
}

func TestLogParallel2(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
	}{
		{"B 1"},
		{"B 2"},
		{"B 3"},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Log(tt.name)
			HelloWorld()
		})
	}
}



//
//// without t.Run() cannot form sub-tests
//func TestPrintWords2(t *testing.T) {
//	for _, cases := range testCases {
//		if PrintWords(cases.s) == "a" {
//			t.Fatal("output is \"a\"")
//		}
//	}
//}

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

