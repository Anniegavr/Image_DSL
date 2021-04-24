package main

import (
	//"container/list"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	//"sync"

	//"github.com/cheekybits/genny/generic"
	//"image/jpeg"
	//"image/gif"
	"log"
	"os"

	"github.com/disintegration/gift"
)
//Resize functions
func resizeSq() gift.Filter{
	var dim int
	fmt.Println("Square dimensions: ")
	fmt.Scanln(&dim)
	//return dim, dim
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

//Crop and rotate functions
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

//Color functions
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
func main() {

	fmt.Println("Enter image and the path to it: ")
	var fille string

	fmt.Scanln(&fille)

	var num_act int
	fmt.Println("How many actions will you perform on the image?")
	fmt.Scanln(&num_act)


	//var act_list[] string




	//for ac := range act {}
	//Prototype definition of actions (left side: what the user should use as input)

	actions := map[string]gift.Filter{
		"resizeSq":          resizeSq(),
		"resize":            resizeManual(),
		"resize2fit":        resize2fit(),
		"resize2fill":       resize2fill(),

		"cropSq":            cropSq(),
		"crop":         	 cropManual(),

		"rotate":            rotate(),
		"rotate30":          gift.Rotate(30, color.Transparent, gift.CubicInterpolation),
		"rotate90":          gift.Rotate90(),
		"rotate180":         gift.Rotate180(),
		"rotate270":         gift.Rotate270(),

		"flipHor":  		 gift.FlipHorizontal(),
		"flipVert":			 gift.FlipVertical(),

		"brightness":        bright(),
		"brightnessUp":  	 gift.Brightness(30),
		"brightnessDn":  	 gift.Brightness(-30),

		"contrast":		 	 contrast(),
		"contrastUp":    	 gift.Contrast(30),
		"contrastDn":    	 gift.Contrast(-30),

		"saturation":	 	 sats(),
		"satsUp":  		 	 gift.Saturation(40),
		"satsDn":  			 gift.Saturation(-40),

		"pixelate":          pixels(),

		"sepia":             gift.Sepia(100),
		"sepiaX":			 sepia(),
		"sepia50":		 	 gift.Sepia(50),

		"gammaX":			 gamma(),
		"gamma_1.5":         gift.Gamma(1.5),
		"gamma_0.5":         gift.Gamma(0.5),

		"invert":            gift.Invert(),

		"gaussianBlur":        gaus_blur(),

		"unsharp_mask":         gift.UnsharpMask(1, 1, 0),



		"sigmoid":              gift.Sigmoid(0.5, 7),

		"colorize":             gift.Colorize(240, 50, 100),
		"grayscale":            gift.Grayscale(),
		//"grayscaleX":			grayscaleX(),  //Allows the user to set the percentage of discoloration


		"mean":                 gift.Mean(5, true),
		"median":               gift.Median(5, true),
		"minimum":              gift.Minimum(5, true),
		"maximum":              gift.Maximum(5, true),
		"hue_rotate":           gift.Hue(45),
		"color_balance":        gift.ColorBalance(10, -10, -10),
		"color_func": gift.ColorFunc(
			func(r0, g0, b0, a0 float32) (r, g, b, a float32) {
				r = 1 - r0   // invert the red channel
				g = g0 + 0.1 // shift the green channel by 0.1
				b = 0        // set the blue channel to 0
				a = a0       // preserve the alpha channel
				return r, g, b, a
			},
		),
		"convolution_emboss": gift.Convolution(
			[]float32{
				-1, -1, 0,
				-1, 1, 1,
				0, 1, 1,
			},
			false, false, false, 0.0,
		),
	}

	var to_do [] gift.Filter
	var i int
	for i=0; i<num_act; i++{
		var act string
		fmt.Print("Actions: ")
		fmt.Scanln(&act)
		to_do = append(to_do, actions[act])
		//for n := range actions {
		//	if n == act{
		//		to_do = append(to_do, actions[act])
		//	}
		//}


	}
	fmt.Println("Save image? Y/N\n")
	var ans string
	fmt.Scanln(&ans)
	//var j int
	if ans == "Y" || ans == "y" {
		//for name, action := range actions {
		//for name, action := range actions {
		//	if to_do = c {
		//		to_do[name] = action
		//	}
		//}
		for name := range to_do{
			//var action string
			src := loadImage(fille)
			//action = actions[name]
			g := gift.New(to_do[name])
			dst := image.NewNRGBA(g.Bounds(src.Bounds()))
			g.Draw(dst, src)
			saveImage("N:\\MINE\\Mentalists\\PBL4\\code\\dsl\\z-product_.png", dst) //path to image + name of the generated image
				//to be reviewed
			//for j := range act_list{
			//
			//}
			}

	} else{
		fmt.Println("Ok")
	}
}

func loadImage(filename string) image.Image {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("os.Open failed: %v", err)
	}
	defer f.Close()
	img, _, err := image.Decode(f)
	if err != nil {
		log.Fatalf("image.Decode failed: %v", err)
	}
	return img
}

func saveImage(filename string, img image.Image) {
	f, err := os.Create(filename)
	if err != nil {
		log.Fatalf("os.Create failed: %v", err)
	}
	defer f.Close()
	err = png.Encode(f, img)
	if err != nil {
		log.Fatalf("png.Encode failed: %v", err)
	}
}
