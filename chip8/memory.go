package chip8

type Memory struct {
    Memory [4096]byte
}

func (mem *Memory) Read(adr uint16) byte {
    return mem.Memory[adr]
}

func (mem *Memory) Write(adr uint16, val byte) {
    mem.Memory[adr] = val
}
