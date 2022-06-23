package imageProcessor

import (
	"bufio"
	"bytes"
	"io"
	"io/fs"
	"log"
	"math"
	"os"
	"path/filepath"
	"pixel_challenge/standards"
	"sort"
)

type BaseImage struct {
	name   string
	pixels []byte
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Compare(image string, imageList []fs.FileInfo, category string) map[string]float64 {
	result := make(map[string]float64)
	values := make([]float64, 0)
	absolutePath, _ := filepath.Abs(category)
	for _, f := range imageList {
		if f.Name() == image {
			continue
		}

		currentImage, err := os.Open(filepath.Join(absolutePath, f.Name()))
		check(err)
		defer currentImage.Close()

		currentBytes := make([]byte, 3)
		currentReader := bufio.NewReader(currentImage)

		matchingPixels := 0
		totalPixels := 0

		baseImage, err := os.Open(filepath.Join(absolutePath, image))
		check(err)
		defer baseImage.Close()

		baseBytes := make([]byte, 3)
		baseReader := bufio.NewReader(baseImage)

		for {
			//_, errCurr := currentReader.Read(currentBytes)
			_, errCurr := io.ReadFull(currentReader, currentBytes)

			//bytesRead, err := baseReader.Read(baseBytes)
			bytesRead, err := io.ReadFull(baseReader, baseBytes)

			if bytesRead == 3 {
				totalPixels++
			}

			if err != nil || errCurr != nil {

				if err != io.EOF || errCurr != io.EOF {

					log.Fatal(err)
				}

				break
			}
			if bytes.Equal(baseBytes, currentBytes) {
				matchingPixels++
			}

		}

		per := standards.GetPercent(matchingPixels, totalPixels)
		result[f.Name()] = per
		values = append(values, per)
	}
	sort.Float64s(values)
	top3 := make(map[string]float64)
	for k, v := range result {
		if standards.Contains(values[len(values)-3:], v) {
			top3[k] = math.Round(v*10000) / 10000
		}
	}
	return top3
}
