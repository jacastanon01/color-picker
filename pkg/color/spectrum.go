package color

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func GenerateSpectrum(w, h int32) *rl.Image {
	const min, max float32 = 0.5, 1.0
	image := rl.GenImageColor(int(w), int(h), rl.Blank)

	// spectrum := make([][]HSB, h)
	for y := int32(0); y < h; y++ {
		// spectrum[y] = make([]HSB, w) // see if there is a way to insert this outside the loop. Calling make every loop may be iniefficient
		for x := int32(0); x < w; x++ {
			var hue float32 = ScaleValue((float32(x) / float32(w)), 360, 0)
			var saturation float32 = 1 // default to full saturation
			var brightness float32 = ScaleValue(float32(y)/float32(h), max, min)

			imageRGB := rl.ColorFromHSV(hue, saturation, brightness)
			rl.ImageDrawPixel(
				image, int32(x), int32(y), imageRGB,
			)
		}
	}
	return image
}

func DisplayColorAtPosition(cachedImg *rl.Image, x, y int32) {
	colors := rl.GetImageColor(*cachedImg, x, y)

	text := fmt.Sprintf("R: %d G: %d B: %d", colors.R, colors.G, colors.B)
	rl.DrawText(text, x, y, 20, rl.White)
}

func ScaleValue(value, max, min float32) float32 {
	return (value * (max - min)) + min
}
