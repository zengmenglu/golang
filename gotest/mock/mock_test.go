package mock

import (
	"fmt"
	"testing"
)

func TestMock(t *testing.T){
	text=func(){fmt.Println("mock print text")	} // 对text函数重新赋值
	if text == nil{

	}
	Mock()
}
