package chip8

// typedef unsigned char Uint8;
// void SquareWave(void *userdata, Uint8 *stream, int len);
import "C"
import (
	"fmt"
	"math"
	"reflect"
	"unsafe"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	totalTime    = 100 // in ms
	toneHz       = 220
	sampleHz     = 48000
	swPeriod     = sampleHz / toneHz
	halfSwP      = swPeriod / 2
	samplePerMs  = sampleHz / 1000
	totalSamples = samplePerMs * totalTime
)

var currentSample uint32
var device sdl.AudioDeviceID
var spec *sdl.AudioSpec

//export SquareWave
func SquareWave(userdata unsafe.Pointer, stream *C.Uint8, length C.int) {
	n := int(length)
	hdr := reflect.SliceHeader{Data: uintptr(unsafe.Pointer(stream)), Len: n, Cap: n}
	buf := *(*[]C.Uint8)(unsafe.Pointer(&hdr))

	for i := 0; i < n; i++ {
		if currentSample > totalSamples {
			return
		}
		sample := C.Uint8(math.Mod(math.Floor(float64(i)/swPeriod), 2) * 64 /* volume */)
		buf[i] = sample
		currentSample++
	}
}

func InitAudio() {
	spec = &sdl.AudioSpec{
		Freq:     sampleHz,
		Format:   sdl.AUDIO_U8,
		Channels: 2,
		Samples:  totalSamples,
		Callback: sdl.AudioCallback(C.SquareWave),
	}

	device, err = sdl.OpenAudioDevice("", false, spec, nil, sdl.AUDIO_ALLOW_ANY_CHANGE)

	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
}

func PlayBeep() {
	currentSample = 0

	//InitAudio()

	for currentSample < totalSamples {
		sdl.PauseAudioDevice(device, false)
	}

	sdl.PauseAudioDevice(device, true)
	//sdl.CloseAudioDevice(device)

	//if currentSample > totalSamples {
	//	return
	//}
}
