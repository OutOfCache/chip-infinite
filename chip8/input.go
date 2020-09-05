package chip8

import (
	"github.com/veandco/go-sdl2/sdl"
)

var keys [16]bool

// HandleInput checks the user input and sets variables accordingly.
// If the user requested to quit the application, the program will exit.
// If the user did not, it will store the key states of every key mapped to
// the hex keypad in keys as bool (true if pressed, false if not)
func HandleInput() {
	for e := sdl.PollEvent(); e != nil; e = sdl.PollEvent() {
		switch e.(type) {
		case *sdl.QuitEvent:
			Quit = true
		}
		currentKeys := sdl.GetKeyboardState()
		keys[0x0] = currentKeys[sdl.SCANCODE_X] == 1
		keys[0x1] = currentKeys[sdl.SCANCODE_1] == 1
		keys[0x2] = currentKeys[sdl.SCANCODE_2] == 1
		keys[0x3] = currentKeys[sdl.SCANCODE_3] == 1
		keys[0x4] = currentKeys[sdl.SCANCODE_Q] == 1
		keys[0x5] = currentKeys[sdl.SCANCODE_W] == 1
		keys[0x6] = currentKeys[sdl.SCANCODE_E] == 1
		keys[0x7] = currentKeys[sdl.SCANCODE_A] == 1
		keys[0x8] = currentKeys[sdl.SCANCODE_S] == 1
		keys[0x9] = currentKeys[sdl.SCANCODE_D] == 1
		keys[0xA] = currentKeys[sdl.SCANCODE_Z] == 1
		keys[0xB] = currentKeys[sdl.SCANCODE_C] == 1
		keys[0xC] = currentKeys[sdl.SCANCODE_4] == 1
		keys[0xD] = currentKeys[sdl.SCANCODE_R] == 1
		keys[0xE] = currentKeys[sdl.SCANCODE_F] == 1
		keys[0xF] = currentKeys[sdl.SCANCODE_V] == 1
	}
}

func waitForKey() int {
	var pressed int = -1

	for e := sdl.PollEvent(); e != nil; e = sdl.PollEvent() {
		switch t := e.(type) {
		case *sdl.QuitEvent:
			Quit = true
		case *sdl.KeyboardEvent:
			keycode := t.Keysym.Sym

			// TODO: find a prettier solution for this
			switch keycode {
			case sdl.K_x:
				pressed = 0x0
			case sdl.K_1:
				pressed = 0x1
			case sdl.K_2:
				pressed = 0x2
			case sdl.K_3:
				pressed = 0x3
			case sdl.K_q:
				pressed = 0x4
			case sdl.K_w:
				pressed = 0x5
			case sdl.K_e:
				pressed = 0x6
			case sdl.K_a:
				pressed = 0x7
			case sdl.K_s:
				pressed = 0x8
			case sdl.K_d:
				pressed = 0x9
			case sdl.K_z:
				pressed = 0xA
			case sdl.K_c:
				pressed = 0xB
			case sdl.K_4:
				pressed = 0xC
			case sdl.K_r:
				pressed = 0xD
			case sdl.K_f:
				pressed = 0xE
			case sdl.K_v:
				pressed = 0xF
			}
			return pressed
		}

	}
	return pressed
}
