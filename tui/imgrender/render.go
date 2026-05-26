package imgrender

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"math"
	"os"
)

// RenderMode defines the avatar rendering style
type RenderMode int

const (
	ModeHalfBlock RenderMode = iota
	ModeASCII
	ModeBraille
	ModeHidden
)

// ConvertImageToString loads an image from a file path and returns its TUI representation.
func ConvertImageToString(filePath string, targetWidth int, mode RenderMode) (string, error) {
	if mode == ModeHidden {
		return "", nil
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	img, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		return "", err
	}

	// Crop to focus on face and hair for ASCII & Braille modes (square 56% region)
	if mode == ModeASCII || mode == ModeBraille {
		img = cropImage(img, 0.22, 0.02, 0.56, 0.56)
	}

	bounds := img.Bounds()
	imgWidth := bounds.Dx()
	imgHeight := bounds.Dy()
	aspectRatio := float64(imgHeight) / float64(imgWidth)

	// In terminal, characters are roughly 2x taller than they are wide.
	// We need to account for this aspect ratio distortion.
	var targetHeight int
	switch mode {
	case ModeHalfBlock:
		// Half-block renders 2 vertical pixels per character cell
		targetHeight = int(float64(targetWidth) * aspectRatio)
	case ModeASCII:
		// Standard ASCII has 1 pixel per character cell. Since cells are tall,
		// we halve the height to avoid a stretched output.
		targetHeight = int(float64(targetWidth) * aspectRatio * 0.5)
	case ModeBraille:
		// Enforce a physically square rendering shape by making character height exactly half of width
		targetHeight = targetWidth / 2
	}

	if targetHeight < 1 {
		targetHeight = 1
	}

	switch mode {
	case ModeHalfBlock:
		// We need an even height of source pixels because 1 char = 2 vertical pixels
		resized := resizeImage(img, targetWidth, targetHeight*2)
		return renderHalfBlock(resized, targetWidth, targetHeight), nil
	case ModeASCII:
		resized := resizedToDimensions(img, targetWidth, targetHeight)
		return renderASCII(resized, targetWidth, targetHeight), nil
	case ModeBraille:
		resized := resizedToDimensions(img, targetWidth*2, targetHeight*4)
		return renderBraille(resized, targetWidth, targetHeight), nil
	}

	return "", nil
}

// cropImage cuts a sub-rectangle out of the image based on percentage bounds
func cropImage(img image.Image, xPercent, yPercent, wPercent, hPercent float64) image.Image {
	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()

	minX := bounds.Min.X + int(float64(width)*xPercent)
	minY := bounds.Min.Y + int(float64(height)*yPercent)
	maxX := minX + int(float64(width)*wPercent)
	maxY := minY + int(float64(height)*hPercent)

	rect := image.Rect(0, 0, maxX-minX, maxY-minY)
	cropped := image.NewRGBA(rect)

	for y := 0; y < rect.Dy(); y++ {
		for x := 0; x < rect.Dx(); x++ {
			cropped.Set(x, y, img.At(minX+x, minY+y))
		}
	}
	return cropped
}

// resizeImage uses simple nearest-neighbor interpolation to resize images without external dependencies.
func resizeImage(img image.Image, width, height int) image.Image {
	rect := image.Rect(0, 0, width, height)
	resized := image.NewRGBA(rect)
	bounds := img.Bounds()
	origW := bounds.Dx()
	origH := bounds.Dy()

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			origX := bounds.Min.X + (x * origW / width)
			origY := bounds.Min.Y + (y * origH / height)
			resized.Set(x, y, img.At(origX, origY))
		}
	}
	return resized
}

// Helper with name matching the main render switch
func resizedToDimensions(img image.Image, w, h int) image.Image {
	return resizeImage(img, w, h)
}

// renderHalfBlock converts an image to colored half-blocks (▄)
func renderHalfBlock(img image.Image, w, h int) string {
	var buf bytes.Buffer
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			topColor := img.At(x, y*2)
			bottomColor := img.At(x, y*2+1)

			tr, tg, tb, _ := topColor.RGBA()
			br, bg, bb, _ := bottomColor.RGBA()

			// Scale down from 16-bit to 8-bit color channels
			r1, g1, b1 := tr>>8, tg>>8, tb>>8
			r2, g2, b2 := br>>8, bg>>8, bb>>8

			// Write ANSI escapes: background for top pixel (\x1b[48;2;r;g;bm), foreground for bottom (\x1b[38;2;r;g;bm)
			buf.WriteString(fmt.Sprintf("\x1b[48;2;%d;%d;%dm\x1b[38;2;%d;%d;%dm▄\x1b[0m", r1, g1, b1, r2, g2, b2))
		}
		buf.WriteByte('\n')
	}
	return buf.String()
}

