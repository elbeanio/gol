package game

import (
	"math/rand"
)

// ToroidalPosition returns an index from a toroidal grid of width * height
func ToroidalPosition(width, height, x, y int) int {
	gx, gy := x, y

	if x < 0 {
		gx = width + x
	}
	if x > width-1 {
		gx = x - width
	}
	if y < 0 {
		gy = height + y
	}
	if y > height-1 {
		gy = y - height
	}
	return Position(width, height, gx, gy)
}

// Position returns a 2d coordinate in a 1d array
func Position(width, height, x, y int) int {
	return (y*width + x)
}

func randomArray(size int, seed int64, maxState int) []int {
	source := rand.NewSource(seed)
	random := rand.New(source)

	arr := make([]int, size)
	for i := range arr {
		if random.Intn(100) < 30 {
			arr[i] = maxState
		} else {
			arr[i] = 0
		}
	}
	return arr
}
