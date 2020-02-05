package main
//This program can count the number of words in the files that are provided as arguments. This uses  multiple go routines
// that access, use and modify a common struct. To prevent race conditions, we need to lock and unlock the resource using mutex
import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	newWordData := newWords()
	for _, eachFile := range os.Args[1:] {
		wg.Add(1)
		go func(filename string) {
			if err := checkWords(filename, newWordData); err!=  nil{
				fmt.Printf("cannot count words for %v \n", filename)
			}
			wg.Done()
		}(eachFile)
	}
wg.Wait()
fmt.Println("Here is the summary on word count....")
	newWordData.Lock()
for eachWord, eachWordCount := range newWordData.wordCount {
	fmt.Printf("%v ---------- %v \n", eachWord, eachWordCount)
}
newWordData.Unlock()
}

type wordData struct {
	//mutex to make sure only one counter function can modify the struct
	sync.Mutex
	wordCount map[string]int
}

func newWords() *wordData {
	return &wordData{wordCount: map[string]int{}}
}

func (w *wordData)  counter(eachWord string, count int){
	//to prevent race condition
	w.Lock()
	defer w.Unlock()
	currentCount,  ok := w.wordCount[eachWord]
	if !ok {
		w.wordCount[eachWord] = count
		return
	}
	w.wordCount[eachWord] = currentCount + count
}

func checkWords(filename string, words *wordData) error {
	f , err := os.Open(filename)
	if err != nil {
		fmt.Printf("cannot open file -  %v for analysis \n", filename)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan(){
		thisWord := strings.ToLower(scanner.Text())
		words.counter(thisWord, 1)
	}
	return scanner.Err()
}


