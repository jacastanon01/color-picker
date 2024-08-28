package color

import rl "github.com/gen2brain/raylib-go/raylib"

type ColorPicker struct {
	Image   *rl.Image
	Texture rl.Texture2D
}

func (cp *ColorPicker) GetColorAtPosition(x, y int32) rl.Color {
	return rl.GetImageColor(*cp.Image, x, y)
}
