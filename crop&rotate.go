package main

import (
	"fmt"
	"github.com/disintegration/gift"
	"image/color"
)

func cropSq() gift.Filter {
	var dim int
	var anchor gift.Anchor
	fmt.Println("Square dimensions: ")
	fmt.Scanln(&dim)
	fmt.Println("Anchor: ")
	fmt.Scanln(&anchor)
	return gift.CropToSize(dim, dim, anchor)
}

func cropManual() gift.Filter {
	var h int
	var l int
	var CropAnchorOption string
	var anchor = map[string]gift.Anchor{
		"1" : gift.TopAnchor,
		"2" : gift.TopLeftAnchor,
		"3" : gift.TopRightAnchor,
		"4" : gift.CenterAnchor,
		"5" : gift.BottomAnchor,
		"6" : gift.BottomLeftAnchor,
		"7" : gift.BottomRightAnchor,
	}
	fmt.Println("Height: ")
	fmt.Scanln(&h)
	fmt.Println("Length: ")
	fmt.Scanln(&l)
	fmt.Println("Anchor: \n 1.Top Anchor\n2.Top Left Anchor\n3.Top Right Anchor\n4.Center Anchor\n5.Bottom Anchor\n6.Bottom Left Anchor\n7.Bottom Right Anchor\n")
	fmt.Scanln(&CropAnchorOption)
	return gift.CropToSize(h, l, anchor[CropAnchorOption])
}

func rotate() gift.Filter {
	var angle int
	fmt.Println("Angle: ")
	fmt.Scanln(&angle)
	return gift.Rotate(float32(angle), color.Transparent, gift.CubicInterpolation)
}



