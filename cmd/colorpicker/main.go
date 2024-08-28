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

	image := color.GenerateSpectrum(w, h)
	texture := rl.LoadTextureFromImage(image)
	colorpicker := &color.ColorPicker{Image: image, Texture: texture}

	displayData := make(map[rl.Vector2]rl.Color)
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)
	var pos rl.Vector2

	for !rl.WindowShouldClose() && rl.IsWindowFocused() {

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		rl.DrawTexture(texture, 0, 0, rl.White)

		if rl.IsMouseButtonReleased(rl.MouseButtonLeft) {
			pos = rl.GetMousePosition()
			colors := colorpicker.GetColorAtPosition(int32(pos.X), int32(pos.Y))
			displayData[pos] = colors
		}

		for pos, colors := range displayData {
			text := fmt.Sprintf("R: %d\nG: %d\nB: %d\n", colors.R, colors.G, colors.B)
			rl.DrawText(text, int32(pos.X), int32(pos.Y), 10, rl.White)
		}

		rl.EndDrawing()
	}
	rl.UnloadTexture(texture)
	rl.UnloadImage(image)
}
