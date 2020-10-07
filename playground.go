package main

import (
	"fmt"
	"sync"
	"time"
)

//wg = "wait group"
var wg sync.WaitGroup

func cleanup() {
	defer wg.Done()
	if r := recover(); r != nil {
		if r == "an example error that stops the script." {
			fmt.Println("you found the example error: ", r)
			say("lol")
		} else {
			fmt.Println("recovered in cleanup: ", r)
		}
	}
}

func say(s string) {
	defer cleanup()
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond * 100)
		if i == 2 {
			panic("this is an example error")
		}
	}
}

func main() {
	defer fmt.Println("this is a defer")
	wg.Add(1)
	go say("hello!")
	wg.Add(1)
	go say("neato")
	wg.Wait()
}