// ditherImage converts an image to a binary boolean grid (black=false, white=true) using Floyd-Steinberg dithering with contrast scaling
func ditherImage(img image.Image, w, h int, contrast float64) [][]bool {
	lum := make([][]float64, h)
	for y := 0; y < h; y++ {
		lum[y] = make([]float64, w)
		for x := 0; x < w; x++ {
			c := img.At(x, y)
			r, g, b, _ := c.RGBA()
			r8, g8, b8 := r>>8, g>>8, b>>8
			yVal := 0.299*float64(r8) + 0.587*float64(g8) + 0.114*float64(b8)

			// Enhance contrast to separate foreground and background
			yVal = (yVal-127.5)*contrast + 127.5
			if yVal < 0 {
				yVal = 0
			} else if yVal > 255 {
				yVal = 255
			}
			lum[y][x] = yVal
		}
	}

	dithered := make([][]bool, h)
	for y := 0; y < h; y++ {
		dithered[y] = make([]bool, w)
	}

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			oldVal := lum[y][x]
			var newVal float64
			if oldVal > 127 {
				newVal = 255
				dithered[y][x] = true
			} else {
				newVal = 0
				dithered[y][x] = false
			}
			err := oldVal - newVal

			if x+1 < w {
				lum[y][x+1] += err * 7.0 / 16.0
			}
			if y+1 < h {
				if x-1 >= 0 {
					lum[y+1][x-1] += err * 3.0 / 16.0
				}
				lum[y+1][x] += err * 5.0 / 16.0
				if x+1 < w {
					lum[y+1][x+1] += err * 1.0 / 16.0
				}
			}
		}
	}
	return dithered
}

// renderASCII converts an image to monochrome ASCII characters using dithered mapping
func renderASCII(img image.Image, w, h int) string {
	// Character ramp from dark to light
	ramp := []rune(" .:-=+*#%@")
	rampLen := len(ramp)

	// Create float buffer for luminance
	lum := make([][]float64, h)
	for y := 0; y < h; y++ {
		lum[y] = make([]float64, w)
		for x := 0; x < w; x++ {
			c := img.At(x, y)
			r, g, b, _ := c.RGBA()
			r8, g8, b8 := r>>8, g>>8, b>>8
			yVal := 0.299*float64(r8) + 0.587*float64(g8) + 0.114*float64(b8)

			// Contrast enhancement for ASCII
			yVal = (yVal-127.5)*1.5 + 127.5
			if yVal < 0 {
				yVal = 0
			} else if yVal > 255 {
				yVal = 255
			}
			lum[y][x] = yVal
		}
	}

	var buf bytes.Buffer
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			oldVal := lum[y][x]
			valFactor := 255.0 / float64(rampLen-1)
			level := math.Round(oldVal / valFactor)
			if level < 0 {
				level = 0
			} else if level >= float64(rampLen) {
				level = float64(rampLen - 1)
			}
			newVal := level * valFactor
			err := oldVal - newVal

			// Diffuse error
			if x+1 < w {
				lum[y][x+1] += err * 7.0 / 16.0
			}
			if y+1 < h {
				if x-1 >= 0 {
					lum[y+1][x-1] += err * 3.0 / 16.0
				}
				lum[y+1][x] += err * 5.0 / 16.0
				if x+1 < w {
					lum[y+1][x+1] += err * 1.0 / 16.0
				}
			}

			buf.WriteRune(ramp[int(level)])
		}
		buf.WriteByte('\n')
	}
	return buf.String()
}

// renderBraille converts an image to Unicode Braille art (2x4 pixels per cell) using binary Floyd-Steinberg dithering
func renderBraille(img image.Image, w, h int) string {
	dithered := ditherImage(img, w*2, h*4, 1.5)

	dotMasks := [4][2]int{
		{0x01, 0x08},
		{0x02, 0x10},
		{0x04, 0x20},
		{0x40, 0x80},
	}

	var buf bytes.Buffer
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			brailleRune := 0x2800
			for dy := 0; dy < 4; dy++ {
				for dx := 0; dx < 2; dx++ {
					if dithered[y*4+dy][x*2+dx] {
						brailleRune |= dotMasks[dy][dx]
					}
				}
			}

			if brailleRune == 0x2800 {
				buf.WriteString(" ")
			} else {
				buf.WriteRune(rune(brailleRune))
			}
		}
		buf.WriteByte('\n')
	}
	return buf.String()
}

// ConvertColorToLipgloss helper to map standard color to Lipgloss color
func ConvertColorToLipgloss(c color.Color) string {
	r, g, b, _ := c.RGBA()
	return fmt.Sprintf("#%02x%02x%02x", r>>8, g>>8, b>>8)
}
