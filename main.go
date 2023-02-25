package main

import (
	"fmt"
	_ "image/jpeg"
	"io/ioutil"

	"github.com/disintegration/imaging"
)

func main() {
	imagePath := "./data"
	files, _ := ioutil.ReadDir(fmt.Sprintf("%s/", imagePath))
	count := 0
	for _, f := range files {
		src, _ := imaging.Open(fmt.Sprintf("%s/%s", imagePath, f.Name()))

		transformedImage := imaging.Resize(src, 320, 240, imaging.Lanczos)
		imaging.Save(transformedImage, fmt.Sprintf("./saved/%s", f.Name()))

		count++
		print := fmt.Sprintf("image with name %s has been saved", f.Name())
		fmt.Println(print)
	}

	print(count)
}
