package chip8

import (
    "fmt"
    "math/rand"
)

type CPU struct {
    // registers
    var V [16]byte
    var I uint16
    var PC uint16
    // memory
    Memory
    // stack
    var stack [16]uint16
    var sp byte
    // instructions?
    var opcode uint16
    ops :=  [16]func(){
	cpuZero, cpuOne, cpuTwo, cpuThree, cpuFour, cpuFive, cpuSix, cpuSeven,
	cpuEight[(opcode & 0x000F)], cpuNine, cpuA, cpuB, cpuC, cpuD, cpuE, cpuF
    }

    // timer registers
    var delay_timer byte
    var sound_timer byte

}

// fetch instruction
func (cpu *CPU) Fetch() {
    cpu.opcode = cpu.Read([cpu.PC]) << 8 | cpu.Read([cpu.PC + 1])
    cpu.PC += 2
}
// decode and execute opcode
func (cpu *CPU) Decode() {
    cpu.ops[(cpu.opcode >> 12)]()
}

// instructions
// TODO: add variables of x and y so I can write cpu.V[x] instead of cpu.V[(cpu.opcode ...]
// TODO: check for invalid opcodes
func (cpu *CPU) cpuZero() {
    switch(cpu.opcode & 0x00FF){
    case 0x00E0:
	// Clears the screen
    case 0x00EE:
	// Returns from subroutine
    default:
	// 0nnn: SYS addr
}

func (cpu *CPU) cpuOne() {
    // 1nnn: Jump to nnn
    cpu.PC = cpu.opcode & 0x0FFF
}

func (cpu *CPU) cpuTwo() {
    // 2nnn: call subroutine at nnn
    cpu.stack[cpu.sp] = cpu.PC
    cpu.sp++
    cpu.PC = cpu.opcode & 0x0FFF
}

func (cpu *CPU) cpuThree() {
    // 3xkk: Skip next instruction if Vx = kk
    var kk byte = opcode & 0x00FF
    if cpu.V[(cpu.opcode & 0x0F00) >> 8] == kk {
	cpu.PC += 2
    }
}

func (cpu *CPU) cpuFour() {
    // 4xkk: Skip next instruction if Vx != kk
    var kk byte = opcode & 0x00FF
    if cpu.V[(cpu.opcode & 0x0F00) >> 8] != kk {
	cpu.PC += 2
    }
}

func (cpu *CPU) cpuFive() {
    // 5xy0: Skip next instruction if Vx = Vy
    if cpu.V[(cpu.opcode & 0x0F00) >> 8] == cpu.V[(cpu.opcode & 0x00F0) >> 4] {
	cpu.PC += 2
    }
}

func (cpu *CPU) cpuSix() {
    // 6xkk: LD Vx, byte - set Vx = byte
    var kk byte = cpu.opcode & 0x00FF
    cpu.V[(cpu.opcode & 0x0F00) >> 8] = kk
}

func (cpu *CPU) cpuSeven() {
    // 7xkk: ADD Vx, byte - set Vx = Vx + kk
    var kk byte = cpu.opcode & 0x00FF
    cpu.V[(cpu.opcode & 0x0F00) >> 8] += kk
}

var cpuEight = [16]func() {
    cpuEightZero, cpuEightOne, cpuEightTwo, cpuEightThree,
    cpuEightFour, cpuEightFive, cpuEightSix, cpuEightSeven,
    cpuNull, cpuNull, cpuNull, cpuNull,
    cpuNull, cpuNull, cpuEightE, cpuNull
}

func (cpu *CPU) cpuEightZero() {
    // 8xy0: LD Vx, Vy - Set Vx = Vy
    cpu.V[(cpu.opcode & 0x0F00) >> 8] = cpu.V[(cpu.opcode & 0x00F0) >> 4]
}

func (cpu *CPU) cpuEightOne() {
    // 8xy1: OR Vx, Vy - set Vx = Vx OR Vy
    var result byte = cpu.V[(cpu.opcode & 0x0F00) >> 8] | cpu.V[(cpu.opcode & 0x00F0 >> 4]
    cpu.V[(cpu.opcode & 0x0F00) >> 8] = result
}

func (cpu *CPU) cpuEightTwo() {
    // 8xy2: AND Vx, Vy - set Vx = Vx AND Vy
    var result byte = cpu.V[(cpu.opcode & 0x0F00) >> 8] & cpu.V[(cpu.opcode & 0x00F0 >> 4]
    cpu.V[(cpu.opcode & 0x0F00) >> 8] = result
}

func (cpu *CPU) cpuEightThree() {
    // 8xy3: XOR Vx, Vy - set Vx = Vx XOR Vy
    var result byte = cpu.V[(cpu.opcode & 0x0F00) >> 8] ^ cpu.V[(cpu.opcode & 0x00F0 >> 4]
    cpu.V[(cpu.opcode & 0x0F00) >> 8] = result
}

func (cpu *CPU) cpuEightFour() {
    // 8xy4: ADD Vx, Vy - set Vx = Vx + Vy; VF = carry
    var result uint16 = cpu.V[(cpu.opcode & 0x0F00) >> 8] + cpu.V[(cpu.opcode & 0x00F0 >> 4]
    if result > 0x00FF {
	cpu.V[16] = 1
    } else {
	cpu.V[16] = 0
    }
    cpu.V[(cpu.opcode & 0x0F00) >> 8] = result & 0x00FF
}

func (cpu *CPU) cpuEightFive() {
    // 8xy5: SUB Vx, Vy - set Vx = Vx - Vy; VF = NOT borrow
    // Is this using - or + with two's complement?
    if cpu.V[(cpu.opcode & 0x0F00) >> 8] > cpu.V[(cpu.opcode & 0x00F0) >> 4] {
	cpu.V[16] = 1
    } else {
	cpu.V[16] = 0
    }
    var result byte = cpu.V[(cpu.opcode & 0x0F00) >> 8] - cpu.V[(cpu.opcode & 0x00F0 >> 4]
    cpu.V[(cpu.opcode & 0x0F00) >> 8] = result
}

func (cpu *CPU) cpuEightSix() {
    // 8xy6: SHR Vx
    var x byte = cpu.opcode & 0x0F00 >> 8
    cpu.V[16] = cpu.V[x] & 0x0001
    cpu.V[x] = cpu.V[x] >> 1
}

func (cpu *CPU) cpuEightSeven() {
    // 8xy7: SUBN Vx, Vy - set Vx = Vy - Vx; VF = NOT borrow
    // Is this using - or + with two's complement?
    if cpu.V[(cpu.opcode & 0x00F0) >> 4] > cpu.V[(cpu.opcode & 0x0F00) >> 8] {
	cpu.V[16] = 1
    } else {
	cpu.V[16] = 0
    }
    var result byte = cpu.V[(cpu.opcode & 0x00F0) >> 4] - cpu.V[(cpu.opcode & 0x0F00 >> 8]
    cpu.V[(cpu.opcode & 0x0F00) >> 8] = result
}

func (cpu *CPU) cpuEightE() {
    // 8xyE: SHL Vx
    var x byte = cpu.opcode & 0x0F00 >> 8
    cpu.V[16] = cpu.V[x] & 0x8000
    cpu.V[x] = cpu.V[x] << 1
}

func (cpu *CPU) cpuNine() {
    // 9xy0: Skip next instruction if Vx != Vy
    if cpu.V[(cpu.opcode & 0x0F00) >> 8] != cpu.V[(cpu.opcode & 0x00F0) >> 4] {
	cpu.PC += 2
    }
}

func (cpu *CPU) cpuA() {
    // Annn: LD I, addr - set I = nnn
    cpu.I = cpu.opcode & 0x0FFF
}

func (cpu *CPU) cpuB() {
    // Bnnn: JP V0, addr - Jump to location nnn + V0
    cpu.PC = (cpu.opcode & 0x0FFF) + cpu.V[0]
}

func (cpu *CPU) cpuC() {
    // Cxkk: RND Vx, byte - Vx = gen rand numb 0-255 & kk
    rand.Seed(time.Now().UnixNano())
    var rand byte = rand.Int(255) & 0x000F
    var kk byte = cpu.opcode & 0x00FF
    cpu.V[(cpu.opcode & 0x0F00) >> 8] = rand & kk
}

func (cpu *CPU) cpuD() {
    // Dxyn: DRW Vx, Vy, nibble - display n-byte sprite starting at mem location I
    // at (Vx, Vy), set VF = collision

}

func (cpu *CPU) cpuE() {
    // Ex9E: skip next instruction if key with the value of Vx is pressed
    // ExA1: skip next instruction if key with the value of Vx is not pressed
}

func (cpu *CPU) cpuF() {
    switch cpu.opcode & 0x00FF {
	case 0x07: // Fx07: LD Vx, DT - set Vx = delay timer value
	    cpu.V[(cpu.opcode & 0x0F00) >> 8] = cpu.delay_timer
	case 0x0A: // Fx0A: LD Vx, K - wait for a key press, store the key in Vx
	    // TODO
	case 0x15: // Fx15: Ld DT, Vx - set delay timer = Vx
	    cpu.delay_timer = cpu.V[(cpu.opcode & 0x0F00) >> 8]
	case 0x18: // Fx18: LD ST, Vx - set sound timer = Vx
	    cpu.sound_timer = cpu.V[(cpu.opcode & 0x0F00) >> 8]
	case 0x1E: // Fx1E: ADD I, Vx - I = I + Vx
	    cpu.I += cpu.V[(cpu.opcode & 0x0F00) >> 8] 
	case 0x29: // Fx29: LD F, Vx - set I = location of sprite for digit Vx
	    // TODO
	case 0x33: // Fx33: LD B, Vx - store BCD of Vx in memory locations I, I+1, I+2
	    // TODO
	case 0x55: // Fx55: LD [I], Vx - store registers V0 to Vx in memory 
		//	starting at the address in I
	    var x byte = (cpu.opcode & 0x0F00) >> 8
	    for var i byte = 0; i <= x; i++ {
	        mem.Write(cpu.I + i, cpu.V[i])
	    }
	case 0x65: // Fx65: LD Vx, [I] - read registers V0 through Vx from memory
			//	starting at location I
	    var x byte = (cpu.opcode & 0x0F00) >> 8
	    for var i byte = 0; i <= x; i++ {
	        cpu.V[i] = mem.Read(cpu.I + i)
	    }
    }
}



func (cpu *CPU) cpuNull() {
    // do nothing
}
