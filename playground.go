package main

import (
	"fmt"
)

func main() {
	a := make([]string, 2)
	a[0], a[1] = "hmm", "ahh"
	s := append(a, "two", "wow", "blam")
	for i := range s {
		fmt.Println(s[i])
	}
}
