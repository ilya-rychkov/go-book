package main // вывести название файлов, в которых слово встречается больше 1 раза

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"unicode"
)

func main() {
	counts := make(map[string]int)
	foundIn := make(map[string][]string)
	files := os.Args[1:]
	//fmt.Println(files)
	
	for _, arg := range files {
		f, err := os.Open(arg)
		if err != nil {
			fmt.Printf("%s\n", err)
			continue
		}
		countWords(f, counts, foundIn)
		f.Close()
	}

	for word, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t%s\n", n, strings.Join(foundIn[word], " "), word)
		}
	}
}

	func in(needle string, strings []string) bool {
		for _, s := range strings {
			if s == needle {
				return true
			}
		}
		return false
	}

func countWords(f *os.File, counts map[string]int, foundIn map[string][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		words := strings.FieldsFunc(line, func(r rune) bool {
    		return !unicode.IsLetter(r)
		})
		filename := filepath.Base(f.Name())
		for _, word := range words {
			word = strings.ToLower(word)
			counts[word]++
			if !in(filename, foundIn[word]) {
				foundIn[word] = append(foundIn[word], filename)
		}
		}
	}
	//fmt.Println(counts)
} 
			