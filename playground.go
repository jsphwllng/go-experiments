package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond * 100)
	}
}

func main() {
	go say("hello!")
	go say("neato")

	/*
		If these are both go routines then nothing happens! Weird! Because the program finishes before the go routine does. If we shove a wait at the end then because it's waiting the whole time the rest of the concurrency happens...
	*/
	time.Sleep(time.Second / 10)
}
