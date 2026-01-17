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

const missingInputExitCode = 2

func main() {
	fmt.Println()
	if len(os.Args) == 1 {
		fmt.Println("missing input filename")
		os.Exit(missingInputExitCode)
	}

	filename := filepath.Clean(os.Args[1])

	slog.Info("Reading input file", "filename", filename)

	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	slog.Info("shannon entropy", "score", calculateEntropy(data))

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

	outputFilename = filepath.Clean(outputFilename)
	f, err := os.Create(outputFilename)
	if err != nil {
		panic(err)
	}
	defer func() {
		err := f.Close()
		if err != nil {
			slog.Error("failed to close file", "error", err)
		}
	}()

	err = png.Encode(f, img)
	if err != nil {
		panic(err)
	}
}

func calculateEntropy(data []byte) float64 {
	if len(data) == 0 {
		return 0.0
	}

	counts := make(map[byte]int)
	for _, b := range data {
		counts[b]++
	}

	var entropy float64
	totalBytes := float64(len(data))

	for _, count := range counts {
		probability := float64(count) / totalBytes
		// The formula is: H = -sum(p_i * log2(p_i))
		// We use math.Log2 for the logarithm to base-2 to get the result in bits.
		entropy -= probability * math.Log2(probability)
	}

	return entropy
}
