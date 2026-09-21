package main//вывести файлы, в которых буква встречается более 1 раза

import (
	"fmt"
	"os"
	"unicode"
	"bufio"
	"path/filepath"
	"strings"
)

func main() {
	counts := make(map[rune]int)
	foundIn := make(map[rune][]string)
	files := os.Args[1:]
	
	for _, arg := range files {
		f, err := os.Open(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			continue
		}
		countLetters(f, counts, foundIn)
		f.Close()
	}
	for letter, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t%s\n", n, strings.Join(foundIn[letter], " "), string(letter))
		}
	}
}

func in(needle string, letters[]string) bool{
	for _, l := range letters {
		if l == needle {
			return true
		}
	}
	return false
}

func countLetters(f *os.File, counts map[rune]int, foundIn map[rune][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		fileName := filepath.Base(f.Name())
		for _, letter := range line {
			if !unicode.IsLetter(letter) {
				continue
			}
			letter = unicode.ToLower(letter)
			counts[letter]++
			if !in(fileName, foundIn[letter]) {
				foundIn[letter] = append(foundIn[letter], fileName)

			}
		}
		
	}
}