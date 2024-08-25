package color

func GenerateSpectrum() [][]HSB {
	spectrum := make([][]HSB, 100)
	for i := range spectrum {
		spectrum[i] = make([]HSB, 360)
		for j := range spectrum[i] {
			spectrum[i][j].Hue = float64(j)
			spectrum[i][j].Saturation = 1.0
			spectrum[i][j].Brightness = (float64(i) / 99) * 100
	}
}

	return spectrum
}
