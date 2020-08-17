package chip8

import (
    "testing"
    "fmt"
)

var cpu CPU = NewCPU()

func TestFetch(t *testing.T) {
    cpu.Write(0x0000, 0x3B)
    cpu.Write(0x0001, 0x45)
    cpu.Fetch()
    if cpu.opcode != 0x3B45 {
	t.Errorf("Expected 0x3B45, received %s", fmt.Sprintf("%x", cpu.opcode))
    }
    if cpu.PC != 2 {
	t.Error("Expected PC to be 2")
    }
}

func TestDecode(t *testing.T) {
    // TODO
}

func TestCpuZero(t *testing.T) {
    // TODO
}

func TestCpuTwo(t *testing.T) {
    // 2nnn: call subroutine at nnn
    cpu.opcode = 0x2F47
    curSP := cpu.sp
    //curPC := cpu.PC
    cpu.cpuTwo()
    if cpu.sp != curSP + 1 {
	t.Errorf("Current stack pointer: %s; expected: %s", fmt.Sprintf("%x", cpu.sp), fmt.Sprintf("%x", curSP + 1))
    }
    if cpu.PC != 0x0F47 {
	t.Errorf("PC: %s, expected 0x0F47", fmt.Sprintf("%x", cpu.PC))
    }
}

func TestCpuThree(t *testing.T) {
    // 3xkk: Skip next instruction if Vx = kk
    var tests = []struct {
	opcode	    uint16
	vx	    byte
	pc	    uint16
	expected    uint16
    }{
	{0x3456, 0x56, 0x1000, 0x1002},
	{0x3456, 0x04, 0x1000, 0x1000},
    }

    for _, test := range tests {
	cpu.opcode = test.opcode
	cpu.V[(cpu.opcode & 0x0F00) >> 8] = test.vx
	cpu.PC = test.pc
	cpu.cpuThree()
	if cpu.PC != test.expected {
	    t.Error("Test failed: opcode {}, Vx: {}, expected: {}, received: {}", test.opcode, test.vx, test.expected, cpu.PC)
	}
    }
}

func TestCpuFour(t *testing.T) {
    // 4xkk: skip next instruction if Vx != kk
    var tests = []struct {
	opcode	    uint16
	vx	    byte
	pc	    uint16
	expected    uint16
    }{
	{0x4456, 0x56, 0x1000, 0x1000},
	{0x4456, 0x04, 0x1000, 0x1002},
    }

    for _, test := range tests {
	cpu.opcode = test.opcode
	cpu.V[(cpu.opcode & 0x0F00) >> 8] = test.vx
	cpu.PC = test.pc
	cpu.cpuThree()
	if cpu.PC != test.expected {
	    t.Error("Test failed: opcode {}, Vx: {}, expected: {}, received: {}", test.opcode, test.vx, test.expected, cpu.PC)
	}
    }
}

func TestCpuFive(t *testing.T) {
    // 5xy0: Skip next instruction if Vx = Vy
    var tests = []struct {
	opcode	    uint16
	vx	    byte
	vy	    byte
	pc	    uint16
	expected    uint16
    }{
	{0x5450, 0x56, 0x56, 0x1000, 0x1002},
	{0x5450, 0x06, 0x56, 0x1000, 0x1000},
    }

    for _, test := range tests {
	cpu.opcode = test.opcode
	cpu.V[(cpu.opcode & 0x0F00) >> 8] = test.vx
	cpu.V[(cpu.opcode & 0x00F0) >> 4] = test.vy
	cpu.PC = test.pc
	cpu.cpuThree()
	if cpu.PC != test.expected {
	    t.Error("Test failed: opcode {}, Vx: {}, expected: {}, received: {}", test.opcode, test.vx, test.expected, cpu.PC)
	}
    }
}

func TestCpuSix(t *testing.T) {
    // 6xkk: LD Vx, byte
    var tests = []struct {
	opcode	    uint16
	expected    byte
    }{
	{0x6450, 0x50},
	{0x6B76, 0x76},
    }

    for _, test := range tests {
	cpu.opcode = test.opcode
	x := cpu.opcode & 0x0F00 >> 8
	cpu.cpuSix()
	if cpu.V[x] != test.expected {
	    t.Error("Test failed: opcode {}, expected: {}, received: {}", test.opcode, test.expected, cpu.V[x])
	}
    }
}

func TestCpuSeven(t *testing.T) {
    // 7xkk: ADD Vx, byte
    var tests = []struct {
	opcode	    uint16
	vx	    byte
	kk	    byte
	expected    byte
    }{
	{0x7450, 0x56, 0x10, 0x66},
	{0x7450, 0x06, 0xB4, 0xBA},
    }

    for _, test := range tests {
	cpu.opcode = test.opcode
	x := (cpu.opcode & 0x0F00) >> 8
	cpu.V[x] = test.vx
	cpu.cpuSeven()
	if cpu.V[x] != test.expected {
	    t.Error("Test failed: opcode {}, Vx: {}, kk: {}, expected: {}, received: {}", test.opcode, test.vx, test.kk, test.expected, cpu.V[x])
	}
    }
}

func TestCpuEightZero(t *testing.T) {
    // 6xkk: LD Vx, byte
    var tests = []struct {
	opcode	    uint16
	vx	    byte
	vy	    byte
	expected    byte
    }{
	{0x7450, 0x12, 0x34, 0x34},
	{0x7B76, 0x76, 0x45, 0x45},
    }

    for _, test := range tests {
	cpu.opcode = test.opcode
	x := cpu.opcode & 0x0F00 >> 8
	y := cpu.opcode & 0x00F0 >> 4
	cpu.V[x] = test.vx
	cpu.V[y] = test.vy
	cpu.cpuEightZero()
	if cpu.V[x] != test.expected {
	    t.Error("Test failed: opcode {}, Vx: {}, Vy: {}, expected: {}, received: {}", test.opcode, test.vx, test.vy, test.expected, cpu.V[x])
	}
    }
}
