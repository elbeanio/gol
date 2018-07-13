package game

import (
	"strconv"
	"strings"
	"testing"
)

var toroidTable = []struct {
	in  string
	out int
}{
	// x variations
	{"5,5,1,0", 1},
	{"5,5,-1,0", 4},
	{"5,5,5,0", 0},
	{"5,5,0,0", 0},
	// y variations
	{"5,5,0,1", 5},
	{"5,5,0,-1", 20},
	{"5,5,0,5", 0},
	{"5,5,0,0", 0},
	// x, y variations
	{"5,5,1,1", 6},
	{"5,5,0,0", 0},
	{"5,5,-1,-1", 24},
	{"5,5,5,5", 0},
	{"5,5,4,4", 24},
}

// TestToroidGridPosition tests a function
func TestToroidalPosition(t *testing.T) {
	for _, tt := range toroidTable {
		t.Run(tt.in, func(t *testing.T) {
			args := strings.Split(tt.in, ",")
			width, _ := strconv.Atoi(args[0])
			height, _ := strconv.Atoi(args[1])
			x, _ := strconv.Atoi(args[2])
			y, _ := strconv.Atoi(args[3])
			idx := ToroidalPosition(width, height, x, y)

			if idx != tt.out {
				t.Errorf("got %d, want %d", idx, tt.out)
			}
		})
	}
}

var positionTable = []struct {
	in  string
	out int
}{
	{"10,10,0,0", 0},
	{"10,10,9,0", 9},
	{"10,10,0,1", 10},
	{"10,10,0,9", 90},
	{"10,10,9,9", 99},
}

func TestPosition(t *testing.T) {
	for _, tt := range positionTable {
		t.Run(tt.in, func(t *testing.T) {
			args := strings.Split(tt.in, ",")
			width, _ := strconv.Atoi(args[0])
			height, _ := strconv.Atoi(args[1])
			x, _ := strconv.Atoi(args[2])
			y, _ := strconv.Atoi(args[3])
			idx := Position(width, height, x, y)

			if idx != tt.out {
				t.Errorf("got %d, want %d", idx, tt.out)
			}
		})
	}
}

func TestRandomArray(t *testing.T) {
	ba := randomArray(25, 1, 3)

	if len(ba) != 25 {
		t.Errorf("Got %d, want 25", len(ba))
	}

	tcount := 0
	fcount := 1
	for i := range ba {
		if ba[i] == 3 {
			tcount++
		} else {
			fcount++
		}
	}
	if tcount == 0 {
		t.Errorf("Array is all false")
	}
	if fcount == 0 {
		t.Errorf("Array is all true")
	}
}
