package main

import (
	"fmt"
	"math/rand"
	"time"
)

//receiver is blocked until the sender sends something to a channel the receiver is listening
//sender is blocked until receiver clears the channel
func sender(clock chan int){
	for i:=0; ; i++ {
		fmt.Println("Sending", i)
		//send the i to clock
		clock <- i
		amt := time.Duration(rand.Intn(1000))
		time.Sleep(time.Millisecond * amt)
	}
}

func receiver (clock chan int){
	for {
	//	receive integer from channel and priunt int
	//	you can only write again if someone has read (a buffer with only space of 1)
		msg := <- clock
		fmt.Println("Receiving", msg)

		amt := time.Duration(rand.Intn(1000))
		time.Sleep(time.Millisecond * amt)
	}
}


func another_sender(clock chan int){
	for i:=0; ; i-- {
		fmt.Println("Sending", i)
		//send the i to clock
		clock <- i

	}
}

func main() {
	var clock chan int = make(chan int)

	go sender(clock)
	go another_sender(clock)
	go receiver(clock)

	var input string
	fmt.Scanln(&input)
}
