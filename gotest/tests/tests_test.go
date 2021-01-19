package tests

import "testing"

func TestHelloWorld(t *testing.T) {
	t.Log("TestHelloWorld")
	HelloWorld()
}

// Run() - subtests
func TestPrintWords(t *testing.T) {
	testCases := []struct {
		name string
		s    string
	}{
		{"single", "a"},
		{"double", "aa"},
		{"trible", "aaa"},
	}
	for _, cases := range testCases {
		t.Run(cases.name, func(t *testing.T) {
			if PrintWords(cases.s) == "a" {
				t.Fatal("output is \"a\"")
			}
		})
	}
}


// without t.Run() cannot form sub-tests
func TestPrintWords2(t *testing.T) {
	testCases := []struct {
		name string
		s    string
	}{
		{"single", "a"},
		{"double", "aa"},
		{"trible", "aaa"},
	}
	for _, cases := range testCases {
		if PrintWords(cases.s) == "a" {
			t.Fatal("output is \"a\"")
		}
	}
}
