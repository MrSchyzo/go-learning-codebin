package stringutils

import "testing"

func TestReverseString(t *testing.T) {

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

	for _, c := range cases {

		then := ReverseString(c.given)

		if then != c.then {
			t.Errorf("ReverseString(%q) returned %q, expected %q", c.given, then, c.then)
		}

	}

}
