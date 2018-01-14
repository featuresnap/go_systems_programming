package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/featuresnap/wordbank"
)

var wordCountFlag, sentenceLengthFlag int
var wordbankPath string

func init() {
	flag.IntVar(&wordCountFlag, "words", 100, "number of words to generate")
	flag.IntVar(&sentenceLengthFlag, "sentence-length", 6, "the length sentences should be")
	flag.StringVar(&wordbankPath, "wordbank", "words.txt", "path to the text file containing the wordbank")
}

func main() {
	flag.Parse()
	rand.Seed(time.Now().UnixNano())
	http.HandleFunc("ipsum", ipsumHandler)
	log.Fatal(http.ListenAndServe(":12345", nil))
}

func ipsumHandler(w http.ResponseWriter, req *http.Request) {
	ipsum, err := generateIpsum(wordCountFlag, sentenceLengthFlag)
	if err != nil {
		fmt.Fprint(w, "Error")
		return
	}
	fmt.Fprint(w, ipsum)
}

func generateIpsum(wordCount, sentenceLength int) (string, error) {
	f, err := os.Open(wordbankPath)
	if err != nil {
		return "", err
	}
	defer f.Close()
	wb, err := wordbank.New(f)
	if err != nil {
		return "", fmt.Errorf("couldn't create the wordbank: %s", err)
	}
	ipsum := ""
	c := make(chan string)
	sentenceCount := 0
	for wordsLeft := wordCount; wordsLeft > 0; wordsLeft -= sentenceLength {
		sentenceCount++
		numWords := sentenceLength
		if wordsLeft < numWords {
			numWords = wordsLeft
		}

		go func() {
			c <- generateSentence(wb, numWords)
		}()

		ipsum := ""
		for i := 0; i < sentenceCount; i++ {
			if ipsum != "" {
				ipsum += " "
			}

			ipsum += <-c
		}
	}

	return ipsum, nil
}

func generateSentence(wb *wordbank.WordBank, wordCount int) string {
	ipsum := ""
	for i := 0; i < wordCount; i++ {
		if ipsum != "" {
			ipsum += " "
		}
		ipsum += wb.GetWord()
	}
	return ipsum + "."
}
