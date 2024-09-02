package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/jacastanon01/color-picker/pkg/color"
)

func main() {
	const w int32 = 900
	const h int32 = 450

	rl.InitWindow(w, h, "Go Color Picker")

	colorpicker := initColorPicker(w, h)

	displayData := make(map[rl.Vector2]string)
	defer rl.CloseWindow()
	defer rl.UnloadTexture(colorpicker.Texture)
	defer rl.UnloadImage(colorpicker.Image)

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() && rl.IsWindowFocused() {

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		rl.DrawTexture(colorpicker.Texture, 0, 0, rl.White)

		if rl.IsMouseButtonReleased(rl.MouseButtonLeft) {
			pos := rl.GetMousePosition()
			colors := colorpicker.GetColorAtPosition(int32(pos.X), int32(pos.Y))
			displayData[pos] = fmt.Sprintf("R: %d\nG: %d\nB: %d\n", colors.R, colors.G, colors.B)
		}

		for pos, text := range displayData {
			rl.DrawText(text, int32(pos.X), int32(pos.Y), 10, rl.White)
		}

		rl.EndDrawing()
	}
}

func initColorPicker(w, h int32) *color.ColorPicker {
	image := color.GenerateSpectrum(w, h)
	texture := rl.LoadTextureFromImage(image)
	colorpicker := &color.ColorPicker{Image: image, Texture: texture}
	return colorpicker
}
