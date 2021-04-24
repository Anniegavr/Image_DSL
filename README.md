# Image_DSL
PBL4 project

Actions: left side is what the user shall input
"resizeSq":          resizeSq(), resize picture in the shape of a square of size n (input)
		"resize":            resizeManual(),  resize to the dimmensions you choose
		"resize2fit":        resize2fit(),   resize to fit in the box of x by y
		"resize2fill":       resize2fill(),   resize to fill the space of a box of x by y

		"cropSq":            cropSq(),      crop into a square
		"crop":         	 cropManual(),    crop x by y

		"rotate":            rotate(),   rotate by x degrees
		"rotate30":          gift.Rotate(30, color.Transparent, gift.CubicInterpolation),
		"rotate90":          gift.Rotate90(),
		"rotate180":         gift.Rotate180(),
		"rotate270":         gift.Rotate270(),

		"flipHor":  		 gift.FlipHorizontal(),
		"flipVert":			 gift.FlipVertical(),

		"brightness":        bright(),     brightness up or down by x percents
		"brightnessUp":  	 gift.Brightness(30),
		"brightnessDn":  	 gift.Brightness(-30),

		"contrast":		 	 contrast(),  contrast up or down by x percents
		"contrastUp":    	 gift.Contrast(30),
		"contrastDn":    	 gift.Contrast(-30),

		"saturation":	 	 sats(),   saturation up or down by x percents
		"satsUp":  		 	 gift.Saturation(40),
		"satsDn":  			 gift.Saturation(-40),

		"pixelate":          pixels(),  pixelate into x pixels

		"sepia":             gift.Sepia(100),
		"sepiaX":			 sepia(),   sepia filter - up to the percentage chosen by user
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
