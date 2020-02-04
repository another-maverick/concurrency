package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func main(){
	go echoMe(os.Stdin, os.Stdout)
	time.Sleep(time.Second * 20)
	fmt.Println("Done Printing")
	os.Exit(0)

}

func echoMe(inp io.Reader, out io.Writer){
	io.Copy(out, inp)
}
