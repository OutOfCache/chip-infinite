package chip8

import(
    "github.com/veandco/go-sdl2/sdl"
)

var Keys [16]bool

// this code is a slightly modified version of Lazy Foo's SDL Tutorial Lesson 04
// found at http://lazyfoo.net
type KeyPressed int

const(
    KPZero  = iota
    KPOne
    KPTwo
    KPThree
    KPFour
    KPFive
    KPSix
    KPSeven
    KPEight
    KPNine
    KPA
    KPB
    KPC
    KPD
    KPE
    KPF
)

func HandleInput() int {
    var pressed int = -1 // for opcode 0xFx0A

    for e:= sdl.PollEvent(); e != nil; e = sdl.PollEvent() {
	switch e.(type) {
	case *sdl.QuitEvent:
	    Quit = true
	}
	currentKeys := sdl.GetKeyboardState()
	Keys[KPZero]	= currentKeys[sdl.SCANCODE_X] == 1
	Keys[KPOne]	= currentKeys[sdl.SCANCODE_1] == 1
	Keys[KPTwo]	= currentKeys[sdl.SCANCODE_2] == 1
	Keys[KPThree]	= currentKeys[sdl.SCANCODE_3] == 1
	Keys[KPFour]	= currentKeys[sdl.SCANCODE_Q] == 1
	Keys[KPFive]	= currentKeys[sdl.SCANCODE_W] == 1
	Keys[KPSix]	= currentKeys[sdl.SCANCODE_E] == 1
	Keys[KPSeven]	= currentKeys[sdl.SCANCODE_A] == 1
	Keys[KPEight]	= currentKeys[sdl.SCANCODE_S] == 1
	Keys[KPNine]	= currentKeys[sdl.SCANCODE_D] == 1
	Keys[KPA]	= currentKeys[sdl.SCANCODE_Z] == 1
	Keys[KPB]	= currentKeys[sdl.SCANCODE_C] == 1
	Keys[KPC]	= currentKeys[sdl.SCANCODE_4] == 1
	Keys[KPD]	= currentKeys[sdl.SCANCODE_R] == 1
	Keys[KPE]	= currentKeys[sdl.SCANCODE_F] == 1
	Keys[KPF]	= currentKeys[sdl.SCANCODE_V] == 1
    }
    return pressed
}
