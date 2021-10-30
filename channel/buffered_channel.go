// In Go, there are mainly two types of channel.
// 1. Unbuffered channel 
// 2. Buffered channel

// Buffered channel can work asyncronously
// Let's see an example of buffered channel

package main

import (
	"fmt"
	"sync"
)

var msgs []string = []string{"msg1", "msg2", "msg3", "msg4", "msg5"}

var wg sync.WaitGroup

func SendMessage(ch chan string,msg string) {
	fmt.Println("Sending msg: ", msg)
	ch <- msg
	wg.Done()
}

func main() {
	len := len(msgs)
	wg.Add(len)
	bufferedChan := make(chan string, len)
	defer close(bufferedChan)
	for _,msg := range msgs {
		go SendMessage(bufferedChan, msg)
	}

	wg.Wait()

	for i:=0;i<len;i++ {
		receivedMsg := <- bufferedChan
		fmt.Println("Received msg: ", receivedMsg)
	}


}





