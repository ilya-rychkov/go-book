package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"strconv"
)

import (
	"log"
	"net/http"
	"time"
)


var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	if len(os.Args) > 1 && os.Args[1] == "web" {
		//!+http
		handler := func(w http.ResponseWriter, r *http.Request) {
			cycles := 5
			size := 100
			nframes := 64
			delay := 8
			if v := r.FormValue("cycles"); v != "" {
				if n, err := strconv.Atoi(v); err == nil {
					cycles = n
				}
			}
			if v := r.FormValue("size"); v != "" {
				if n, err := strconv.Atoi(v); err == nil {
					size = n
				}
			}
			if v := r.FormValue("nframes"); v != "" {
				if n, err := strconv.Atoi(v); err == nil {
					nframes = n
				}
			}
		if v := r.FormValue("delay"); v != "" {
			if n, err := strconv.Atoi(v); err == nil {
				delay = n
			}
		}
		
	lissajous(w, cycles, size, nframes, delay)
		}
		http.HandleFunc("/", handler)
		//!-http
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}
	
	lissajous(os.Stdout, 5, 100, 64, 8)
}

func lissajous(out io.Writer, cycles, size, nframes, delay int) {
	const res = 0.001
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5),
				blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

