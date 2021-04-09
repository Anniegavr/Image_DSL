package main

import (
	"fmt"
	"github.com/disintegration/gift"
	"image"
	"image/draw"
)
type Filter interface {
	// Draw applies the filter to the src image and outputs the result to the dst image.
	Draw(dst draw.Image, src image.Image, options *Options)
	// Bounds calculates the appropriate bounds of an image after applying the filter.
	Bounds(srcBounds image.Rectangle) (dstBounds image.Rectangle)
}

// Options is the parameters passed to image processing filters.
type Options struct {
	Parallelization bool
}



type pixel struct{
	r, g, b, a float32
}
type colorFilter struct {
	fn func(pixel) pixel
}

func bright() gift.Filter {
	var strong int
	fmt.Println("Percentage: ")
	fmt.Scanln(&strong)
	return gift.Brightness(float32(strong))
}

func contrast() gift.Filter {
	var strong int
	fmt.Println("Percentage: ")
	fmt.Scanln(&strong)
	return gift.Contrast(float32(strong))
}

func sats() gift.Filter {
	var strong int
	fmt.Println("Percentage: ")
	fmt.Scanln(&strong)
	return gift.Saturation(float32(strong))
}

func pixels() gift.Filter {
	var size int
	fmt.Println("Size: ")
	fmt.Scanln(&size)
	return gift.Pixelate(size)
}

func sepia() gift.Filter {
	var strong int
	fmt.Println("Percentage: ")
	fmt.Scanln(&strong)
	return gift.Sepia(float32(strong))
}

func gamma() gift.Filter {
	var strong int
	fmt.Println("Percentage: ")
	fmt.Scanln(&strong)
	return gift.Gamma(float32(strong))
}

func gaus_blur() gift.Filter {
	var strong int
	fmt.Println("Percentage: ")
	fmt.Scanln(&strong)
	return gift.GaussianBlur(float32(strong))
}
//
//func(px pixel) pixel {
//	var r, g, b float32
//	fmt.Println("Decrease red (0 to 1): ")
//	fmt.Scanln(&r)
//	fmt.Println("Decrease green (0 to 1): ")
//	fmt.Scanln(&g)
//	fmt.Println("Decrease blue (0 to 1): ")
//	fmt.Scanln(&b)
//	y := r*px.r + g*px.g + b*px.b
//	return pixel{y, y, y, px.a}
//}
//
//func grayscaleX() gift.Filter {
//	var r, g, b float32
//
//	return gift.Filter{
//
//	}
//}


