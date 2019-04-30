package main

import (
	"bytes"
	"compress/flate"
	"fmt"
	"github.com/klauspost/compress/fse"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	_ "image/png"
	"io/ioutil"
	"log"
	"math"
	"math/rand"
	"os"
	"sort"
	"strings"
)

func main() {
	inBytes, err := ioutil.ReadFile("./fse/rp-orig.png")
	exitOnErr(err)

	inImg, _, err := image.Decode(bytes.NewBuffer(inBytes))
	exitOnErr(err)
	img := toGray(inImg)
	writeout(img, "./fse/rp-orig-hist.png")
	writeout(scramble(img), "./fse/rp-scrambled.png")
	writeout(sorted(img), "./fse/rp-sorted.png")
	writeout(posterize(img, 4), "./fse/rp-post-4.png")
	writeout(posterize(img, 50), "./fse/rp-post-50.png")
	writeout(downLeft(img), "./fse/rp-downleft.png")
	writeout(downLeft(posterize(img, 32)), "./fse/rp-post-32-downleft.png")
}

// writeout will write out
func writeout(img *image.Gray, name string) {
	combined, symbols, limit := combinedWithHist(img)
	out, err := os.Create(name)
	exitOnErr(err)
	defer out.Close()
	err = png.Encode(out, combined)
	exitOnErr(err)
	out, err = os.Create(strings.Replace(name, ".png", ".txt", 1))
	exitOnErr(err)
	defer out.Close()
	compFSE, err := fse.Compress(img.Pix, nil)
	compDef := bytes.Buffer{}
	fw, err := flate.NewWriter(&compDef, flate.HuffmanOnly)
	exitOnErr(err)
	_, err = fw.Write(img.Pix)
	exitOnErr(err)
	fw.Close()
	bits := int(0.9999 + math.Log2(float64(symbols)))
	fmt.Fprintf(out, "// In: %d, Symbols: %d, Lim: %d, FSE: %d (%.2f:1), Huff: %d (%.2f:1), %d bits: %d\n",
		len(img.Pix), symbols, limit,
		len(compFSE), float64(len(img.Pix))/float64(len(compFSE)),
		compDef.Len(), float64(len(img.Pix))/float64(compDef.Len()),
		bits, int((bits*len(img.Pix)+7)/8))

}

// combinedWithHist returns a grayscale image (256x256) with histogram
func combinedWithHist(in *image.Gray) (out *image.Gray, symbols, limit int) {
	hist, symbols, limit := histogram(in)

	cRect := image.Rect(0, 0, 512+20, 256)
	out = image.NewGray(cRect)
	draw.Draw(out, cRect, image.NewUniform(color.Gray{255}), image.Pt(0, 0), draw.Over)
	draw.Draw(out, cRect, in, image.Pt(0, 0), draw.Over)
	cRect.Min.X = 256 + 20
	draw.Draw(out, cRect, hist, image.Pt(0, 0), draw.Over)
	return out, symbols, limit
}

func histogram(in *image.Gray) (out *image.Gray, symbols, limit int) {
	var hist [256]int
	w := in.Rect.Max.X
	h := in.Rect.Max.Y
	for y := 0; y < h; y++ {
		line := in.Pix[y*in.Stride : y*in.Stride+w]
		for _, v := range line {
			hist[v]++
		}
	}

	// Maximum for normalization.
	shannon := float64(0)
	total := float64(w * h)
	max := 0
	for i := range hist[:] {
		if hist[i] > max {
			max = hist[i]
		}
		n := float64(hist[i])
		if n > 0 {
			shannon += math.Log2(total/n) * n
		}
	}
	out = image.NewGray(image.Rect(0, 0, 256, 256))
	invMax := 1.0 / float64(max)
	for x := range hist[:] {
		if hist[x] > 0 {
			symbols++
		}
		height := 256 * float64(hist[x]) * invMax
		height = 256 - height
		for y := 0; y < 256; y++ {
			weight := (height - float64(y)) * 255
			weight = math.Min(math.Max(weight, 0), 255.5)
			out.SetGray(x, y, color.Gray{uint8(weight)})
		}
	}
	return out, symbols, int(shannon+7) / 8
}

