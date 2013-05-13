package main

import (
	"github.com/scottferg/Go-SDL/sdl"
	"fmt"
)

const (
	PadBtnA = iota
	PadBtnB
	PadBtnSelect
	PadBtnStart
	PadBtnUp
	PadBtnDown
	PadBtnLeft
	PadBtnRight
)

type Controller struct {
	ActualButtonState [2][8]Word
	ReportedButtonState [2][8]Word
	ButtonIndex [2]int
	StrobeOn    bool
}

var keyToBtn = map[uint32]int{
	sdl.K_h: PadBtnA,
	sdl.K_t: PadBtnB,
	sdl.K_LSHIFT: PadBtnSelect,
	sdl.K_RETURN: PadBtnStart,
	sdl.K_COMMA: PadBtnUp,
	sdl.K_o: PadBtnDown,
	sdl.K_a: PadBtnLeft,
	sdl.K_e: PadBtnRight,
}

func (c *Controller) Init() {
	for padIndex := 0; padIndex <= 1; padIndex++ {
		for i := 0; i < 8; i++ {
			c.ActualButtonState[padIndex][i] = 0x40
			c.ReportedButtonState[padIndex][i] = 0x40
		}
	}
}


func (c *Controller) SetButtonState(k sdl.KeyboardEvent, state Word) {
	btn, ok := keyToBtn[k.Keysym.Sym]
	if !ok {
		return
	}
	movie = append(movie, MovieFrame{
		uint64(totalCpuCycles),
		0,
		uint8(btn),
		uint8(state),
	})
	c.ActualButtonState[0][btn] = state
	if c.StrobeOn {
		c.ReportedButtonState[0][btn] = state
	}
}

func (c *Controller) Write(v Word) {
	fmt.Printf("pad_write $%02x\n", v)
	c.StrobeOn = v&0x1 == 1
	if c.StrobeOn {
		// copy actual to reported
		for padIndex := 0; padIndex <= 1; padIndex++ {
			for i, v := range c.ActualButtonState[padIndex] {
				c.ReportedButtonState[padIndex][i] = v
			}
			c.ButtonIndex[padIndex] = 0
		}
	}
}

func (c *Controller) Read(padIndex int) Word {
	if c.ButtonIndex[padIndex] >= len(c.ReportedButtonState[padIndex]) {
		return 1
	}
	v := c.ReportedButtonState[padIndex][c.ButtonIndex[padIndex]]
	if c.StrobeOn {
		c.ButtonIndex[padIndex] = 0
	} else {
		c.ButtonIndex[padIndex] += 1
	}
	return v
}

func ReadInput(r chan [2]int, i chan int) {
	for {
		select {
		case ev := <-sdl.Events:
			switch e := ev.(type) {
			case sdl.ResizeEvent:
				r <- [2]int{int(e.W), int(e.H)}
			case sdl.QuitEvent:
				running = false
				video.Close()
			case sdl.KeyboardEvent:
				switch e.Keysym.Sym {
				case sdl.K_SPACE:
					saveMovie()
					running = false
					video.Close()
				case sdl.K_ESCAPE:
					running = false
				case sdl.K_r:
					// Trigger reset interrupt
					if e.Type == sdl.KEYDOWN {
						cpu.RequestInterrupt(InterruptReset)
					}
				case sdl.K_l:
					if e.Type == sdl.KEYDOWN {
						i <- LoadState
					}
				case sdl.K_p:
					if e.Type == sdl.KEYDOWN {
						// Enable/disable scanline sprite limiter flag
						ppu.SpriteLimitEnabled = !ppu.SpriteLimitEnabled
					}
				case sdl.K_s:
					if e.Type == sdl.KEYDOWN {
						i <- SaveState
					}
				case sdl.K_z:
					if e.Type == sdl.KEYDOWN {
						ppu.OverscanEnabled = !ppu.OverscanEnabled
					}
				case sdl.K_i:
					if e.Type == sdl.KEYDOWN {
						audioEnabled = !audioEnabled
					}
				case sdl.K_1:
					if e.Type == sdl.KEYDOWN {
						r <- [2]int{256, 240}
					}
				case sdl.K_2:
					if e.Type == sdl.KEYDOWN {
						r <- [2]int{512, 480}
					}
				case sdl.K_3:
					if e.Type == sdl.KEYDOWN {
						r <- [2]int{768, 720}
					}
				case sdl.K_4:
					if e.Type == sdl.KEYDOWN {
						r <- [2]int{1024, 960}
					}
				}

				switch e.Type {
				case sdl.KEYDOWN:
					pads.SetButtonState(e, 0x41)
				case sdl.KEYUP:
					pads.SetButtonState(e, 0x40)
				}
			}
		}
	}
}
