package main

import (
	"fmt"
	"os"
	"time"
)

//Echo using channels. What we put in std input gets pushed to stdout in a channel that is shared
func main(){
	endTime := time.After(time.Second * 30)
	biDirChan := make(chan []byte)

	go readStdIn(biDirChan)
		for{
			select {
			case buf := <-biDirChan:
				os.Stdout.Write(buf)
			case <-endTime:
				fmt.Println("Timer has ended")
				os.Exit(0)
			}
		}
}

func readStdIn(out chan<- []byte) {
	for{
		inpData := make([]byte, 1024)
		dataLength, _ := os.Stdin.Read(inpData)
		if dataLength > 0{
			out  <- inpData
		}
	}
}