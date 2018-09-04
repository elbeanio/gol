package io

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/elbeanio/gol/game"
)

// GetTermSize returns the size of the tty we're running in
func GetTermSize() (int, int) {
	// query the term size
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, _ := cmd.Output()

	// get the values out
	size := strings.Split(strings.Trim(string(out), "\n"), " ")
	sy, err := strconv.Atoi(size[0])
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	sx, err := strconv.Atoi(size[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	if sx > 0 && sy > 0 {
		return sx, sy
	}

	return 10, 10

}

// ClearScreen prints the escape code to clear the term
func ClearScreen() {
	fmt.Print("\033[2J")
}

// TopLeft prints the escape code to move the cursor to 0,0
func TopLeft() {
	fmt.Print("\033[0;0H")
}

// CursorLeft prints the escape code to move the cursor left n cols
func CursorLeft(chars int) {
	fmt.Printf("\033[%dD", chars+1)
}

// CursorUp prints the escape code to move the cursor up n rows
func CursorUp(n int) {
	fmt.Printf("\033[%dA", n-1)
}

func printStatusLine(g *game.Game, b *bytes.Buffer, oType string) {
	b.WriteString(fmt.Sprintf("Output Type: %s        Frame Delay: %d", oType, g.FrameDelay))
}

// PrintGridBasic prints a game state with simple true / false values
func PrintGridBasic(g *game.Game) {
	var buffer bytes.Buffer
	buffer.Grow(g.Width * g.Height)
	for y := 0; y < g.Height; y++ {
		for x := 0; x < g.Width; x++ {
			pState := g.State[game.Position(g.Width, g.Height, x, y)]
			if pState == g.StateMax {
				buffer.WriteString("█")
			} else {
				buffer.WriteString(" ")
			}
		}
	}
	printStatusLine(g, &buffer, "Basic")
	fmt.Print(buffer.String())
}

// PrintGridFade prints a game state where a cell death fades out over a number of states
func PrintGridFade(g *game.Game) {
	var buffer bytes.Buffer
	buffer.Grow(g.Width * g.Height)
	for y := 0; y < g.Height; y++ {
		for x := 0; x < g.Width; x++ {
			pState := g.State[game.Position(g.Width, g.Height, x, y)]
			buffer.WriteString(g.StateChars[pState])
		}
	}
	printStatusLine(g, &buffer, "Fading")
	fmt.Print(buffer.String())
}
