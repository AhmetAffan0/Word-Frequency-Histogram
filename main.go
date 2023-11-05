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
	input []string
}

func main() {
	text, err := os.Open("lorem.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer text.Close()

	m := manipulate{
		words: map[string]int{},
		input: []string{},
	}

	scanner := bufio.NewScanner(text)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		m.input = strings.Fields(scanner.Text())
		for _, word := range m.input {
			_, matched := m.words[word]
			if matched {
				m.words[word] += 1
			} else {
				m.words[word] = 1
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for index, element := range m.words {
		fmt.Println(index, "=", element)
	}
	fmt.Println(m.words)

}

func (m *manipulate) countWords(text string, turn int) {
	m.words[text] = turn
}
