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
	    t.Errorf("Test failed: opcode %x, Vx: %x, expected: %x, received: %x", test.opcode, test.vx, test.expected, cpu.PC)
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
	cpu.cpuFour()
	if cpu.PC != test.expected {
	    t.Errorf("Test failed: opcode %x, Vx: %x, expected: %x, received: %x", test.opcode, test.vx, test.expected, cpu.PC)
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
	cpu.cpuFive()
	if cpu.PC != test.expected {
	    t.Errorf("Test failed: opcode %x, Vx: %x, expected: %x, received: %x", test.opcode, test.vx, test.expected, cpu.PC)
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
	    t.Errorf("Test failed: opcode %x, expected: %x, received: %x", test.opcode, test.expected, cpu.V[x])
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
	vf	    byte
    }{
	{0x7410, 0x56, 0x10, 0x66, 0},
	// {0x7410, 0xF0, 0x10, 0x00, 1},
	{0x74B4, 0x06, 0xB4, 0xBA, 0},
    }

    for _, test := range tests {
	cpu.opcode = test.opcode
	x := (cpu.opcode & 0x0F00) >> 8
	cpu.V[x] = test.vx
	cpu.cpuSeven()
	if cpu.V[x] != test.expected {
	    t.Errorf("Test failed: opcode %x, Vx: %x, kk: %x, expected: %x, received: %x", test.opcode, test.vx, test.kk, test.expected, cpu.V[x])
	}
	// if cpu.V[15] != test.vf {
	//     t.Errorf("Test failed: VF %x, expected %x", cpu.V[15], test.vf)
	// }
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
	{0x8450, 0x12, 0x34, 0x34},
	{0x8B70, 0x76, 0x45, 0x45},
    }

    for _, test := range tests {
	cpu.opcode = test.opcode
	x := cpu.opcode & 0x0F00 >> 8
	y := cpu.opcode & 0x00F0 >> 4
	cpu.V[x] = test.vx
	cpu.V[y] = test.vy
	cpu.cpuEightZero()
	if cpu.V[x] != test.expected {
	    t.Errorf("Test failed: opcode %x, Vx: %x, Vy: %x, expected: %x, received: %x", test.opcode, test.vx, test.vy, test.expected, cpu.V[x])
	}
    }
}

func TestCpuEightOne(t *testing.T) {
    // 8xy1: OR Vx, Vy -set Vx = Vx OR Vy
    var tests = []struct {
	opcode	    uint16
	vx	    byte
	vy	    byte
	expected    byte
    }{
	{0x8451, 0x12, 0x34, 0x34},
	{0x8B71, 0x76, 0x45, 0x45},
    }

    for _, test := range tests {
	cpu.opcode = test.opcode
	x := cpu.opcode & 0x0F00 >> 8
	y := cpu.opcode & 0x00F0 >> 4
	cpu.V[x] = test.vx
	cpu.V[y] = test.vy
	cpu.cpuEightZero()
	if cpu.V[x] != test.expected {
	    t.Errorf("Test failed: opcode %x, Vx: %x, Vy: %x, expected: %x, received: %x", test.opcode, test.vx, test.vy, test.expected, cpu.V[x])
	}
    }
}

func TestCpuEightTwo(t *testing.T) {
    // 8xy2: AND V, Vy
    var tests = []struct {
	opcode	    uint16
	vx	    byte
	vy	    byte
	expected    byte
    }{
	{0x8452, 0x15, 0x34, 0x14},
	{0x8B72, 0xB6, 0x57, 0x16},
    }

    for _, test := range tests {
	cpu.opcode = test.opcode
	x := cpu.opcode & 0x0F00 >> 8
	y := cpu.opcode & 0x00F0 >> 4
	cpu.V[x] = test.vx
	cpu.V[y] = test.vy
	cpu.cpuEightTwo()
	if cpu.V[x] != test.expected {
	    t.Errorf("Test failed: Vx: %x, Vy: %x, expected: %x, received: %x", test.vx, test.vy, test.expected, cpu.V[x])
	}
    }
}

