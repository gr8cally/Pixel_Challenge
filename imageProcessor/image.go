package imageProcessor

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
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

func Compare(image string, imageList []fs.FileInfo, category string) {
	result := make(map[string]float64)
	values := make([]float64, 0)

	for _, f := range imageList {
		if f.Name() == image {
			continue
		}

		currentImage, err := os.Open(category + f.Name())
		check(err)
		defer currentImage.Close()

		currentBytes := make([]byte, 3)
		currentReader := bufio.NewReader(currentImage)

		matchingPixels := 0
		totalPixels := 0

		baseImage, err := os.Open(category + image)
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
	fmt.Printf("The closest images to %s in %s category are:\n", image, category)
	for k, v := range result {
		if standards.Contains(values[len(values)-3:], v) {
			fmt.Println(k, fmt.Sprintf("%.4f%%", v))
		}
	}
}
