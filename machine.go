package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
	"strings"
)

var (
	cpuClockSpeed = 1789773
	running       = true

	ppu   Ppu
	video Video
	pads  [2]*Controller

	gamename       string
	saveStateFile  string
	batteryRamFile string

	cpuprofile = flag.String("cprof", "", "write cpu profile to file")
)

const (
	SaveState = iota
	LoadState
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	if len(os.Args) < 2 {
		fmt.Println("Please specify a ROM file")
		return
	}

	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}

		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	pads[0] = new(Controller)
	pads[1] = new(Controller)
	pads[0].Init(0)
	pads[1].Init(0)

	v := ppu.Init()

	if contents, err := ioutil.ReadFile(os.Args[len(os.Args)-1]); err == nil {

		if err = LoadRom(contents); err != nil {
			fmt.Println(err.Error())
			return
		}

		// Set the game name for save states
		path := strings.Split(os.Args[1], "/")
		gamename = strings.Split(path[len(path)-1], ".")[0]
		saveStateFile = fmt.Sprintf(".%s.state", gamename)
		batteryRamFile = fmt.Sprintf(".%s.battery", gamename)

	} else {
		fmt.Println(err.Error())
		return
	}

	interrupt := make(chan int)

	// Main runloop, in a separate goroutine so that
	// the video rendering can happen on this one
	go func(c <-chan int) {
		StartCpu(func(cycles int) {
			for i := 0; i < 3*cycles; i++ {
				ppu.Step()
			}
		})
	}(interrupt)

	r := video.Init(v, gamename)

	go ReadInput(r, interrupt)

	// This needs to happen on the main thread for OSX
	runtime.LockOSThread()
	video.Render()

	return
}
