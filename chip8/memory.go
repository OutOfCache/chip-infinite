package chip8

type Memory struct {
    var mem [4096]byte
}

func (mem Memory) Read(adr uint16) byte {
    return mem[adr]
}

func (mem Memory) Write(adr uint16, val byte) {
    mem[adr] = val
}
