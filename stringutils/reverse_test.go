package stringutils

import "testing"

/*
Tests - 101.
1- Create a file with <name>_test.go
2- Package may be the same of the tested functions, remember to import "testing" package
3- Write a void function named Test<whatToTest> with argument of type *testing.T
4- After tests are written, "cd" on the same directory of *_test.go and execute "go test"
*/
func TestReverseString(t *testing.T) {

	// Creates a struct on the fly for acceptance test
	cases := []struct {
		given string
		then  string
	}{
		{"Otto", "ottO"},
		{"abc", "cba"},
		{"azza", "azza"},
		{"", ""},
		{"öl lë", "ël lö"},
	}

	//Remember, for in go must enumerate key and value, that's why you have and underscore at first
	for _, c := range cases {

		then := ReverseString(c.given)

		//Is there any equivalent of string templating? My guess is "no" seeing the "simplicity" of Go
		if then != c.then {
			t.Errorf("ReverseString(%q) returned %q, expected %q", c.given, then, c.then)
		}

	}

}
