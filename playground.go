package main

import (
	"fmt"
	"sync"
	"time"
)

//wg = "wait group"
var wg sync.WaitGroup

//This function only accepts a channel for sending values. It would be a compile-time error to try to receive on this channel.
func recieve(recieveChannel chan<- string, msg string) {
	recieveChannel <- msg
}

//The pong function accepts one channel for receives (recieveChannel) and a second for sends (sendChannel).
func send(recieveChannel <-chan string, sendChannel chan<- string) {
	msg := <-recieveChannel
	sendChannel <- msg
}

func main() {
	recieveChannel := make(chan string, 2)
	sendChannel := make(chan string, 2)
	recieve(recieveChannel, "passed message")
	send(recieveChannel, sendChannel)
	//print what is stored in the sendChannel
	fmt.Println(<-sendChannel)

	channel1, channel2 := make(chan string, 1), make(chan string, 1)

	go func() {
		time.Sleep(1 * time.Second)
		channel1 <- "first message"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		channel2 <- "second message"
	}()

	for i := 0; i < 2; i++ {
		select {
		/*select operates whenevr a case is met so in this example once a case is satisfied "channel1 has a message" then it will print the message. if we change this to have a default then it will rush through the select and print the example twice*/
		case msg1 := <-channel1:
			fmt.Println("received", msg1)
		case msg2 := <-channel2:
			fmt.Println("received", msg2)
			// default:
			// 	fmt.Println("example")
		}
	}
}
