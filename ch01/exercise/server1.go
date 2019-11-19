// Modify the Lissajous server to read parameter values from the URL.
// For exam- ple, you might arrange it so that a URL like http://localhost:8000/?cycles=20
// sets the number of cycles to 20 instead of the default 5.
// Use the strconv.Atoi function to convert the string parameter into an integer.
// You can see its documentation with go doc strconv.Atoi.
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			log.Print(err)
		}

		cycles := r.Form["cycles"]
		var cyclesCount = 5
		if len(cycles) > 0 {
			cyclesCount, err := strconv.Atoi(cycles[0])
			if err != nil {
				cyclesCount = 5
				log.Print(err)
			}
		}
		lissajous(cyclesCount, w)
	})

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

func lissajous(cycles int, out io.Writer) {
	const (
		//cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)

	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}

		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
