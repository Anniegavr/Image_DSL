package main

import (
    "fmt"
    "github.com/disintegration/gift"
)

func resizeSq() gift.Filter {
    var dim int
    fmt.Println("Square dimensions: ")
    fmt.Scanln(&dim)
    return gift.Resize(dim, dim, gift.CubicResampling)
}

func resizeManual() gift.Filter {  // Doesn't preserve the original aspect ratio
    var h int
    var w int
    fmt.Println("Height: ")
    fmt.Scanln(&h)
    fmt.Println("Length: ")
    fmt.Scanln(&w)
    return gift.Resize(h, w, gift.CubicResampling)
}

func resize2fit() gift.Filter { // Preserves the original aspect ratio
    var h int
    var w int
    fmt.Println("Height: ")
    fmt.Scanln(&h)
    fmt.Println("Length: ")
    fmt.Scanln(&w)
    return gift.ResizeToFit(h, w, gift.CubicResampling)
}

func resize2fill() gift.Filter { // Preserves the original aspect ratio
    var h int
    var w int
    var resAnchorOption string
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
    fmt.Scanln(&w)
    fmt.Println("Anchor: \n 1.Top Anchor\n2.Top Left Anchor\n3.Top Right Anchor\n4.Center Anchor\n5.Bottom Anchor\n6.Bottom Left Anchor\n7.Bottom Right Anchor\n")
    fmt.Scanln(&resAnchorOption)
    return gift.ResizeToFill(h, w, gift.BoxResampling, anchor[resAnchorOption])
}
