package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	//"image/jpeg"
	//"image/gif"
	"log"
	"os"

	"github.com/disintegration/gift"
)
func main() {
	fmt.Println("Enter image and the path to it: ")
	var fille string
	fmt.Scanln(&fille)
	src := loadImage(fille)
	fmt.Print("Action: ")
	var act string
	fmt.Scanln(&act)
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


		//"mean":                 gift.Mean(5, true),
		//"median":               gift.Median(5, true),
		//"minimum":              gift.Minimum(5, true),
		//"maximum":              gift.Maximum(5, true),
		//"hue_rotate":           gift.Hue(45),
		//"color_balance":        gift.ColorBalance(10, -10, -10),
		//"color_func": gift.ColorFunc(
		//	func(r0, g0, b0, a0 float32) (r, g, b, a float32) {
		//		r = 1 - r0   // invert the red channel
		//		g = g0 + 0.1 // shift the green channel by 0.1
		//		b = 0        // set the blue channel to 0
		//		a = a0       // preserve the alpha channel
		//		return r, g, b, a
		//	},
		//),
		//"convolution_emboss": gift.Convolution(
		//	[]float32{
		//		-1, -1, 0,
		//		-1, 1, 1,
		//		0, 1, 1,
		//	},
		//	false, false, false, 0.0,
		//),
	}

	fmt.Println("Save image? Y/N\n")
	var ans string
	fmt.Scanln(&ans)
	if ans == "Y" || ans == "y" {
		for name, action := range actions {
			g := gift.New(action)
			dst := image.NewNRGBA(g.Bounds(src.Bounds()))
			g.Draw(dst, src)
			saveImage("..\\z-product_"+name+".png", dst) //path to image + name of the generated image
			//to be reviewed 
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


