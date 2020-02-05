package main

import (
	"fmt"
	"time"
)

func main(){
	//data is a bi directional channel that contains the message
	data := make(chan string)
	//endTime is a channel that gets message after 60 secs
	endTime := time.After(60 * time.Second)
	//done is channel that is used by receiver to signal sender that I am done
	done := make(chan bool)

	go sendMsg(data, done)
	for {
		select {
		case <- data:
			fmt.Println("Received a message")
		case <-endTime:
			done  <- true
			time.Sleep(time.Millisecond * 500)
			return
		}
	}
}

func sendMsg(msg chan<- string, status <-chan bool){
	for{
		select {
		case <-status:
			fmt.Println("Looks like timer has timed out...closing the message channel")
			close(msg)
			return
		default:
			msg<- "Random Message"
			time.Sleep(time.Second * 1)

		}
	}
}
