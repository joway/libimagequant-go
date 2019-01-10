# libimagequant-go

![GitHub release](https://img.shields.io/github/tag/joway/libimagequant-go.svg?label=release)
[![Go Report Card](https://goreportcard.com/badge/github.com/joway/libimagequant-go)](https://goreportcard.com/report/github.com/joway/libimagequant-go)
[![codecov](https://codecov.io/gh/joway/libimagequant-go/branch/master/graph/badge.svg)](https://codecov.io/gh/joway/libimagequant-go)
[![CircleCI](https://circleci.com/gh/joway/libimagequant-go.svg?style=shield)](https://circleci.com/gh/joway/libimagequant-go)

A golang wrapper for [libimagequant](https://github.com/ImageOptim/libimagequant).

## Used by

- [imagic](https://github.com/joway/imagic): An easy and fast tool to process images.

## Install

```shell
go get github.com/joway/libimagequant-go
```

## Usage

### High Level API

```
// compress png iamge
func Compress(img image.Image, quality int, speed int) (image.Image, error)
```

### Low Level API

[GoDoc](https://godoc.org/github.com/joway/libimagequant-go/pngquant)

## Example

```golang
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
	quality := 70
	speed := 5 // 1~10
	output, err := pngquant.Compress(img, quality, speed)
}
```
