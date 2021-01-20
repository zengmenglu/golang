package mock

import "fmt"

var text = func() {
	fmt.Println("print text")
}

func Mock() {
	text()
}
