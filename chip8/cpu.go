package chip8

import (
    "fmt"
)

type CPU struct {
    // registers
    var V [16]byte
    var I uint16
    var PC uint16
    // memory
    Memory
    // instructions?
    var opcode uint16
    ops :=  [16]func(){
	cpuZero, cpuOne, cpuTwo, cpuThree, cpuFour, cpuFive, cpuSix, cpuSeven,
	cpuEight[(opcode & 0x000F)], cpuNine, cpuA, cpuB, cpuC, cpuD, cpuE, cpuF
    }

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
func (cpu *CPU) cpuZero() {
}

func (cpu *CPU) cpuOne() {
}

func (cpu *CPU) cpuTwo() {
}

func (cpu *CPU) cpuThree() {
}

func (cpu *CPU) cpuFour() {
}

func (cpu *CPU) cpuFive() {
}

func (cpu *CPU) cpuSix() {
}

func (cpu *CPU) cpuSeven() {
}

var cpuEight = [16]func() {
    cpuEightZero, cpuEightOne, cpuEightTwo, cpuEightThree,
    cpuEightFour, cpuEightFive, cpuEightSix, cpuEightSeven,
    cpuNull, cpuNull, cpuNull, cpuNull,
    cpuNull, cpuNull, cpuEightE, cpuNull
}

func (cpu *CPU) cpuEightZero() {
}

func (cpu *CPU) cpuEightOne() {
}

func (cpu *CPU) cpuEightTwo() {
}

func (cpu *CPU) cpuEightThree() {
}

func (cpu *CPU) cpuEightFour() {
}

func (cpu *CPU) cpuEightFive() {
}

func (cpu *CPU) cpuEightSix() {
}

func (cpu *CPU) cpuEightSeven() {
}

func (cpu *CPU) cpuEightE() {
}

func (cpu *CPU) cpuNine() {
}

func (cpu *CPU) cpuA() {
}

func (cpu *CPU) cpuB() {
}

func (cpu *CPU) cpuC() {
}

func (cpu *CPU) cpuD() {
}

func (cpu *CPU) cpuE() {
}

func (cpu *CPU) cpuF() {
}

func (cpu *CPU) cpuNull() {
}
