package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Dup struct {
	Frequency int
	Files     []string
}

func main() {
	counts := make(map[string]Dup)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, "Stdin", counts)
	} else {
		for _, file := range files {
			f, err := os.Open(file)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, file, counts)
			f.Close()
		}
	}
	for text, dup := range counts {
		if dup.Frequency > 1 {
			fmt.Printf("%d\t%s\t%s\n", dup.Frequency, text, strings.Join(dup.Files, ", "))
		}
	}
}

func countLines(f *os.File, file string, counts map[string]Dup) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		text := input.Text()
		prevDup := counts[text]
		newFiles := prevDup.Files
		if !contains(newFiles, file) {
			newFiles = append(newFiles, file)
		}
		counts[text] = Dup{
			Frequency: prevDup.Frequency + 1,
			Files:     newFiles,
		}
	}

}

func contains(strs []string, str string) bool {
	for _, strComp := range strs {
		if str == strComp {
			return true
		}
	}
	return false
}