// posterize reduces the number of colors by posterizing.
func posterize(img *image.Gray, n uint8) *image.Gray {
	inRect := img.Bounds()
	grey := image.NewGray(image.Rect(0, 0, inRect.Dx(), inRect.Dy()))
	for y := 0; y < grey.Rect.Dy(); y++ {
		for x := 0; x < grey.Rect.Dx(); x++ {
			pix := img.GrayAt(x+inRect.Min.X, y+inRect.Min.Y)
			pix.Y -= pix.Y % n
			grey.Set(x, y, pix)
		}
	}
	return grey
}

// scramble moves around pixels. Image must be 256x256.
func scramble(img *image.Gray) *image.Gray {
	inRect := img.Bounds()
	rng := rand.New(rand.NewSource(1337))
	grey := image.NewGray(image.Rect(0, 0, inRect.Dx(), inRect.Dy()))
	draw.Draw(grey, inRect, img, image.Pt(0, 0), draw.Over)
	for y := 0; y < grey.Rect.Dy(); y++ {
		for x := 0; x < grey.Rect.Dx(); x++ {
			rnd := int(rng.Uint32())
			x2, y2 := rnd&255, (rnd>>8)&255
			p1 := grey.GrayAt(x, y)
			p2 := grey.GrayAt(x2, y2)
			// Swap p1 & p2
			grey.Set(x, y, p2)
			grey.Set(x2, y2, p1)
		}
	}
	return grey
}

// sorted orders pixels by frequency in histogram.
func sorted(img *image.Gray) *image.Gray {
	inRect := img.Bounds()
	type histEntry struct {
		n   int
		org uint8
	}
	var hist [256]histEntry
	w := img.Rect.Max.X
	h := img.Rect.Max.Y
	for y := 0; y < h; y++ {
		line := img.Pix[y*img.Stride : y*img.Stride+w]
		for _, v := range line {
			hist[v].n++
		}
	}

	// Record original position before sorting
	for i := range hist[:] {
		hist[i].org = uint8(i)
	}
	sort.Slice(hist[:], func(i, j int) bool {
		a, b := hist[i], hist[j]
		if a.n != b.n {
			return a.n > b.n
		}
		// Equal count, just use pixel value
		return a.org > b.org
	})

	// Create input -> output mapping.
	var mapping [256]uint8
	for i := range hist[:] {
		mapping[hist[i].org] = uint8(i)
	}

	grey := image.NewGray(image.Rect(0, 0, inRect.Dx(), inRect.Dy()))
	for y := 0; y < grey.Rect.Dy(); y++ {
		for x := 0; x < grey.Rect.Dx(); x++ {
			pix := img.GrayAt(x, y)
			pix.Y = mapping[pix.Y]
			grey.Set(x, y, pix)
		}
	}
	return grey
}

// downLeft does down+left prediction on image
func downLeft(img *image.Gray) *image.Gray {
	inRect := img.Bounds()
	grey := image.NewGray(image.Rect(0, 0, inRect.Dx(), inRect.Dy()))
	last := color.Gray{Y: 0}
	for y := 0; y < grey.Rect.Dy(); y++ {
		if y > 0 {
			last = img.GrayAt(0, y-1)
		}
		for x := 0; x < grey.Rect.Dx(); x++ {
			if x > 0 {
				last = img.GrayAt(x-1, y)
			}
			pix := img.GrayAt(x, y)
			pix.Y = pix.Y - last.Y
			grey.Set(x, y, pix)
		}
	}
	return grey
}

// toGray converts an image to greyscale.
// Side effect: move image to 0,0.
func toGray(img image.Image) *image.Gray {
	inRect := img.Bounds()
	grey := image.NewGray(image.Rect(0, 0, inRect.Dx(), inRect.Dy()))
	for y := 0; y < grey.Rect.Dy(); y++ {
		for x := 0; x < grey.Rect.Dx(); x++ {
			grey.Set(x, y, color.GrayModel.Convert(img.At(x+inRect.Min.X, y+inRect.Min.Y)))
		}
	}
	return grey
}

func exitOnErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
