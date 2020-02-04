package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Starting the main function. Normal course of action...")
	fmt.Println("Noe calling a go routine")
	go func(){
		fmt.Println("Now inside a go routine. We are inside a closure now")
	}()
	fmt.Println("Go routine now finished. we are again in normal course of execution")
	// Lets make sure we make time for scheduling the go routine. If we dont have this line and you have one CPU, the go routine may never run.
	runtime.Gosched()

}
