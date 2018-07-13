package main

import (
	"bufio"
	"fmt"
	"gol/game"
	"gol/io"
	"os"
	"strings"
	"time"

	"golang.org/x/crypto/ssh/terminal"
)

func introScreen(w, h int) {
	io.ClearScreen()
	io.TopLeft()
	lines := 7
	boxStart := strings.Repeat(" ", (w-20)/2)
	fmt.Println(strings.Repeat("\n", (h-lines)/2))
	fmt.Println(boxStart, "  Game of Life")
	fmt.Println(boxStart, "--------------------------------")
	fmt.Println(boxStart, " | q or <esc>: quit           |")
	fmt.Println(boxStart, " | <space>   : pause          |")
	fmt.Println(boxStart, " | r         : restart        |")
	fmt.Println(boxStart, " | o         : change output  |")
	fmt.Println(boxStart, "--------------------------------")
	fmt.Println(boxStart, " Press Enter to start")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func main() {
	sx, sy := io.GetTermSize()
	frameDelay := 100
	running := true

	introScreen(sx, sy)
	state := game.New(sx, sy, time.Now().UnixNano())

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
		io.PrintGrid(sx, sy, state)

		select {
		case chr := <-input:
			switch chr {
			case 'q':
				fallthrough
			case '\033':
				terminal.Restore(0, oldState)
				os.Exit(0)
			case '+':
				if frameDelay > 10 {
					frameDelay -= 5
				}
			case '-':
				if frameDelay < 100 {
					frameDelay += 5
				}
			case ' ':
				running = !running
			case 'r':
				state.Init(sx, sy, time.Now().UnixNano())
			}
		default: // empty select default so it doesn't block
		}

		// TODO turn this into a fixed refresh rate rather than just waiting
		time.Sleep(time.Duration(frameDelay) * time.Millisecond)
	}
}
