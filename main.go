package main

import (
	"fmt"
	"github.com/joway/libimagequant-go/pngquant"
	"image/png"
	"os"
)

func main() {
	source, _ := os.OpenFile("testdata/1.png", os.O_RDONLY, 0444)
	img, _ := png.Decode(source)
	output, err := pngquant.Compress(img, 10)

	fmt.Println(output)
	fmt.Println(err)
}
