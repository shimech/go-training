package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/shimech/go-training/chapter02/question2.2/lengthconv"
	"github.com/shimech/go-training/chapter02/question2.2/tempconv"
	"github.com/shimech/go-training/chapter02/question2.2/weightconv"
)

func main() {
	tArg := flag.String("temp", "", "Temperature")
	lArg := flag.String("length", "", "Length")
	wArg := flag.String("weight", "", "Weight")
	flag.Parse()

	if *tArg != "" {
		t, err := strconv.ParseFloat(*tArg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		fmt.Printf("%s = %s, %s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))
	}

	if *lArg != "" {
		l, err := strconv.ParseFloat(*lArg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
		f := lengthconv.Feet(l)
		m := lengthconv.Meter(l)
		fmt.Printf("%s = %s, %s = %s\n", f, lengthconv.FToM(f), m, lengthconv.MToF(m))
	}

	if *wArg != "" {
		w, err := strconv.ParseFloat(*wArg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
		g := weightconv.Gram(w)
		p := weightconv.Pound(w)
		fmt.Printf("%s = %s, %s = %s\n", g, weightconv.GToP(g), p, weightconv.PToG(p))
	}
}
