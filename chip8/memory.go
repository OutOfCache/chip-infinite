package chip8

type Memory struct {
    var mem [4096]byte
}

func (mem Memory) read(adr uint16) byte {
    return mem[adr]
}

func (mem Memory) write(adr uint16, val byte) {
    mem[adr] = val
}
