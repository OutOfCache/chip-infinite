package chip8

import (
	"math/rand"
	"time"
)

// CLOCKSPEED defines the 60Hz at which the timers (sound and
// delay) are count down.
const CLOCKSPEED = time.Second / 60 // timers count down at 60Hz

// CPU contains all the registers, memory, the stack, timers,
// opcodes and the drawflag
type CPU struct {
	V          [16]byte   // registers V0-VF
	I          uint16     // registers I
	PC         uint16     // program counter
	Memory                // memory
	stack      [16]uint16 // stack
	sp         byte       // stack pointer
	opcode     uint16     // current opcode
	Delaytimer byte       // delay timer
	Soundtimer byte       // sound timer
	ops        [16]func() // possible opcodes
	opEight    [16]func() // opcodes for 0x8xyz
	Drawflag   bool       // shows whenever sprites have been drawn to the display
}

// NewCPU initializes a new CPU with most fields except the program counter
// initialized to 0
// the following code is mostly derived from github.com/fogleman/nes
func NewCPU() *CPU {
	cpu := &CPU{}
	cpu.PC = 0x200 // starting point for programs
	cpu.I = 0
	cpu.opcode = 0
	cpu.sp = 0

	for i := range cpu.stack {
		cpu.stack[i] = 0
	}

	// load fontset
	for i, font := range fontset {
		cpu.Write(uint16(i), font)
	}

	cpu.Delay_timer = 0xFF
	cpu.Sound_timer = 0xFF

	cpu.opEight = [16]func(){
		cpu.cpuEightZero, cpu.cpuEightOne, cpu.cpuEightTwo, cpu.cpuEightThree,
		cpu.cpuEightFour, cpu.cpuEightFive, cpu.cpuEightSix, cpu.cpuEightSeven,
		cpu.cpuNull, cpu.cpuNull, cpu.cpuNull, cpu.cpuNull,
		cpu.cpuNull, cpu.cpuNull, cpu.cpuEightE, cpu.cpuNull,
	}
	cpu.ops = [16]func(){
		cpu.cpuZero, cpu.cpuOne, cpu.cpuTwo, cpu.cpuThree,
		cpu.cpuFour, cpu.cpuFive, cpu.cpuSix, cpu.cpuSeven,
		cpu.cpuEight, cpu.cpuNine, cpu.cpuA, cpu.cpuB,
		cpu.cpuC, cpu.cpuD, cpu.cpuE, cpu.cpuF,
	}
	return cpu
}

// Fetch sets the drawflag to 0 and fetches the next opcode
func (cpu *CPU) Fetch() {
	cpu.Drawflag = false
	cpu.opcode = uint16(cpu.Read(cpu.PC))<<8 + uint16(cpu.Read(cpu.PC+1))
}

// Decode jumps to the right function in opcodes and executes that function
func (cpu *CPU) Decode() {
	cpu.ops[(cpu.opcode >> 12)]()
}

// instructions
// TODO: add variables of x and y so I can write cpu.V[x] instead of cpu.V[(cpu.opcode ...]
// TODO: check for invalid opcodes
func (cpu *CPU) cpuZero() {
	switch cpu.opcode & 0x00FF {
	case 0x00E0: // Clears the screen
		for i := range Display {
			Display[i] = 0
		}
		cpu.Drawflag = true
		cpu.PC += 2
	case 0x00EE: // Returns from subroutine
		cpu.PC = cpu.stack[cpu.sp] + 2
		cpu.sp--

	default: // 0nnn: SYS addr, apparently omitted in most emulators
	}
}

func (cpu *CPU) cpuOne() {
	// 1nnn: Jump to nnn
	cpu.PC = cpu.opcode & 0x0FFF
}

func (cpu *CPU) cpuTwo() {
	// 2nnn: call subroutine at nnn
	cpu.sp++
	cpu.stack[cpu.sp] = cpu.PC
	cpu.PC = cpu.opcode & 0x0FFF
}

func (cpu *CPU) cpuThree() {
	// 3xkk: Skip next instruction if Vx = kk
	var kk byte = byte(cpu.opcode & 0x00FF)
	if cpu.V[(cpu.opcode&0x0F00)>>8] == kk {
		cpu.PC += 4
	} else {
		cpu.PC += 2
	}
}

func (cpu *CPU) cpuFour() {
	// 4xkk: Skip next instruction if Vx != kk
	var kk byte = byte(cpu.opcode & 0x00FF)
	if cpu.V[(cpu.opcode&0x0F00)>>8] != kk {
		cpu.PC += 4
	} else {
		cpu.PC += 2
	}
}

func (cpu *CPU) cpuFive() {
	// 5xy0: Skip next instruction if Vx = Vy
	if cpu.V[(cpu.opcode&0x0F00)>>8] == cpu.V[(cpu.opcode&0x00F0)>>4] {
		cpu.PC += 4
	} else {
		cpu.PC += 2
	}
}

