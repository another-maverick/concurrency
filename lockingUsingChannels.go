package main

import (
	"fmt"
	"time"
)

// you can have sensitive portions of the code locked by using a "Buffered Channel" of 1
//Sender can send a message as long as buffer has space in it. Next senders will be blocked until the buffer has space again
// In normal unbuffered channels, sender sends a message and is blocked until it is read
func main(){
	lockUsingChannel := make(chan bool, 1)
	for i:=0; i <= 10; i++{
		go performJob(i, lockUsingChannel)
	}
	// Make sure all go routines complete. You can also do runtime.gosched or implement wait groups
	time.Sleep(time.Second * 20)
}

func performJob(jobNum int, lockChan chan bool){
	fmt.Printf("Job number  %v wants to take control \n", jobNum)
	 lockChan <- true
	 fmt.Printf("Job number %v has taken control \n", jobNum)
	 time.Sleep(time.Second * 1)
	 fmt.Printf("job number %v is releasing conreol \n", jobNum)
	 <-lockChan
}
