package main

import (
	"bufio"
	"fmt"
	"image/color"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func makeUI(text *canvas.Text, textBox *widget.Entry) *fyne.Container {
	return container.New(layout.NewGridLayout(0),
		text,
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
	w := a.NewWindow("Word Frequency Counter")
	w.Resize(fyne.NewSize(640, 480))

	input := widget.NewEntry()
	input.MultiLine = true
	input.TextStyle.Bold = true
	input.TextStyle.Italic = true

	label := canvas.NewText("Word Frequency Counter", color.NRGBA{175, 175, 175, 255})
	label.TextSize = 30
	label.Alignment = fyne.TextAlignCenter
	label.TextStyle = fyne.TextStyle{Italic: false, Bold: true, Monospace: false}

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

	progress := input.Text

	for _, kv := range sortBigToLow {
		time.Sleep(10 * time.Millisecond)
		progress += fmt.Sprintf("%s-%d\n", kv.Key, kv.Value)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	input.SetText(progress)

	input.Disable()

	w.SetContent(makeUI(label, input))

	w.ShowAndRun()
}

func (w *wordsHistogram) countWords(text string, turn int) {
	w.words[text] = turn
}
