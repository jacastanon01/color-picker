package main

import (
	"fmt"

	"github.com/jacastanon01/color-picker/pkg/color"
)

func main(){
	colors := color.GenerateSpectrum()
	fmt.Println(colors)
}