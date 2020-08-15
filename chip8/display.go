package chip8

import(
    "github.com/veandco/go-sdl2/gfx"
    "github.com/veandco/go-sdl2/sdl"

    "fmt"
)

var display [2048]bool

// sdl window dimensions
var winWidth, winHeight int32 = 64, 32
var err error

var gWindow *sdl.Window
var gScreenSurface *sdl.Surface

func start() bool {
    // Initialization flag
    var success bool = true

    // Initialize SDL
    if err = sdl.Init(sdl.INIT_VIDEO); err != nil {
	fmt.Printf("SDL could not initialize! Error: %s\n", err)
	success = false
    } else {
	// create Window
	gWindow, err = sdl.CreateWindow("Chip 8", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, winWidth, winHeight, sdl.WINDOW_SHOWN)
	if (err != nil) {
	    fmt.Printf("Window could not be created! Error: %s\n", err)
	    success = false
	} else {
	    // get Window surface
	    gScreenSurface, err = gWindow.GetSurface()
	}
    }
    return success
}

func loadMedia() bool {
    // loading success flag
    var success bool = true
}

func main() {
    // the window we'll be rendering to
