package io

import (
	"bytes"
	"fmt"
	"gol/game"
	"os"
	"os/exec"
	"strconv"
	"strings"
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

// PrintGrid prints a grid in the terminal
func PrintGrid(sizeX, sizeY int, g *game.Game) {
	var buffer bytes.Buffer
	buffer.Grow(sizeX * sizeY)
	for y := 0; y < sizeY; y++ {
		for x := 0; x < sizeX; x++ {
			pState := g.State[game.Position(sizeX, sizeY, x, y)]
			buffer.WriteString(g.StateChars[pState])
		}

		buffer.WriteString("\n")
	}
	fmt.Print(buffer.String())
}
