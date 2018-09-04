package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/elbeanio/gol/game"
	"github.com/elbeanio/gol/io"

	"golang.org/x/crypto/ssh/terminal"
)

func introScreen(w, h int) {
	io.ClearScreen()
	io.TopLeft()
	lines := 9
	boxStart := strings.Repeat(" ", (w-20)/2)
	fmt.Println(strings.Repeat("\n", (h-lines)/2))
	fmt.Println(boxStart, "---------------------------------------")
	fmt.Println(boxStart, "            Game of Life               ")
	fmt.Println(boxStart, "---------------------------------------")
	fmt.Println(boxStart, " | q or <esc>: quit                  | ")
	fmt.Println(boxStart, " | <space>   : pause                 | ")
	fmt.Println(boxStart, " | r         : restart               | ")
	fmt.Println(boxStart, " | o         : change output         | ")
	fmt.Println(boxStart, "---------------------------------------")
	fmt.Println(boxStart, "       Press Enter to continue         ")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func main() {
	sx, sy := io.GetTermSize()
	frameDelay := 100
	running := true

	introScreen(sx, sy)
	state := game.New(sx, sy, frameDelay, time.Now().UnixNano())

	// put term in raw modez
	oldState, err := terminal.MakeRaw(0)
	if err != nil {
		panic(err)
	}

	input := make(chan rune, 1)
	go io.GetInput(input)

	for {
		// update
		if running {
			state.Update()
		}

		// write to screen
		io.ClearScreen()
		io.TopLeft()
		switch state.OutputType {
		case 0:
			io.PrintGridBasic(state)
		case 1:
			io.PrintGridFade(state)
		}

		select {
		case chr := <-input:
			switch chr {
			case 'q':
				fallthrough
			case '\033':
				terminal.Restore(0, oldState)
				os.Exit(0)
			case 'o':
				state.OutputType++
				if state.OutputType > 1 {
					state.OutputType = 0
				}
			case '+':
				if state.FrameDelay > 20 {
					state.FrameDelay -= 5
				}
			case '-':
				if state.FrameDelay < 100 {
					state.FrameDelay += 5
				}
			case ' ':
				running = !running
			case 'r':
				state.Init(time.Now().UnixNano())
			}
		default: // empty select default so it doesn't block
		}

		// TODO turn this into a fixed refresh rate rather than just waiting
		time.Sleep(time.Duration(state.FrameDelay) * time.Millisecond)
	}
}
