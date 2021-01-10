package normaltest

import "testing"

func TestHelloWorld(t *testing.T) {

	t.Error("TestHelloWorld")
	HelloWorld()
}

func TestHelloWorld2(t *testing.T) {
	t.Logf("TestHelloWorld2")
}