func (cpu *CPU) cpuSix() {
	// 6xkk: LD Vx, byte - set Vx = byte
	var kk byte = byte(cpu.opcode & 0x00FF)
	cpu.V[(cpu.opcode&0x0F00)>>8] = kk
	cpu.PC += 2
}

func (cpu *CPU) cpuSeven() {
	// 7xkk: ADD Vx, byte - set Vx = Vx + kk
	var kk byte = byte(cpu.opcode & 0x00FF)
	cpu.V[(cpu.opcode&0x0F00)>>8] += kk
	cpu.PC += 2
}

func (cpu *CPU) cpuEight() {
	cpu.opEight[cpu.opcode&0x000F]()
}

func (cpu *CPU) cpuEightZero() {
	// 8xy0: LD Vx, Vy - Set Vx = Vy
	cpu.V[(cpu.opcode&0x0F00)>>8] = cpu.V[(cpu.opcode&0x00F0)>>4]
	cpu.PC += 2
}

func (cpu *CPU) cpuEightOne() {
	// 8xy1: OR Vx, Vy - set Vx = Vx OR Vy
	var result byte = cpu.V[(cpu.opcode&0x0F00)>>8] | cpu.V[(cpu.opcode&0x00F0>>4)]
	cpu.V[(cpu.opcode&0x0F00)>>8] = result
	cpu.PC += 2
}

func (cpu *CPU) cpuEightTwo() {
	// 8xy2: AND Vx, Vy - set Vx = Vx AND Vy
	var result byte = cpu.V[(cpu.opcode&0x0F00)>>8] & cpu.V[(cpu.opcode&0x00F0>>4)]
	cpu.V[(cpu.opcode&0x0F00)>>8] = result
	cpu.PC += 2
}

func (cpu *CPU) cpuEightThree() {
	// 8xy3: XOR Vx, Vy - set Vx = Vx XOR Vy
	var result byte = cpu.V[(cpu.opcode&0x0F00)>>8] ^ cpu.V[(cpu.opcode&0x00F0>>4)]
	cpu.V[(cpu.opcode&0x0F00)>>8] = result
	cpu.PC += 2
}

func (cpu *CPU) cpuEightFour() {
	// 8xy4: ADD Vx, Vy - set Vx = Vx + Vy; VF = carry
	var result uint16 = uint16(cpu.V[(cpu.opcode&0x0F00)>>8]) + uint16(cpu.V[(cpu.opcode&0x00F0>>4)])
	if result > 0x00FF {
		cpu.V[15] = 1
	} else {
		cpu.V[15] = 0
	}
	cpu.V[(cpu.opcode&0x0F00)>>8] = byte(result & 0x00FF)
	cpu.PC += 2
}

func (cpu *CPU) cpuEightFive() {
	// 8xy5: SUB Vx, Vy - set Vx = Vx - Vy; VF = NOT borrow
	// Is this using - or + with two's complement?
	var x byte = byte((cpu.opcode & 0x0F00) >> 8)
	var y byte = byte((cpu.opcode & 0x00F0) >> 4)

	if cpu.V[x] > cpu.V[y] {
		cpu.V[0xF] = 1
	} else {
		cpu.V[0xF] = 0
	}
	var result byte = cpu.V[x] - cpu.V[y]
	cpu.V[x] = result
	cpu.PC += 2
}

func (cpu *CPU) cpuEightSix() {
	// 8xy6: SHR Vx
	var x byte = byte(cpu.opcode & 0x0F00 >> 8)
	cpu.V[15] = cpu.V[x] & 0x0001
	cpu.V[x] = cpu.V[x] >> 1
	cpu.PC += 2
}

func (cpu *CPU) cpuEightSeven() {
	// 8xy7: SUBN Vx, Vy - set Vx = Vy - Vx; VF = NOT borrow
	// Is this using - or + with two's complement?
	var x byte = byte((cpu.opcode & 0x0F00) >> 8)
	var y byte = byte((cpu.opcode & 0x00F0) >> 4)
	if cpu.V[y] > cpu.V[x] {
		cpu.V[15] = 1
	} else {
		cpu.V[15] = 0
	}
	var result byte = cpu.V[y] - cpu.V[x]
	cpu.V[x] = result
	cpu.PC += 2
}

func (cpu *CPU) cpuEightE() {
	// 8xyE: SHL Vx
	var x byte = byte(cpu.opcode & 0x0F00 >> 8)
	cpu.V[15] = cpu.V[x] & 0x80 >> 7
	cpu.V[x] = cpu.V[x] << 1
	cpu.PC += 2
}

func (cpu *CPU) cpuNine() {
	// 9xy0: Skip next instruction if Vx != Vy
	if cpu.V[(cpu.opcode&0x0F00)>>8] != cpu.V[(cpu.opcode&0x00F0)>>4] {
		cpu.PC += 4
	} else {
		cpu.PC += 2
	}
}

