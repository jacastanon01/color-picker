package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/jacastanon01/color-picker/pkg/color"
)

func main() {
	const w int32 = 900
	const h int32 = 450
	rl.InitWindow(w, h, "Go Color Picker")

	image := color.GenerateSpectrum(w, h)
	texture := rl.LoadTextureFromImage(image)
	cachedImg := image
	// rl.UnloadImage(image)

	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() && rl.IsWindowFocused() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		rl.DrawTexture(texture, 0, 0, rl.White)

		if rl.IsMouseButtonReleased(rl.MouseButtonLeft) {
			// color.DisplayRGBText()
			pos := rl.GetMousePosition()
			color.DisplayColorAtPosition(cachedImg, int32(pos.X), int32(pos.Y))
		}

		rl.EndDrawing()
	}
	rl.UnloadTexture(texture)
	rl.UnloadImage(image)
}
