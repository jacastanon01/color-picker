package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/jacastanon01/color-picker/pkg/color"
)

func main() {
	const w int32 = 900
	const h int32 = 450

	rl.InitWindow(w, h, "Go Color Picker")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() && rl.IsWindowFocused() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		color.GenerateSpectrum(w, h)
		color.DisplayPosText()

		rl.EndDrawing()
	}
}
