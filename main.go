package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type manipulate struct {
	words map[string]int
	or    int
	to    int
	of    int
}

func main() {
	text, err := os.Open("lorem.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer text.Close()

	m := manipulate{
		words: map[string]int{},
		or:    0,
		to:    0,
		of:    0,
	}

	scanner := bufio.NewScanner(text)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		or := strings.Count(strings.ToLower(scanner.Text()), "or")
		m.or += or
		to := strings.Count(strings.ToLower(scanner.Text()), "to")
		m.to += to
		of := strings.Count(strings.ToLower(scanner.Text()), "of")
		m.of += of

		m.countWords(scanner.Text(), 0)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(m.or)
	fmt.Println(m.to)
	fmt.Println(m.of)
	fmt.Println(m.words)
}

func (m *manipulate) countWords(text string, turn int) {
	m.words[text] = turn
}
