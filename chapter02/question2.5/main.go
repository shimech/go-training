package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/shimech/go-training/chapter02/question2.5/popcount"
)

func main() {
	x, err := strconv.ParseUint(os.Args[1], 10, 64)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(popcount.PopCountLogicalAndDecrement(x))
}
