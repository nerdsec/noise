package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log/slog"
	"math"
	"os"
	"path/filepath"
)

func main() {
	fmt.Println()
	if len(os.Args) == 1 {
		fmt.Println("missing input filename")
		os.Exit(2)
	}

	filename := filepath.Clean(os.Args[1])

	slog.Info("Reading input file", "filename", filename)

	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	size := len(data)
	sq := math.Sqrt(float64(size))
	width := int(sq)
	height := int(sq)

	colors := make([]color.Gray, len(data))

	for i, v := range data {
		colors[i] = color.Gray{uint8(v)}
	}

	img := image.NewGray(image.Rect(0, 0, width, height))

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			k := y*width + x
			img.SetGray(x, y, colors[k])
		}
	}

	outputFilename := "entropy." + filename + ".png"

	slog.Info("writing output file", "filename", outputFilename)

	f, err := os.Create(outputFilename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = png.Encode(f, img)
	if err != nil {
		panic(err)
	}
}
