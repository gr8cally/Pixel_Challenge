package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"pixel_challenge/imageProcessor"
	"time"
)

func main() {
	start := time.Now()
	category := "Bronze/"
	files, err := ioutil.ReadDir(category)
	if err != nil {
		log.Fatal(err)
	}

	image := "1d25ea94-4562-4e19-848e-b60f1b58deee.raw"
	top3 := imageProcessor.Compare(image, files, category)
	fmt.Printf("The closest images to %s in %s category are:\n", image, category)
	for k, v := range top3 {
		fmt.Println(k, fmt.Sprintf("%.4f%%", v))
	}
	elapsed := time.Since(start)
	fmt.Printf("%s took %s", category, elapsed)
}
