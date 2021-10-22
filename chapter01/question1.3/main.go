package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(stringifyArgsWithLoop(os.Args))
}

func stringifyArgsWithLoop(arguments []string) string {
	s, separator := "", ""
	for _, argument := range arguments {
		s += separator + argument
		separator = " "
	}
	return s
}

func stringifyArgsWithJoin(arguments []string) string {
	return strings.Join(arguments, " ")
}
