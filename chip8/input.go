package chip8

import(
    "github.com/veandco/go-sdl2/sdl"
)

var keys [16]bool

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
    for i, _ := range keys {
	keys[i] = false
    }

    for e:= sdl.PollEvent(); e != nil; e = sdl.PollEvent() {
	switch t := e.(type) {
	case *sdl.QuitEvent:
	    Quit = true
	    break
	case *sdl.KeyboardEvent:
	    keyCode := t.Keysym.Sym

	    switch(keyCode) {
	    case sdl.K_x:
		keys[KPZero] = true
		pressed = KPZero
		break
	    case sdl.K_1:
		keys[KPOne] = true
		pressed = KPOne
		break
	    case sdl.K_2:
		keys[KPTwo] = true
		pressed = KPTwo
		break
	    case sdl.K_3:
		keys[KPThree] = true
		pressed = KPThree
		break
	    case sdl.K_q:
		keys[KPFour] = true
		pressed = KPFour
		break
	    case sdl.K_w:
		keys[KPFive] = true
		pressed = KPFive
		break
	    case sdl.K_e:
		keys[KPSix] = true
		pressed = KPSix
		break
	    case sdl.K_a:
		keys[KPSeven] = true
		pressed = KPSeven
		break
	    case sdl.K_s:
		keys[KPEight] = true
		pressed = KPEight
		break
	    case sdl.K_d:
		keys[KPNine] = true
		pressed = KPNine
		break
	    case sdl.K_z:
		keys[KPA] = true
		pressed = KPA
		break
	    case sdl.K_c:
		keys[KPB] = true
		pressed = KPB
		break
	    case sdl.K_4:
		keys[KPC] = true
		pressed = KPC
		break
	    case sdl.K_r:
		keys[KPD] = true
		pressed = KPD
		break
	    case sdl.K_f:
		keys[KPE] = true
		pressed = KPE
		break
	    case sdl.K_v:
		keys[KPF] = true
		pressed = KPF
		break
	    }
	}
    }
    return pressed
}