func TestCpuEightThree(t *testing.T) {
    // 8xy3: XOR Vx, Vy
    var tests = []struct {
	opcode	    uint16
	vx	    byte
	vy	    byte
	expected    byte
    }{
	{0x8453, 0x15, 0x34, 0x21},
	{0x8B73, 0xB6, 0x57, 0xE1},
    }

    for _, test := range tests {
	cpu.opcode = test.opcode
	x := cpu.opcode & 0x0F00 >> 8
	y := cpu.opcode & 0x00F0 >> 4
	cpu.V[x] = test.vx
	cpu.V[y] = test.vy
	cpu.cpuEightThree()
	if cpu.V[x] != test.expected {
	    t.Errorf("Test failed: Vx: %x, Vy: %x, expected: %x, received: %x", test.vx, test.vy, test.expected, cpu.V[x])
	}
    }
}

func TestCpuEightFour(t *testing.T) {
    // 8xy4: ADD Vx, Vy
    var tests = []struct {
	opcode	    uint16
	vx	    byte
	vy	    byte
	expected    byte
	vf	    byte
    }{
	{0x8454, 0x15, 0x34, 0x49, 0},
	{0x8B74, 0xC4, 0xC0, 0x84, 1},
    }

    for _, test := range tests {
	cpu.opcode = test.opcode
	x := cpu.opcode & 0x0F00 >> 8
	y := cpu.opcode & 0x00F0 >> 4
	cpu.V[x] = test.vx
	cpu.V[y] = test.vy
	cpu.cpuEightFour()
	if cpu.V[x] != test.expected {
	    t.Errorf("Test failed: Vx: %x, Vy: %x, expected: %x, received: %x", test.vx, test.vy, test.expected, cpu.V[x])
	}
	if cpu.V[15] != test.vf {
	    t.Errorf("Test failed: VF: %x, expected %x", cpu.V[15], test.vf)
	}
    }
}


func TestCpuEightFive(t *testing.T) {
    // 8xy5: SUB Vx, Vy
    var tests = []struct {
	opcode	    uint16
	vx	    byte
	vy	    byte
	expected    byte
	vf	    byte
    }{
	{0x8455, 0xC0, 0xC4, 0xFC, 0},
	{0x8B75, 0xC4, 0xC0, 0x04, 1},
    }

    for _, test := range tests {
	cpu.opcode = test.opcode
	x := cpu.opcode & 0x0F00 >> 8
	y := cpu.opcode & 0x00F0 >> 4
	cpu.V[x] = test.vx
	cpu.V[y] = test.vy
	cpu.cpuEightFive()
	if cpu.V[x] != test.expected {
	    t.Errorf("Test failed: Vx: %x, Vy: %x, expected: %x, received: %x", test.vx, test.vy, test.expected, cpu.V[x])
	}
	if cpu.V[15] != test.vf {
	    t.Errorf("Test failed: VF: %x, expected %x", cpu.V[15], test.vf)
	}
    }
}


func TestCpuEightSix(t *testing.T) {
    // 8xy6: SHR Vx
    var tests = []struct {
	opcode	    uint16
	vx	    byte
	expected    byte
	vf	    byte
    }{
	{0x8456, 0xC4, 0x62, 0},
	{0x8B76, 0x31, 0x18, 1},
    }

    for _, test := range tests {
	cpu.opcode = test.opcode
	x := cpu.opcode & 0x0F00 >> 8
	cpu.V[x] = test.vx
	cpu.cpuEightSix()
	if cpu.V[x] != test.expected {
	    t.Errorf("Test failed: Vx: %x, expected: %x", cpu.V[x], test.expected)
	}
	if cpu.V[15] != test.vf {
	    t.Errorf("Test failed: VF: %x, expected %x", cpu.V[15], test.vf)
	}
    }
}

func TestCpuEightSeven(t *testing.T) {
    // 8xy5: SUBN Vx, Vy
    var tests = []struct {
	opcode	    uint16
	vx	    byte
	vy	    byte
	expected    byte
	vf	    byte
    }{
	{0x8457, 0xC4, 0xC0, 0xFC, 0},
	{0x8B77, 0xC0, 0xC4, 0x04, 1},
    }

    for _, test := range tests {
	cpu.opcode = test.opcode
	x := cpu.opcode & 0x0F00 >> 8
	y := cpu.opcode & 0x00F0 >> 4
	cpu.V[x] = test.vx
	cpu.V[y] = test.vy
	cpu.cpuEightSeven()
	if cpu.V[x] != test.expected {
	    t.Errorf("Test failed: Vx: %x, Vy: %x, expected: %x, received: %x", test.vx, test.vy, test.expected, cpu.V[x])
	}
	if cpu.V[15] != test.vf {
	    t.Errorf("Test failed: VF: %x, expected %x", cpu.V[15], test.vf)
	}
    }
}

