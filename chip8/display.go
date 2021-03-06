package chip8

import (
	"github.com/veandco/go-sdl2/sdl"
	"os"

	"fmt"
)

// Display represents the 64x32 display as a one-dimensional array, with
// Display[0] being the pixel in the top left corner and Display[2047]
// being the pixel in the bottom right
var Display [2048]byte

// sdl window dimensions
// TODO: variable scaling factor?
var winWidth, winHeight int32 = 64 * 8, 32 * 8
var err error

var gWindow *sdl.Window
var gRenderer *sdl.Renderer

// Quit controls whether we want to exit the main loop and the program
// TODO: get rid of Quit to avoid unnecessary comparison in main loop
var Quit bool

var fontset [80]byte = [80]byte{
	0xF0, 0x90, 0x90, 0x90, 0xF0, // 0
	0x20, 0x60, 0x20, 0x20, 0x70, // 1
	0xF0, 0x10, 0xF0, 0x80, 0xF0, // 2
	0xF0, 0x10, 0xF0, 0x10, 0xF0, // 3
	0x90, 0x90, 0xF0, 0x10, 0x10, // 4
	0xF0, 0x80, 0xF0, 0x10, 0xF0, // 5
	0xF0, 0x80, 0xF0, 0x90, 0xF0, // 6
	0xF0, 0x10, 0x20, 0x40, 0x40, // 7
	0xF0, 0x90, 0xF0, 0x90, 0xF0, // 8
	0xF0, 0x90, 0xF0, 0x10, 0xF0, // 9
	0xF0, 0x90, 0xF0, 0x90, 0x90, // A
	0xE0, 0x90, 0xE0, 0x90, 0xE0, // B
	0xF0, 0x80, 0x80, 0x80, 0xF0, // C
	0xE0, 0x90, 0x90, 0x90, 0xE0, // D
	0xF0, 0x80, 0xF0, 0x80, 0xF0, // E
	0xF0, 0x80, 0xF0, 0x80, 0x80, // F
}

// StartSDL initializes SDL and creates the window and returns true if it
// succeeded and returns false if there was and error
// the following start and end functions are taken from Lazy Foo Production's
// SDL Tutorials found at http://lazyfoo.net
// TODO: use log instead?
func StartSDL() {
	// Initialize SDL
	if err = sdl.Init(sdl.INIT_VIDEO | sdl.INIT_AUDIO); err != nil {
		fmt.Printf("SDL could not initialize! Error: %s\n", err)
		os.Exit(3) // 3 is for SDL errors now
	}
	// create Window
	gWindow, err = sdl.CreateWindow("Chip 8", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, winWidth, winHeight, sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Printf("Window could not be created! Error: %s\n", err)
		os.Exit(3)
	}
	gRenderer, err = sdl.CreateRenderer(gWindow, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Printf("Renderer could not be created! SDL Error: %s\n", err)
		os.Exit(3)
	}
}

// End frees the resources needed by sdl and quits sdl entirely
func End() {
	gRenderer.Destroy()
	gWindow.Destroy()

	sdl.CloseAudioDevice(device) // from audio.go

	sdl.Quit()
}

// Render draws an 8x8 rectangle on the screen for every pixel
// and updates the screen
func Render() {
	gRenderer.SetDrawColor(0x00, 0x00, 0x00, 0xFF)
	gRenderer.Clear()

	// set every pixel to the right color
	for i, pixel := range Display {
		var y int32 = int32(8 * (i / 64))
		var x int32 = int32(8 * (i % 64))
		rect := sdl.Rect{x, y, 8, 8}
		if pixel != 0 {
			gRenderer.SetDrawColor(0xFF, 0xFF, 0xFF, 0xFF)
		} else {
			gRenderer.SetDrawColor(0x00, 0x00, 0x00, 0xFF)
		}
		gRenderer.FillRect(&rect)
	}

	gRenderer.Present()
}
