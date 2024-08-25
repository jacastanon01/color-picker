package color

type HSB struct {
	Hue float64 // 0 to 360
	Saturation float64 // 0 to 1.0
	Brightness float64 // 0 to 1.0
}

type RGB struct {
	Red uint8 // 0 to 255
	Green uint8	// 0 to 255
	Blue uint8 // 0 to 255
}
