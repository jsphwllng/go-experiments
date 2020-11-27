package main

import (
	"fmt"
	"testing"
)

func TestString(t *testing.T) {
	var tests = []struct {
		in   string
		want string
	}{
		{"hello", "eo"},
		{"joseph", "oe"},
		{"hello, world!", "eoo"},
		{"lorem ipsum haha", "oeiuaa"},
		{"UPCASETEST", "uaee"},
		{"n0-v0w3l5", ""},
	}
	for _, testingTable := range tests {
		testname := fmt.Sprintf("%s", testingTable.in)
		// t.Run enables subtests
		t.Run(testname, func(t *testing.T) {
			ans := OnlyVowels(testingTable.in)
			if ans != testingTable.want {
				t.Errorf("got %s, want %s", ans, testingTable.want)
			}
		})
	}
}
