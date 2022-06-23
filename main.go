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
	files, err := ioutil.ReadDir("./" + category)
	if err != nil {
		log.Fatal(err)
	}

	imageProcessor.Compare("1d25ea94-4562-4e19-848e-b60f1b58deee.raw", files, category)
	elapsed := time.Since(start)
	fmt.Printf("%s took %s", category, elapsed)
}