func (cpu *CPU) cpuA() {
	// Annn: LD I, addr - set I = nnn
	cpu.I = cpu.opcode & 0x0FFF
	cpu.PC += 2
}

func (cpu *CPU) cpuB() {
	// Bnnn: JP V0, addr - Jump to location nnn + V0
	cpu.PC = (cpu.opcode & 0x0FFF) + uint16(cpu.V[0])
}

func (cpu *CPU) cpuC() {
	// Cxkk: RND Vx, byte - Vx = gen rand numb 0-255 & kk
	rand.Seed(time.Now().UnixNano())
	var rand byte = byte(rand.Intn(255))
	var kk byte = byte(cpu.opcode & 0x00FF)
	cpu.V[(cpu.opcode&0x0F00)>>8] = rand & kk
	cpu.PC += 2
}

func (cpu *CPU) cpuD() {
	// Dxyn: DRW Vx, Vy, nibble - display n-byte sprite starting at mem location I
	// at (Vx, Vy), set VF = collision
	var x byte = byte((cpu.opcode & 0x0F00) >> 8)
	var y byte = byte((cpu.opcode & 0x00F0) >> 4)
	var n byte = byte((cpu.opcode & 0x000F))

	cpu.V[15] = 0

	for h := 0; h < int(n); h++ {
		var curbyte byte = cpu.Read(cpu.I + uint16(h))
		for w := 0; w < 8; w++ {
			if curbyte&(0x80>>w) != 0 {
				var pos = (int(cpu.V[y])+h)%32*64 + ((int(cpu.V[x]) + w) % 64)
				if Display[pos] != 0 {
					cpu.V[15] = 1
				}
				Display[pos] ^= 1
			}
		}
	}
	cpu.Drawflag = true
	cpu.PC += 2
}

func (cpu *CPU) cpuE() {
	var x byte = byte(cpu.opcode & 0x0F00 >> 8)
	switch cpu.opcode & 0x00FF {
	case 0x9E:
		// Ex9E: skip next instruction if key with the value of Vx is pressed
		if Keys[cpu.V[x]] {
			cpu.PC += 4
		} else {
			cpu.PC += 2
		}
		break
	// ExA1: skip next instruction if key with the value of Vx is not pressed
	case 0xA1:
		if !Keys[cpu.V[x]] {
			cpu.PC += 4
		} else {
			cpu.PC += 2
		}
		break
	}
}

func (cpu *CPU) cpuF() {
	switch cpu.opcode & 0x00FF {
	case 0x07: // Fx07: LD Vx, DT - set Vx = delay timer value
		cpu.V[(cpu.opcode&0x0F00)>>8] = cpu.Delay_timer
	case 0x0A: // Fx0A: LD Vx, K - wait for a key press, store the key in Vx
		var x byte = byte(cpu.opcode & 0x0F00 >> 8)
		var key int16
		for key := waitForKey(); key == -1; key = waitForKey() {
			// wait
		}
		cpu.V[x] = byte(key)

	case 0x15: // Fx15: Ld DT, Vx - set delay timer = Vx
		cpu.Delay_timer = cpu.V[(cpu.opcode&0x0F00)>>8]
	case 0x18: // Fx18: LD ST, Vx - set sound timer = Vx
		cpu.Sound_timer = cpu.V[(cpu.opcode&0x0F00)>>8]
	case 0x1E: // Fx1E: ADD I, Vx - I = I + Vx
		cpu.I += uint16(cpu.V[(cpu.opcode&0x0F00)>>8])
	case 0x29: // Fx29: LD F, Vx - set I = location of sprite for digit Vx
		var fontStart uint16 = uint16(0x000)
		var x byte = byte(cpu.opcode & 0x0F00 >> 8)
		var digit uint16 = uint16(cpu.V[x])
		cpu.I = fontStart + (5 * digit) // every digit is 5 byte long
	case 0x33: // Fx33: LD B, Vx - store BCD of Vx in memory locations I, I+1, I+2
		var x byte = byte(cpu.opcode & 0x0F00 >> 8)
		cpu.Write(cpu.I, cpu.V[x]/100)
		cpu.Write(cpu.I+1, (cpu.V[x]/10)%10)
		cpu.Write(cpu.I+2, (cpu.V[x] % 10))
	case 0x55: // Fx55: LD [I], Vx - store registers V0 to Vx in memory
		//	starting at the address in I
		var x byte = byte((cpu.opcode & 0x0F00) >> 8)
		for i := uint16(0); i <= uint16(x); i++ {
			cpu.Write(cpu.I+i, cpu.V[i])
		}
	case 0x65: // Fx65: LD Vx, [I] - read registers V0 through Vx from memory
		//	starting at location I
		var x byte = byte((cpu.opcode & 0x0F00) >> 8)
		for i := byte(0); i <= x; i++ {
			cpu.V[i] = cpu.Read(cpu.I + uint16(i))
		}
	}
	cpu.PC += 2
}

func (cpu *CPU) cpuNull() {
	// do nothing
}
