package main

import (
	"image"
	"image/color"
	"tour_of_go/pkg/display"
)

func Pic(dx, dy int) [][]uint8 {
	s := make([][]uint8, dy)
	for i := range s {
		s[i] = make([]uint8, dx)
		for j := range s[i] {
			// s[i][j] = uint8(i) ^ uint8(j)
			// s[i][j] = uint8(i) ^ uint8(j)
			s[i][j] = (uint8(i) + uint8(j)) / 2
		}
	}
	return s
}

func main() {
	const width, height = 256, 256

	data := Pic(width, height)

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for y := range height {
		for x := range width {
			v := data[y][x]
			img.Set(x, y, color.RGBA{v, v, 255, 255})
		}
	}

	display.DisplayImageInTerminal(img)
}
