package imageProcessor

import (
	"bytes"
	"io/fs"
	"io/ioutil"
	"log"
	"math"
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

	baseImageBytes := parseFileToBytes(category, image)

	for _, f := range imageList {
		if f.Name() == image {
			continue
		}
		currentImage := parseFileToBytes(category, f.Name())

		matchingPixels := 0
		totalPixels := len(currentImage) / 3

		start, end := 0, 3

		for end <= len(currentImage) {
			if bytes.Equal(baseImageBytes[start:end], currentImage[start:end]) {
				matchingPixels++
			}
			start += 3
			end += 3

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

func parseFileToBytes(category, fileName string) []byte {
	absolutePath, _ := filepath.Abs(category)

	content, err := ioutil.ReadFile(filepath.Join(absolutePath, fileName))
	if err != nil {
		log.Fatal(err)
	}
	return content
}
