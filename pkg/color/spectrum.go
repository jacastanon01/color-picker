package color

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func GenerateSpectrum(w, h int32) *rl.Image {
	const min, max float32 = 0.5, 1.0
	image := rl.GenImageColor(int(w), int(h), rl.Blank)

	for y := int32(0); y < h; y++ {
		for x := int32(0); x < w; x++ {
			var hue float32 = scaleValue((float32(x) / float32(w)), 360, 0)
			var saturation float32 = 1 // default to full saturation
			var brightness float32 = scaleValue(float32(y)/float32(h), max, min)

			imageRGB := rl.ColorFromHSV(hue, saturation, brightness)
			rl.ImageDrawPixel(
				image, x, y, imageRGB,
			)
		}
	}
	return image
}

func scaleValue(value, max, min float32) float32 {
	return (value * (max - min)) + min
}
