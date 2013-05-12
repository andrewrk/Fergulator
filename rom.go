package main

import (
	"errors"
	"fmt"
)

const (
	Size1k  = 0x0400
	Size4k  = 0x1000
	Size8k  = 0x2000
	Size16k = 0x4000
	Size32k = 0x8000
)

type Mapper interface {
	Write(v uint8, a int)
	BatteryBacked() bool
}

// Nrom
type Rom struct {
	VromBanks [][]uint8

	PrgBankCount int
	ChrRomCount  int
	Battery      bool
	Data         []byte
}

func WriteVramBank(rom [][]uint8, bank, dest, size int) {
	bank %= len(rom)

	for i := 0; i < size; i++ {
		ppu.Vram[i+dest] = rom[bank][i]
	}
}

func (m *Rom) BatteryBacked() bool {
	return m.Battery
}

func LoadRom(rom []byte) error {
	r := new(Rom)

	if string(rom[0:3]) != "NES" {
		return errors.New("Invalid ROM file")

		if rom[3] != 0x1a {
			return errors.New("Invalid ROM file")
		}
	}

	r.PrgBankCount = int(rom[4])
	r.ChrRomCount = int(rom[5])

	fmt.Printf("-----------------\nROM:\n  ")

	fmt.Printf("PRG-ROM banks: %d\n  ", r.PrgBankCount)
	fmt.Printf("CHR-ROM banks: %d\n  ", r.ChrRomCount)

	fmt.Printf("Mirroring: ")
	switch rom[6] & 0x1 {
	case 0x0:
		fmt.Printf("Horizontal\n  ")
		ppu.Nametables.SetMirroring(MirroringHorizontal)
	case 0x1:
		fmt.Printf("Vertical\n  ")
		ppu.Nametables.SetMirroring(MirroringVertical)
	}

	if (rom[6]>>0x1)&0x1 == 0x1 {
		r.Battery = true
	}

	r.Data = rom[16:]

	// Everything after PRG-ROM
	chrRom := r.Data[0x4000*r.PrgBankCount:]

	r.VromBanks = make([][]uint8, r.ChrRomCount*2)
	for i := 0; i < r.ChrRomCount*2; i++ {
		// Move 16kb chunk to 16kb bank
		bank := make([]uint8, 0x1000)
		for x := 0; x < 0x1000; x++ {
			bank[x] = uint8(chrRom[(0x1000*i)+x])
		}

		r.VromBanks[i] = bank
	}

	// If we have CHR-ROM, load the first two banks
	// into VRAM region 0x0000-0x1000
	if r.ChrRomCount > 0 {
		if r.ChrRomCount == 1 {
			WriteVramBank(r.VromBanks, 0, 0x0000, Size4k)
			WriteVramBank(r.VromBanks, 1, 0x1000, Size4k)
		} else {
			WriteVramBank(r.VromBanks, 0, 0x0000, Size4k)
			WriteVramBank(r.VromBanks, len(r.VromBanks)-1, 0x1000, Size4k)
		}
	}

	// Check mapper, get the proper type
	mapper := (uint8(rom[6])>>4 | (uint8(rom[7]) & 0xF0))
	fmt.Printf("Mapper: 0x%X -> ", mapper)
	switch mapper {
	case 0x00:
		fallthrough
	case 0x40:
		fallthrough
	case 0x41:
		// NROM
		fmt.Printf("NROM\n")
		return nil
	default:
		// Unsupported
		fmt.Printf("Unsupported\n")
		return errors.New(fmt.Sprintf("Unsupported memory mapper: 0x%X", mapper))
	}

	fmt.Printf("-----------------\n")

	return nil
}
