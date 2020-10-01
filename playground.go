package main

import (
	"fmt"
	"math/rand"
	"time"
)

func multipleReturns(x int) (int, int) {
	doubled := x * 2
	return x, doubled
}

func playWithScope(words *string) {
	*words = "played with scope"
}

func main() {
	rand.Seed(time.Now().UnixNano())
	var value, doubled int
	words := "playing with scope"
	random := rand.Intn(10000)
	//deferred statements run when the function scope has ended
	defer fmt.Println("\nthis is deferred!")

	value, doubled = multipleReturns(random)

	playWithScope(&words)

	defer fmt.Print(value, " doubled is ", doubled, "\n", words)

}
