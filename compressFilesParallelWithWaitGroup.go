package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"sync"
)

func main(){
	fmt.Println("looking for files to  be compressed ....")
	var wg sync.WaitGroup
	var eachFile string
	for _, eachFile = range os.Args[1:]{
		wg.Add(1)
		go func(filename string) {
			fmt.Printf("Compressing file -- %v \n", filename)
			compressMe(filename)
			wg.Done()
		}(eachFile)
		wg.Wait()


	}
}

func compressMe(flatfile string) error {
	inpptr, err := os.Open(flatfile)

	if err != nil {
		fmt.Println("cannot open the file")
	}
	//Making sure the file pointer gets closed in the end
	defer inpptr.Close()

	outptr, err := os.Create(flatfile + ".gz")
	if err != nil {
		fmt.Println("cannot open the file")
	}
	//Making sure the file pointer gets closed in the end
	defer outptr.Close()

	gzipPtr := gzip.NewWriter(outptr)
	_, err = io.Copy(gzipPtr, inpptr)
	gzipPtr.Close()

	return err


}
// Usage: go run <executable> /tmp/compress/*