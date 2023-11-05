package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
	"time"
)

type wordsHistogram struct {
	words map[string]int
	input []string
}

var nonAlphanumericRegex = regexp.MustCompile(`[^\p{L}\p{N} ]+`)

func clearString(str string) string {
	return nonAlphanumericRegex.ReplaceAllString(str, "")
}

func main() {
	text, err := os.Open("lorem.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer text.Close()

	w := wordsHistogram{
		words: map[string]int{},
		input: []string{},
	}

	scanner := bufio.NewScanner(text)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		str := clearString(scanner.Text())
		w.input = strings.Fields(str)
		for _, word := range w.input {
			_, matched := w.words[word]
			if matched {
				w.words[word] += 1
			} else {
				w.words[word] = 1
			}
		}
	}

	type KeyValue struct {
		Key   string
		Value int
	}

	var sortBigToLow []KeyValue

	for k, v := range w.words {
		sortBigToLow = append(sortBigToLow, KeyValue{k, v})
	}

	sort.Slice(sortBigToLow, func(i, j int) bool {
		return sortBigToLow[i].Value > sortBigToLow[j].Value
	})

	for _, kv := range sortBigToLow {
		time.Sleep(10 * time.Millisecond)
		fmt.Printf("%s-%d\n", kv.Key, kv.Value)
	}

	fmt.Println(len(w.words))
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func (w *wordsHistogram) countWords(text string, turn int) {
	w.words[text] = turn
}
