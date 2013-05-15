package main

import (
	"fmt"
)

var (
	c  *Cpu
	pc uint16
)

func immediateAddress() int {
	val, _ := Ram.Read(pc)
	return int(val)
}

func absoluteAddress() (result int) {
	// Switch to an int (or more appropriately uint16) since we 
	// will overflow when shifting the high byte
	high, _ := Ram.Read(pc + 1)
	low, _ := Ram.Read(pc)

	return (int(high) << 8) + int(low)
}

func zeroPageAddress() int {
	res, _ := Ram.Read(pc)

	return int(res)
}

func indirectAbsoluteAddress() (result int) {
	high, _ := Ram.Read(pc + 1)
	low, _ := Ram.Read(pc)

	result = (int(high) << 8) + int(low)
	pc++
	return
}

func absoluteIndexedAddress(index Word) (result int) {
	// Switch to an int (or more appropriately uint16) since we 
	// will overflow when shifting the high byte
	high, _ := Ram.Read(pc + 1)
	low, _ := Ram.Read(pc)

	return (int(high) << 8) + int(low)
}

func zeroPageIndexedAddress(index Word) int {
	location, _ := Ram.Read(pc)
	return int(location)
}

func indexedIndirectAddress() int {
	location, _ := Ram.Read(pc)
	location = location + c.X

	// Switch to an int (or more appropriately uint16) since we 
	// will overflow when shifting the high byte
	high, _ := Ram.Read(location + 1)
	low, _ := Ram.Read(location)

	return (int(high) << 8) + int(low)
}

func indirectIndexedAddress() int {
	val, _ := Ram.Read(pc)
	return int(val)
}

func relativeAddress() int {
	val, _ := Ram.Read(pc)
	return int(pc + 1) + int(int8(val))
}

func accumulatorAddress() int {
	return 0
}

