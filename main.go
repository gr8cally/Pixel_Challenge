package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"pixel_challenge/imageProcessor"
	"time"
)

func main() {
	runCategory("Bronze/", "1d25ea94-4562-4e19-848e-b60f1b58deee.raw")
	//runCategory("Silver/", "0c9ec855-f5ad-4586-bf3b-da489f447219.raw")
	//runCategory("Gold/", "0a0f8f44-3b78-4bff-adee-14bc708e4ba7.raw")
}

func runCategory(category, image string) {
	start := time.Now()
	files, err := ioutil.ReadDir(category)
	if err != nil {
		log.Fatal(err)
	}

	top3 := imageProcessor.Compare(image, files, category)
	fmt.Printf("The closest images to %s in %s category are:\n", image, category)
	for k, v := range top3 {
		fmt.Println(k, fmt.Sprintf("%.4f%%", v))
	}
	elapsed := time.Since(start)
	fmt.Printf("%s took %s\n", category, elapsed)
}
