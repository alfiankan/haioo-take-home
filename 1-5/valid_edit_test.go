package main

import (
	"fmt"
	"testing"
)

func IsValidEdit(inA, inB string) (valid bool) {

	attempts := map[rune]int{}

	for _, c := range inA {
		if attempts[c] == 0 {
			attempts[c] = 1
			continue
		}
		attempts[c] += 1

	}

	for _, c := range inB {
		if attempts[c] != 0 {
			attempts[c] -= 1
			if attempts[c] == 0 {
				delete(attempts, c)
			}
		} else {
			attempts[c] += 1
		}
	}

	return len(attempts) == 1
}

func TestValidEdit(t *testing.T) {

	/*

		allowed process -> insert, remove, replace

		telkom -> tlkom

		so we need delete 'e' that only one operation

		telkom -> telecom

		so we need to :
			- delete 'k'
			- add 'e'
			- add 'c'

		{

		}

		telkom -> tlkom

		t = 0
		e = 1
		l = 0
		k = 0
		o = 0
		m = 0

		tlkom

		-> delete e because e still 1

		------------

		telkom -> telecom

		t = 0
		e = -1
		l = 0
		k = 1
		o = 0
		m = 0
		c = 1

		telecom

		------

		jooi -> joodi

		j = 0
		o = 0
		i = 0
		d = 1

		joodi


	*/

	testCases := []struct {
		InputA   string
		InputB   string
		Expected bool
	}{
		{"telkom", "telecom", false},
		{"telkom", "tlkom", true},
		{"joi", "jodi", true},
		{"mantra", "matrei", false},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%s == %s", tc.InputA, tc.InputB), func(t *testing.T) {
			res := IsValidEdit(tc.InputA, tc.InputB)

			if res != tc.Expected {
				t.Error("Expected :", tc.Expected, "Actual :", res)
			}
		})
	}

}
