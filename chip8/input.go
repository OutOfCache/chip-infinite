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

func HandleInput() {
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
		break
	    case sdl.K_1:
		keys[KPOne] = true
		break
	    case sdl.K_2:
		keys[KPTwo] = true
		break
	    case sdl.K_3:
		keys[KPThree] = true
		break
	    case sdl.K_q:
		keys[KPFour] = true
		break
	    case sdl.K_w:
		keys[KPFive] = true
		break
	    case sdl.K_e:
		keys[KPSix] = true
		break
	    case sdl.K_a:
		keys[KPSeven] = true
		break
	    case sdl.K_s:
		keys[KPEight] = true
		break
	    case sdl.K_d:
		keys[KPNine] = true
		break
	    case sdl.K_z:
		keys[KPA] = true
		break
	    case sdl.K_c:
		keys[KPB] = true
		break
	    case sdl.K_4:
		keys[KPC] = true
		break
	    case sdl.K_r:
		keys[KPD] = true
		break
	    case sdl.K_f:
		keys[KPE] = true
		break
	    case sdl.K_v:
		keys[KPF] = true
		break
	    }
	}
    }
}
