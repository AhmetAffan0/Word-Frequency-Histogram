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

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func makeUI(textBox *widget.Entry) *fyne.Container {
	return container.New(layout.NewGridLayout(1),
		textBox,
	)
}

type wordsHistogram struct {
	words map[string]int
	input []string
}

var nonAlphanumericRegex = regexp.MustCompile(`[^\p{L}\p{N} ]+`)

func clearString(str string) string {
	return nonAlphanumericRegex.ReplaceAllString(str, "")
}

func main() {
	a := app.New()
	w := a.NewWindow("Word Frequency Histogram")

	text, err := os.Open("text.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer text.Close()

	wh := wordsHistogram{
		words: map[string]int{},
		input: []string{},
	}

	scanner := bufio.NewScanner(text)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		str := clearString(scanner.Text())
		wh.input = strings.Fields(str)
		for _, word := range wh.input {
			_, matched := wh.words[word]
			if matched {
				wh.words[word] += 1
			} else {
				wh.words[word] = 1
			}
		}
	}

	type KeyValue struct {
		Key   string
		Value int
	}

	var sortBigToLow []KeyValue

	for k, v := range wh.words {
		sortBigToLow = append(sortBigToLow, KeyValue{k, v})
	}

	sort.Slice(sortBigToLow, func(i, j int) bool {
		return sortBigToLow[i].Value > sortBigToLow[j].Value
	})

	for _, kv := range sortBigToLow {
		time.Sleep(10 * time.Millisecond)
		fmt.Printf("%s-%d\n", kv.Key, kv.Value)
	}

	fmt.Println(len(wh.words))
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	w.ShowAndRun()
}

func (w *wordsHistogram) countWords(text string, turn int) {
	w.words[text] = turn
}
