package parallel

import (
	"testing"
)
var testCases = []struct{
	name string
	s string
}{
	{"single","a"},
	{"double","aa"},
	{"triple","aaa"},
}

//func TestPrintWord(t *testing.T) {
//	for _,c:=range testCases{
//		cases:=c //
//		t.Run(cases.name, func(t *testing.T) {
//			t.Parallel()
//			if PrintWord(cases.s) == "a"{
//				t.Fatal("s is \"a\"")
//			}
//		})
//	}
//}
//
//func TestGorouteArrayIndex(t *testing.T){
//	a:=[]int{1,2,3}
//	var c = make(chan interface{})
//	for i:=range a{
//		go func(){
//			fmt.Println(i)
//			if i==2{
//				c<-0
//			}
//		}()
//	}
//	<-c
//}

func BenchmarkPrintWord(b *testing.B) {
	for _,c:=range testCases{
		cases:=c
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next(){
				PrintWord(cases.s)
			}
		})
	}

}