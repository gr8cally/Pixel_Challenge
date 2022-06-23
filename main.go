package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"time"
)

type BaseImage struct {
	name   string
	pixels []byte
}

func main() {
	start := time.Now()
	category := "Bronze/"
	files, err := ioutil.ReadDir("./" + category)
	if err != nil {
		log.Fatal(err)
	}

	compare("1d25ea94-4562-4e19-848e-b60f1b58deee.raw", files, category)
	elapsed := time.Since(start)
	fmt.Printf("%s took %s", category, elapsed)
}

var f1 = "Bronze/1d25ea94-4562-4e19-848e-b60f1b58deee.raw"
var f2 = "Bronze/6c9952ef-e5bf-4de2-817b-fd0073be8449.raw"
var f3 = "Bronze/1d25ea94-4562-4e19-848e-b60f1b58deee.raw"
var f4 = "Bronze/6c9952ef-e5bf-4de2-817b-fd0073be8449.raw"
var f5 = "Bronze/1d25ea94-4562-4e19-848e-b60f1b58deee.raw"
var f6 = "Bronze/6c9952ef-e5bf-4de2-817b-fd0073be8449.raw"
var f7 = "Bronze/1d25ea94-4562-4e19-848e-b60f1b58deee.raw"
var f8 = "Bronze/6c9952ef-e5bf-4de2-817b-fd0073be8449.raw"
var f9 = "Bronze/1d25ea94-4562-4e19-848e-b60f1b58deee.raw"
var f10 = "Bronze/6c9952ef-e5bf-4de2-817b-fd0073be8449.raw"

func contains(s []float64, e float64) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func compare(image string, imageList []fs.FileInfo, category string) {

	//baseImage, err := os.Open(image)
	//check(err)
	//defer baseImage.Close()
	//
	//currentImage, err := os.Open(f2)
	//check(err)
	//defer currentImage.Close()

	mirror := BaseImage{
		name:   image,
		pixels: nil,
	}
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

		if mirror.pixels == nil {
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
		}
		per := (float64(matchingPixels) / float64(totalPixels)) * 100
		result[f.Name()] = per
		values = append(values, per)
	}
	sort.Float64s(values)
	for k, v := range result {
		if contains(values[len(values)-3:], v) {
			fmt.Println(k, fmt.Sprintf("%.4f%%", v))
		}
	}

	//currentBytes := make([]byte, 3)
	//baseBytes := make([]byte, 3)
	//
	//currentReader := bufio.NewReader(currentImage)
	//baseReader := bufio.NewReader(baseImage)
	//
	//matchingPixels := 0
	//totalPixels := 0
	//
	//for {
	//	//_, errCurr := currentReader.Read(currentBytes)
	//	_, errCurr := io.ReadFull(currentReader, currentBytes)
	//
	//	//bytesRead, err := baseReader.Read(baseBytes)
	//	bytesRead, err := io.ReadFull(baseReader, baseBytes)
	//
	//	if bytesRead == 3 {
	//		totalPixels++
	//	}
	//
	//	if err != nil || errCurr != nil {
	//
	//		if err != io.EOF || errCurr != io.EOF {
	//
	//			log.Fatal(err)
	//		}
	//
	//		break
	//	}
	//	if bytes.Equal(baseBytes, currentBytes) {
	//		matchingPixels++
	//	}
	//
	//}
	//fmt.Printf("match: %v, countedPixels: %v \n", matchingPixels, totalPixels)
	//per := (float64(matchingPixels) / float64(totalPixels)) * 100
	//fmt.Println(per)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
