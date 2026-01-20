package main

import (
	"image/jpeg"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatal("引数が足りません")
	}
	imagePath := os.Args[1]
	w, err := os.Open(imagePath)
	if err != nil {
		log.Fatal(err)
	}

	image, err := jpeg.Decode(w)
	if err != nil {
		log.Fatal(err)
	}

	outputImagePath := os.Args[2]
	nw, err := os.Create(outputImagePath)
	if err != nil {
		log.Fatal(err)
	}

	if err := jpeg.Encode(nw, image, &jpeg.Options{Quality: 60}); err != nil {
		log.Fatal(err)
	}
}
