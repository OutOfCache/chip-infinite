// Package chip8 contains the necessary components for emulation of the Chip-8
package chip8

// Memory contains an array that represents the entirety of the 4K RAM.
// A separate struct is most likely not appropriate for a simple structure
// like the chip-8, however I wanted to figure out a working file structure
// for future, more complicated projects
type Memory struct {
	Memory [4096]byte
}

// Read takes an adress and returns the value stored at that adress
func (mem *memory) Read(adr uint16) byte {
	return mem.Memory[adr]
}

// Write takes and adress and a value and stores the value at the given adress
func (mem *Memory) Write(adr uint16, val byte) {
	mem.Memory[adr] = val
}
