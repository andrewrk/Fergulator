package main

/*
#cgo LDFLAGS: zelda.o
#include "stdint.h"
#include "rom.h"

extern uint8_t rom_ppustatus();
extern void rom_cycle(uint8_t);
extern void rom_ppuctrl(uint8_t);
extern void rom_ppumask(uint8_t);
extern void rom_ppuaddr(uint8_t);
extern void rom_setppudata(uint8_t);
extern void rom_oamaddr(uint8_t);
extern void rom_setoamdata(uint8_t);
extern void rom_setppuscroll(uint8_t);
*/
import "C"

//export rom_ppustatus
func rom_ppustatus() C.uint8_t {
	v, _ := ppu.ReadStatus()
	return C.uint8_t(v)
}

//export rom_ppuctrl
func rom_ppuctrl(v C.uint8_t) {
	ppu.WriteControl(uint8(v))
}

//export rom_ppumask
func rom_ppumask(v C.uint8_t) {
	ppu.WriteMask(uint8(v))
}

//export rom_ppuaddr
func rom_ppuaddr(v C.uint8_t) {
	ppu.WriteAddress(uint8(v))
}

//export rom_setppudata
func rom_setppudata(v C.uint8_t) {
	ppu.WriteData(uint8(v))
}

//export rom_oamaddr
func rom_oamaddr(v C.uint8_t) {
	ppu.WriteOamAddress(uint8(v))
}

//export rom_setoamdata
func rom_setoamdata(v C.uint8_t) {
	ppu.WriteOamData(uint8(v))
}

//export rom_setppuscroll
func rom_setppuscroll(v C.uint8_t) {
	ppu.WriteScroll(uint8(v))
}

var cycleCbFn func(int)

//export rom_cycle
func rom_cycle(v C.uint8_t) {
	cycleCbFn(int(v))
}

func StartCpu(cb func(int)) {
	cycleCbFn = cb
	C.rom_start()
}
