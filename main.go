package main

import (
	"fmt"
	_ "image/jpeg"
	"io/ioutil"

	"github.com/disintegration/imaging"
)

func main() {
	imagePath := "/home/ardanhilal/Documents/coding/golang/imageResize"
	files, _ := ioutil.ReadDir(fmt.Sprintf("%s/%s", imagePath, "data"))
	count := 0
	for _, f := range files {
		src, _ := imaging.Open(fmt.Sprintf("%s/%s/%s", imagePath, "data", f.Name()))

		transformedImage := imaging.Resize(src, 320, 240, imaging.Lanczos)

		imaging.Save(transformedImage, fmt.Sprintf("%s/saved/%s", imagePath, f.Name()))

		count++
		print := fmt.Sprintf("image with name %s has been saved", f.Name())
		fmt.Println(print)
	}

	print(count)
}
