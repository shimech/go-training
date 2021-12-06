package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
)

var red = color.RGBA{0xFF, 0x00, 0x00, 0xFF}
var green = color.RGBA{0x00, 0xFF, 0x00, 0xFF}
var blue = color.RGBA{0x00, 0x00, 0xFF, 0xFF}
var palette = []color.Color{color.Black, red, green, blue}

type Param struct {
	Cycles  float64
	Res     float64
	Size    int
	Nframes int
	Delay   int
}

func lissajous(out io.Writer, param Param) {
	cycles := param.Cycles
	res := param.Res
	size := param.Size
	nframes := param.Nframes
	delay := param.Delay
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5), uint8(rand.Intn(3)+1))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
