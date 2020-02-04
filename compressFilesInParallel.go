package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
)

func main(){
	fmt.Println("looking for files to  be compressed ....")
	for _, eachFile := range os.Args[1:]{
		fmt.Printf("Compressing file -- %v \n", eachFile)
		compressMe(eachFile)
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