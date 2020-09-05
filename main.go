package main

import (
	"./chip8"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

var cpu *chip8.CPU

func initialize() *chip8.CPU {
	cpu = chip8.NewCPU()

	for i := range chip8.Display {
		chip8.Display[i] = 0
	}
	return cpu
}

func loadProgram() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go path/to/rom")
	}
	var path string = os.Args[1]

	rom, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("File could not be read")
		fmt.Println(err)
		os.Exit(2)
	}

	for i, buffer := range rom {
		cpu.Write(uint16(0x200+i), buffer)
	}
}

func main() {
	// set up
	cpu = initialize()
	loadProgram()

	now := time.Now()
	chip8.Quit = false
	if !chip8.StartSDL() {
		fmt.Println("Failed to initialize")
	} else {
		chip8.InitAudio()
		for !chip8.Quit {
			cpu.Fetch()
			cpu.Decode()

			if time.Since(now) >= chip8.CLOCKSPEED {
				now = time.Now()
				if cpu.Delay_timer > 0 {
					cpu.Delay_timer--
				}

				if cpu.Sound_timer > 0 {
					if cpu.Sound_timer == 1 {
						chip8.PlayBeep()
					}
					cpu.Sound_timer--
				}
			}

			chip8.HandleInput()
			if cpu.Drawflag {
				chip8.Render()
			} else {
				time.Sleep(time.Microsecond * 900) // reduce CPU usage and speed of emulation
			}
		}
	}
	chip8.End()
}
