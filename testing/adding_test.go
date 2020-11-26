package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestAdding(t *testing.T) {
	ans := Adding(2, 2)
	if ans != 4 {
		t.Errorf("expected 2+2 to equal 4, got: %d", ans)
	}
}

func TestTableDriven(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())
	random1, random2 := rand.Intn(100000), rand.Intn(100000)
	var tests = []struct {
		a, b int
		want int
	}{
		{0, 1, 1},
		{2, 2, 4},
		{99, 1, 100},
		{-1, 1, 0},
		{random1, random2, (random1 + random2)},
	}
	for _, testTable := range tests {
		testname := fmt.Sprintf("%d,%d", testTable.a, testTable.b)
		// t.Run enables subtests
		t.Run(testname, func(t *testing.T) {
			ans := Adding(testTable.a, testTable.b)
			if ans != testTable.want {
				t.Errorf("got %d, want %d", ans, testTable.want)
			}
		})
	}
}
