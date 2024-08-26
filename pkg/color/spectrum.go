package color

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func GenerateSpectrum(w, h int32) [][]HSB {
	const min, max float32 = 0.5, 1.0
	spectrum := make([][]HSB, h)
	for y := range spectrum {
		spectrum[y] = make([]HSB, w) // see if there is a way to insert this outside the loop. Calling make every loop may be iniefficient
		percent := float32(y) / float32(h)

		for x := range spectrum[y] {
			spectrum[y][x].Hue = GenerateThreshold((float32(x) / float32(w)), 0, 360)
			spectrum[y][x].Saturation = 1.0 // default to full saturation
			spectrum[y][x].Brightness = GenerateThreshold(percent, max, min)

			rl.DrawPixel(
				int32(x), int32(y), rl.ColorFromHSV(spectrum[y][x].Hue, spectrum[y][x].Saturation, spectrum[y][x].Brightness),
			)
		}
	}

	return spectrum
}

func DisplayPosText() {
	if rl.IsMouseButtonReleased(rl.MouseButtonLeft) {
		pos := rl.GetMousePosition()
		text := fmt.Sprintf("X: %d, Y: %d", int(pos.X), int(pos.Y))
		rl.DrawText(text, int32(pos.X-float32(25)), int32(pos.Y), 20, rl.White)
	}
}

func GenerateThreshold(value, max, min float32) float32 {
	return (value * (max - min)) + min
}
