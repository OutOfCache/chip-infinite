package chip8

import (
	"testing"
)

var mem Memory

func TestWrite(t *testing.T) {
	var adr uint16 = 0x0016
	mem.Write(adr, 0x08)

	if mem.Memory[0x0016] != 0x08 {
		t.Error("Expected mem[0x0016] to be 0x08")
	}
}

func TestRead(t *testing.T) {
	var adr uint16 = 0x0016
	mem.Memory[adr] = 0x08
	if mem.Read(adr) != 0x08 {
		t.Error("Expected mem[0x0016] to be 0x08")
	}
}
