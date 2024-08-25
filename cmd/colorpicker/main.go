package main

import rl "github.com/gen2brain/raylib-go/raylib"

type HSB struct {
	Hue        float32 // 0 to 360
	Saturation float32 // 0 to 1.0
	Brightness float32 // 0 to 1.0
}

func main() {
	const w int32 = 900
	const h int32 = 450
	const min float32 = 0.5
	const max float32 = 1.0

	rl.InitWindow(w, h, "raylib [core] example - basic window")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() && rl.IsWindowFocused() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)
		rl.DrawText("Congrats! You created your first window!", 190, 200, 20, rl.LightGray)

		spectrum := make([][]HSB, h)
		for y := range spectrum {
			spectrum[y] = make([]HSB, w)
			for x := range spectrum[y] {
				spectrum[y][x].Hue = float32(x) / float32(w) * 360
				spectrum[y][x].Saturation = 1.0
				percent := float32(y) / float32(h) // percentage of height
		
				spectrum[y][x].Brightness = (percent * (max - min)) + min

				rl.DrawPixel(
					int32(x), int32(y), rl.ColorFromHSV(spectrum[y][x].Hue, spectrum[y][x].Saturation, spectrum[y][x].Brightness),
				)
			}
		}
		rl.EndDrawing()
	}
}
