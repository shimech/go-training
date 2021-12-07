package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/shimech/go-training/chapter02/question2.1/tempconv"
)

func main() {
	temp, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(tempconv.CToK(tempconv.Celsius(temp)).String())
}