func TestCpuEightE(t *testing.T) {
    // 8xy6: SHL Vx
    var tests = []struct {
	opcode	    uint16
	vx	    byte
	expected    byte
	vf	    byte
    }{
	{0x845E, 0xC4, 0x88, 1},
	{0x8B7E, 0x18, 0x30, 0},
    }

    for _, test := range tests {
	cpu.opcode = test.opcode
	x := cpu.opcode & 0x0F00 >> 8
	cpu.V[x] = test.vx
	cpu.cpuEightE()
	if cpu.V[x] != test.expected {
	    t.Errorf("Test failed: Vx: %x, expected: %x", cpu.V[x], test.expected)
	}
	if cpu.V[15] != test.vf {
	    t.Errorf("Test failed: VF: %x, expected %x", cpu.V[15], test.vf)
	}
    }
}

func TestCpuNine(t *testing.T) {
    // 9xy0: skip next instruction if Vx != Vy
    var tests = []struct {
	opcode	    uint16
	pc	    uint16
	vx	    byte
	vy	    byte
	expected    uint16
    }{
	{0x9450, 0x1234, 0xC4, 0xC0, 0x1236},
	{0x9B70, 0x1234, 0xC4, 0xC4, 0x1234},
    }

    for _, test := range tests {
	cpu.opcode = test.opcode
	cpu.PC = test.pc
	x := cpu.opcode & 0x0F00 >> 8
	y := cpu.opcode & 0x00F0 >> 4
	cpu.V[x] = test.vx
	cpu.V[y] = test.vy
	cpu.cpuNine()
	if cpu.PC != test.expected {
	    t.Errorf("Test failed: Vx: %x, Vy: %x, expected: %x, received: %x", test.vx, test.vy, test.expected, cpu.PC)
	}
    }
}

func TestCpuA(t *testing.T) {
    // Annn: LD I, nnn
    var tests = []struct {
	opcode	    uint16
	expected    uint16
    }{
	{0xA450, 0x0450},
	{0xAB70, 0x0B70},
    }

    for _, test := range tests {
	cpu.opcode = test.opcode
	cpu.cpuA()
	if cpu.I != test.expected {
	    t.Errorf("Test failed: I: %x, expected: %x", cpu.I, test.expected)
	}
    }
}

func TestCpuB(t *testing.T) {
    // Bnnn: JP V0, nnn
    var tests = []struct {
	opcode	    uint16
	pc	    uint16
	v0	    byte
	expected    uint16
    }{
	{0xB450, 0x1234, 0xC4, 0x0514},
	{0xBB70, 0x1234, 0xC4, 0x0C34},
    }

    for _, test := range tests {
	cpu.opcode = test.opcode
	cpu.PC = test.pc
	cpu.V[0] = test.v0
	cpu.cpuB()
	if cpu.PC != test.expected {
	    t.Errorf("Test failed: opcode: %x, V0: %x, expected: %x, received: %x", test.opcode, test.v0, test.expected, cpu.PC)
	}
    }
}

func TestCpuE(t *testing.T) {
    var tests = []struct {
	opcode	    uint16
	pc	    uint16
	vx	    byte
	key	    bool
	expected    uint16
    }{
	{0xE49E, 0x1234, 0x04, false, 0x1236},
	{0xEB9E, 0x1234, 0x04, true,  0x1238},
	{0xEBA1, 0x1234, 0x04, false, 0x1238},
	{0xEBA1, 0x1234, 0x04, true,  0x1236},
    }

    for _, test := range tests {
	cpu.opcode = test.opcode
	cpu.PC = test.pc
	var x byte = byte(test.opcode & 0x0F00 >> 8)
	cpu.V[x] = test.vx
	keys[test.vx] = test.key
	cpu.cpuE()
	if cpu.PC != test.expected {
	    t.Errorf("Test failed: opcode: %x, Vx: %x, key: %t, expected: %x, received: %x", test.opcode, test.vx, test.key, test.expected, cpu.PC)
	}
    }
}

func TestCpuF(t *testing.T) {
    // TODO
    // var tests = []struct {
    //     opcode	    uint16
    //     expected    uint16
    // }
}
