package main

import (
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			log.Print(err)
		}

		param := Param{}
		cycles, err := strconv.ParseFloat(r.Form.Get("cycles"), 64)
		if err != nil {
			param.Cycles = 5.0
		} else {
			param.Cycles = cycles
		}
		res, err := strconv.ParseFloat(r.Form.Get("res"), 64)
		if err != nil {
			param.Res = 0.001
		} else {
			param.Res = res
		}
		size, err := strconv.Atoi(r.Form.Get("size"))
		if err != nil {
			param.Size = 100
		} else {
			param.Size = size
		}
		nframes, err := strconv.Atoi(r.Form.Get("nframes"))
		if err != nil {
			param.Nframes = 64
		} else {
			param.Nframes = nframes
		}
		delay, err := strconv.Atoi(r.Form.Get("delay"))
		if err != nil {
			param.Delay = 8
		} else {
			param.Delay = delay
		}

		lissajous(w, param)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