func Disassemble(opcode Word, cpu *Cpu, p uint16) {
	c = cpu
	pc = p

	//fmt.Printf("0x%x: 0x%x ", pc-1, opcode)

	switch opcode {
	// ADC
	case 0x69:
		fmt.Printf("adc #$%02x\n", immediateAddress())
	case 0x65:
		fmt.Printf("adc $%02x\n", zeroPageAddress())
	case 0x75:
		fmt.Printf("adc $%02x, X\n", zeroPageIndexedAddress(c.X))
	case 0x6D:
		fmt.Printf("adc $%04x\n", absoluteAddress())
	case 0x7D:
		fmt.Printf("adc $%04x, X\n", absoluteIndexedAddress(c.X))
	case 0x79:
		fmt.Printf("adc $%04x, Y\n", absoluteIndexedAddress(c.Y))
	case 0x61:
		fmt.Printf("adc ($%02x, X)\n", indexedIndirectAddress())
	case 0x71:
		fmt.Printf("adc ($%02x), Y\n", indirectIndexedAddress())
	// LDA
	case 0xA9:
		fmt.Printf("lda #$%02x\n", immediateAddress())
	case 0xA5:
		fmt.Printf("lda $%02x\n", zeroPageAddress())
	case 0xB5:
		fmt.Printf("lda $%02x, X\n", zeroPageIndexedAddress(c.X))
	case 0xAD:
		fmt.Printf("lda $%04x\n", absoluteAddress())
	case 0xBD:
		fmt.Printf("lda $%04x, X\n", absoluteIndexedAddress(c.X))
	case 0xB9:
		fmt.Printf("lda $%04x, Y\n", absoluteIndexedAddress(c.Y))
	case 0xA1:
		fmt.Printf("lda ($%02x, X)\n", indexedIndirectAddress())
	case 0xB1:
		fmt.Printf("lda ($%02x), Y\n", indirectIndexedAddress())
	// LDX
	case 0xA2:
		fmt.Printf("ldx #$%02x\n", immediateAddress())
	case 0xA6:
		fmt.Printf("ldx $%02x\n", zeroPageAddress())
	case 0xB6:
		fmt.Printf("ldx $%02x, X\n", zeroPageIndexedAddress(c.X))
	case 0xAE:
		fmt.Printf("ldx $%04x\n", absoluteAddress())
	case 0xBE:
		fmt.Printf("ldx $%04x, Y\n", absoluteIndexedAddress(c.Y))
	// LDY
	case 0xA0:
		fmt.Printf("ldy #$%02x\n", immediateAddress())
	case 0xA4:
		fmt.Printf("ldy $%02x\n", zeroPageAddress())
	case 0xB4:
		fmt.Printf("ldy $%02x, X\n", zeroPageIndexedAddress(c.X))
	case 0xAC:
		fmt.Printf("ldy $%04x\n", absoluteAddress())
	case 0xBC:
		fmt.Printf("ldy $%04x, X\n", absoluteIndexedAddress(c.X))
	// STA
	case 0x85:
		fmt.Printf("sta $%02x\n", zeroPageAddress())
	case 0x95:
		fmt.Printf("sta $%02x, X\n", zeroPageIndexedAddress(c.X))
	case 0x8D:
		fmt.Printf("sta $%04x\n", absoluteAddress())
	case 0x9D:
		fmt.Printf("sta $%04x, X\n", absoluteIndexedAddress(c.X))
	case 0x99:
		fmt.Printf("sta $%04x, Y\n", absoluteIndexedAddress(c.Y))
	case 0x81:
		fmt.Printf("sta ($%02x, X)\n", indexedIndirectAddress())
	case 0x91:
		fmt.Printf("sta ($%02x), Y\n", indirectIndexedAddress())
	// STX
	case 0x86:
		fmt.Printf("stx $%02x\n", zeroPageAddress())
	case 0x96:
		fmt.Printf("stx $%02x, X\n", zeroPageIndexedAddress(c.X))
	case 0x8E:
		fmt.Printf("stx $%04x\n", absoluteAddress())
	// STY
	case 0x84:
		fmt.Printf("sty $%02x\n", zeroPageAddress())
	case 0x94:
		fmt.Printf("sty $%02x, X\n", zeroPageIndexedAddress(c.X))
	case 0x8C:
		fmt.Printf("sty $%04x\n", absoluteAddress())
	// JMP
	case 0x4C:
		fmt.Printf("jmp $%04x\n", absoluteAddress())
	case 0x6C:
		fmt.Printf("jmp ($%04x)\n", indirectAbsoluteAddress())
	// JSR
	case 0x20:
		fmt.Printf("jsr $%04x\n", absoluteAddress())
	// Register Instructions
	case 0xAA:
		fmt.Println("tax")
	case 0x8A:
		fmt.Println("txa")
	case 0xCA:
		fmt.Println("dex")
	case 0xE8:
		fmt.Println("inx")
	case 0xA8:
		fmt.Println("tay")
	case 0x98:
		fmt.Println("tya")
	case 0x88:
		fmt.Println("dey")
	case 0xC8:
		fmt.Println("iny")
	// Branch Instructions
	case 0x10:
		fmt.Printf("bpl $%04x\n", relativeAddress())
	case 0x30:
		fmt.Printf("bmi $%04x\n", relativeAddress())
	case 0x50:
		fmt.Printf("bvc $%04x\n", relativeAddress())
	case 0x70:
		fmt.Printf("bvs $%04x\n", relativeAddress())
	case 0x90:
		fmt.Printf("bcc $%04x\n", relativeAddress())
	case 0xB0:
		fmt.Printf("bcs $%04x\n", relativeAddress())
	case 0xD0:
		fmt.Printf("bne $%04x\n", relativeAddress())
	case 0xF0:
		fmt.Printf("beq $%04x\n", relativeAddress())
	// CMP
	case 0xC9:
		fmt.Printf("cmp #$%02x\n", immediateAddress())
	case 0xC5:
		fmt.Printf("cmp $%02x\n", zeroPageAddress())
	case 0xD5:
		fmt.Printf("cmp $%02x, X\n", zeroPageIndexedAddress(c.X))
	case 0xCD:
		fmt.Printf("cmp $%04x\n", absoluteAddress())
	case 0xDD:
		fmt.Printf("cmp $%04x, X\n", absoluteIndexedAddress(c.X))
	case 0xD9:
		fmt.Printf("cmp $%04x, Y\n", absoluteIndexedAddress(c.Y))
	case 0xC1:
		fmt.Printf("cmp ($%02x, X)\n", indexedIndirectAddress())
	case 0xD1:
		fmt.Printf("cmp ($%02x), Y\n", c.indirectIndexedAddress())
	// CPX
	case 0xE0:
		fmt.Printf("cpx #$%02x\n", immediateAddress())
	case 0xE4:
		fmt.Printf("cpx $%02x\n", zeroPageAddress())
	case 0xEC:
		fmt.Printf("cpx $%04x\n", absoluteAddress())
	// CPY
	case 0xC0:
		fmt.Printf("cpy #$%02x\n", immediateAddress())
	case 0xC4:
		fmt.Printf("cpy $%02x\n", zeroPageAddress())
	case 0xCC:
		fmt.Printf("cpy $%04x\n", absoluteAddress())
	// SBC
	case 0xE9:
		fmt.Printf("sbc #$%02x\n", immediateAddress())
	case 0xE5:
		fmt.Printf("sbc $%02x\n", zeroPageAddress())
	case 0xF5:
		fmt.Printf("sbc $%02x, X\n", zeroPageIndexedAddress(c.X))
	case 0xED:
		fmt.Printf("sbc $%04x\n", absoluteAddress())
	case 0xFD:
		fmt.Printf("sbc $%04x, X\n", absoluteIndexedAddress(c.X))
	case 0xF9:
		fmt.Printf("sbc $%04x, Y\n", absoluteIndexedAddress(c.Y))
	case 0xE1:
		fmt.Printf("sbc ($%02x, X)\n", indexedIndirectAddress())
	case 0xF1:
		fmt.Printf("sbc ($%02x), Y\n", indirectIndexedAddress())
	// Flag Instructions
	case 0x18:
		fmt.Println("clc")
	case 0x38:
		fmt.Println("sec")
	case 0x58:
		fmt.Println("cli")
	case 0x78:
		fmt.Println("sei")
	case 0xB8:
		fmt.Println("clv")
	case 0xD8:
		fmt.Println("cld")
	case 0xF8:
		fmt.Println("sed")
	// Stack instructions
	case 0x9A:
		fmt.Println("txs")
	case 0xBA:
		fmt.Println("tsx")
	case 0x48:
		fmt.Println("pha")
	case 0x68:
		fmt.Println("pla")
	case 0x08:
		fmt.Println("php")
	case 0x28:
		fmt.Println("plp")
	// AND
	case 0x29:
		fmt.Printf("and #$%02x\n", immediateAddress())
	case 0x25:
		fmt.Printf("and $%02x\n", zeroPageAddress())
	case 0x35:
		fmt.Printf("and $%02x, X\n", zeroPageIndexedAddress(c.X))
	case 0x2d:
		fmt.Printf("and $%04x\n", absoluteAddress())
	case 0x3d:
		fmt.Printf("and $%04x, X\n", absoluteIndexedAddress(c.X))
	case 0x39:
		fmt.Printf("and $%04x, Y\n", absoluteIndexedAddress(c.Y))
	case 0x21:
		fmt.Printf("and ($%02x, X)\n", indexedIndirectAddress())
	case 0x31:
		fmt.Printf("and ($%02x), Y\n", indirectIndexedAddress())
	// ORA
	case 0x09:
		fmt.Printf("ora #$%02x\n", immediateAddress())
	case 0x05:
		fmt.Printf("ora $%02x\n", zeroPageAddress())
	case 0x15:
		fmt.Printf("ora $%02x, X\n", zeroPageIndexedAddress(c.X))
	case 0x0d:
		fmt.Printf("ora $%04x\n", absoluteAddress())
	case 0x1d:
		fmt.Printf("ora $%04x, X\n", absoluteIndexedAddress(c.X))
	case 0x19:
		fmt.Printf("ora $%04x, Y\n", absoluteIndexedAddress(c.Y))
	case 0x01:
		fmt.Printf("ora ($%02x, X)\n", indexedIndirectAddress())
	case 0x11:
		fmt.Printf("ora ($%02x), Y\n", indirectIndexedAddress())
	// EOR
	case 0x49:
		fmt.Printf("eor #$%02x\n", immediateAddress())
	case 0x45:
		fmt.Printf("eor $%02x\n", zeroPageAddress())
	case 0x55:
		fmt.Printf("eor $%02x, X\n", zeroPageIndexedAddress(c.X))
	case 0x4d:
		fmt.Printf("eor $%04x\n", absoluteAddress())
	case 0x5d:
		fmt.Printf("eor $%04x, X\n", absoluteIndexedAddress(c.X))
	case 0x59:
		fmt.Printf("eor $%04x, Y\n", absoluteIndexedAddress(c.Y))
	case 0x41:
		fmt.Printf("eor ($%02x, X)\n", indexedIndirectAddress())
	case 0x51:
		fmt.Printf("eor ($%02x), Y\n", indirectIndexedAddress())
	// dec
	case 0xc6:
		fmt.Printf("dec $%02x\n", zeroPageAddress())
	case 0xd6:
		fmt.Printf("dec $%02x, X\n", zeroPageIndexedAddress(c.X))
	case 0xce:
		fmt.Printf("dec $%04x\n", absoluteAddress())
	case 0xde:
		fmt.Printf("dec $%04x, X\n", absoluteIndexedAddress(c.X))
	// INC
	case 0xe6:
		fmt.Printf("inc $%02x\n", zeroPageAddress())
	case 0xf6:
		fmt.Printf("inc $%02x, X\n", zeroPageIndexedAddress(c.X))
	case 0xee:
		fmt.Printf("inc $%04x\n", absoluteAddress())
	case 0xfe:
		fmt.Printf("inc $%04x, X\n", absoluteIndexedAddress(c.X))
	// BRK
	case 0x00:
		fmt.Println("brk")
	// RTI
	case 0x40:
		fmt.Println("rti")
	// RTS
	case 0x60:
		fmt.Println("rts")
	// NOP
	case 0xea:
		fmt.Println("nop")
	// LSR
	case 0x4a:
		fmt.Println("lsr")
	case 0x46:
		fmt.Printf("lsr $%02x\n", zeroPageAddress())
	case 0x56:
		fmt.Printf("lsr $%02x, X\n", zeroPageIndexedAddress(c.X))
	case 0x4e:
		fmt.Printf("lsr $%04x\n", absoluteAddress())
	case 0x5e:
		fmt.Printf("lsr $%04x, X\n", absoluteIndexedAddress(c.X))
	// ASL
	case 0x0a:
		fmt.Println("asl")
	case 0x06:
		fmt.Printf("asl $%02x\n", zeroPageAddress())
	case 0x16:
		fmt.Printf("asl $%02x, X\n", zeroPageIndexedAddress(c.X))
	case 0x0e:
		fmt.Printf("asl $%04x\n", absoluteAddress())
	case 0x1e:
		fmt.Printf("asl $%04x, X\n", absoluteIndexedAddress(c.X))
	// ROL
	case 0x2a:
		fmt.Println("rol")
	case 0x26:
		fmt.Printf("rol $%02x\n", zeroPageAddress())
	case 0x36:
		fmt.Printf("rol $%02x, X\n", zeroPageIndexedAddress(c.X))
	case 0x2e:
		fmt.Printf("rol $%04x\n", absoluteAddress())
	case 0x3e:
		fmt.Printf("rol $%04x, X\n", absoluteIndexedAddress(c.X))
	// ROR
	case 0x6a:
		fmt.Println("ror")
	case 0x66:
		fmt.Printf("ror $%02x\n", zeroPageAddress())
	case 0x76:
		fmt.Printf("ror $%02x, X\n", zeroPageIndexedAddress(c.X))
	case 0x6e:
		fmt.Printf("ror $%04x\n", absoluteAddress())
	case 0x7e:
		fmt.Printf("ror $%04x, X\n", absoluteIndexedAddress(c.X))
	// BIT
	case 0x24:
		fmt.Printf("bit $%02x\n", zeroPageAddress())
	case 0x2c:
		fmt.Printf("bit $%04x\n", absoluteAddress())
	}
}
