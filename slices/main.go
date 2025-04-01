package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/color"
	"image/png"
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

func displayImageInTerminal(img image.Image) {
	var buf bytes.Buffer
	err := png.Encode(&buf, img)
	if err != nil {
		panic(err)
	}

	imgBase64 := base64.StdEncoding.EncodeToString(buf.Bytes())

	// This needs iTerm2 graphics protocol
	fmt.Printf("\033]1337;File=inline=1;width=auto;height=auto:%s\a\n", imgBase64)

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

	displayImageInTerminal(img)
}
