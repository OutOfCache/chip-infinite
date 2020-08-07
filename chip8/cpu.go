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
}

// instructions
