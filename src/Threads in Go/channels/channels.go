package main

import (
	"fmt"
)

/********************************************************************************************************
Unbuffered Channels
Channels are used for communication between two goroutines.
They are typed so you need to specifiy the type while making the goroutines.
The type specifies what type of data will be transfered between 2 goroutines.
Channels are by default unbuffered if one task sends data to a channel then it'll wait for the other task
to receive the data.
So by default they are synchronised in nature.
So because of this property we can use channels just for synchronisation and not for passing values.
********************************************************************************************************/

/********************************************************************************************************
Buffered Channels
Channels by default don't store any data so buffer has 0 capacity or are unbuffered in nature.
To make buffered channels pass another argument to the make function which defines the size of the buffer
In buffered channel synchronisation(blocking) works differently.
The channel is not blocked until the buffer capacity is full. Like if buffer size is 3 then channel can
do three sends and not get blocked if no receives are done.
Same vice versa, receiving only blocks when buffer is empty.
********************************************************************************************************/

func prod(a, b int, c chan int) {
	c <- a * b // This sends data to the channel c
}

func main() {
	c := make(chan int, 1) // make function is used to make channels
	// Arrow syntax is used to send and recieve data from channels
	go prod(1, 2, c)
	go prod(3, 4, c)
	x := <-c // Recieve data from channel c
	y := <-c
	fmt.Println(x * y)
}
