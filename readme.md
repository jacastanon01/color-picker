# Go Color Picker

<img width="898" alt="Screenshot 2024-08-28 at 8 34 08 AM" src="https://github.com/user-attachments/assets/20fcbd2c-437b-4661-bd5f-12e091fd9c4f">

## Overview

This is an exploration of how pixels store color. My goal with this project is to learn more about differents way to generate data that can be used for colors. I used [raylib]("https://www.raylib.com/") to create a simple color spectrum where you can pick a color and see the resulting value of the pixel in RGB format.

## Understanding color and its measurement

Color is a visual perception created by the interaction of light with our eyes. Light consists of electromagnetic waves, and when it strikes an object, certain wavelengths are absorbed while others are reflected. Our eyes perceive these reflected wavelengths as color. In digital systems, color is often represented using the RGB (Red, Green, Blue) model, which combines these three primary colors of light at varying intensities to create a spectrum of colors. Alternatively, the HSL (Hue, Saturation, Lightness) model measures color differently by defining its hue (the type of color), saturation (the intensity of the color), and lightness (the brightness of the color). When creating a color palette with only two pigments, the resulting spectrum is limited, often focusing on gradients between the two, producing a harmonious and visually cohesive range of colors. I generated the HSL values based on pixel position and then converted that data into an RGB format:

```go
func GenerateSpectrum(w, h int32) *rl.Image {
	const min, max float32 = 0.5, 1.0
	image := rl.GenImageColor(int(w), int(h), rl.Blank)

	for y := int32(0); y < h; y++ {
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
```

## Made with

- Go's standard library
- Raylib
- ❤️ and 🧐

## How to run

- [Install Go](https://go.dev/doc/install)
- Clone this repo `git clone https://github.com/jacastanon01/color-picker.git`
- Navigate to the cloned directory. If using VSCode, use the build options to launch the program or run `go run cmd/colorpicker/main.go` at the root
