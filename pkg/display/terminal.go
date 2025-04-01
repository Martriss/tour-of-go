package display

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/png"
	"os"
	"strings"
)

// DisplayImageInTerminal displays an image directly in Wezterm or Ghostty terminal
func DisplayImageInTerminal(img image.Image) {
	// Create a buffer to store the PNG data
	var buf bytes.Buffer
	err := png.Encode(&buf, img)
	if err != nil {
		panic(err)
	}

	// Encode the PNG data as base64
	imgBase64 := base64.StdEncoding.EncodeToString(buf.Bytes())

	// Determine which terminal emulator we're using
	term := os.Getenv("TERM")
	termProgram := os.Getenv("TERM_PROGRAM")
	wezterm := strings.Contains(term, "wezterm") || strings.Contains(termProgram, "WezTerm")
	ghostty := strings.Contains(term, "ghostty") || strings.Contains(termProgram, "Ghostty")

	// Print the image using the appropriate escape sequences
	if wezterm || ghostty {
		// Both Wezterm and Ghostty support iTerm2 graphics protocol
		fmt.Printf("\033]1337;File=inline=1;width=auto;height=auto:%s\a\n", imgBase64)
	} else {
		fmt.Println("Your terminal doesn't appear to support inline images.")
		fmt.Println("Image saved as pattern.png instead.")

		// Save to file as fallback
		f, err := os.Create("pattern.png")
		if err != nil {
			panic(err)
		}
		defer f.Close()

		// Write the buffer to the file
		_, err = f.Write(buf.Bytes())
		if err != nil {
			panic(err)
		}
	}
}
