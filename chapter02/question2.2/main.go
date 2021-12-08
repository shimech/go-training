package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/shimech/go-training/chapter02/question2.2/lengthconv"
	"github.com/shimech/go-training/chapter02/question2.2/tempconv"
	"github.com/shimech/go-training/chapter02/question2.2/weightconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	tArg := flag.String("temp", "", "Temperature")
	lArg := flag.String("length", "", "Length")
	wArg := flag.String("weight", "", "Weight")
	flag.Parse()

	tStr := *tArg
	if tStr == "" {
		fmt.Print("Please input temperature >> ")
		sc.Scan()
		tStr = sc.Text()
	}
	t, err := strconv.ParseFloat(tStr, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	fa := tempconv.Fahrenheit(t)
	ce := tempconv.Celsius(t)
	fmt.Printf("%s = %s, %s = %s\n", fa, tempconv.FToC(fa), ce, tempconv.CToF(ce))

	lStr := *lArg
	if lStr == "" {
		fmt.Print("Please input length >> ")
		sc.Scan()
		lStr = sc.Text()
	}
	l, err := strconv.ParseFloat(lStr, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	fe := lengthconv.Feet(l)
	me := lengthconv.Meter(l)
	fmt.Printf("%s = %s, %s = %s\n", fe, lengthconv.FToM(fe), me, lengthconv.MToF(me))

	wStr := *wArg
	if wStr == "" {
		fmt.Print("Please input weight >> ")
		sc.Scan()
		wStr = sc.Text()
	}
	w, err := strconv.ParseFloat(wStr, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	gr := weightconv.Gram(w)
	po := weightconv.Pound(w)
	fmt.Printf("%s = %s, %s = %s\n", gr, weightconv.GToP(gr), po, weightconv.PToG(po))

}
