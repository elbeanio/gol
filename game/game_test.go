package game

import (
	"testing"
)

func TestNewGame(t *testing.T) {
	g := New(5, 5, 1)
	if g.width != 5 {
		t.Errorf("Expected width of 5, got %d", g.width)
	}
	if g.height != 5 {
		t.Errorf("Expected height of 5, got %d", g.height)
	}
	if len(g.State) != 25 {
		t.Errorf("Expected State to be length 25, got %d", len(g.State))
	}
	if len(g.Neighbours) != 25 {
		t.Errorf("Expected neighbours to be length 25, got %d", len(g.Neighbours))
	}
}

func TestSetNeighbours(t *testing.T) {
	var seed int64 = 1
	g := New(5, 5, seed)
	g.setNeighbours(2, 2, g.StateMax)
	g.setNeighbours(3, 3, g.StateMax)
	neighbours := []int{
		3, 3, 2, 2, 2,
		4, 3, 3, 3, 2,
		7, 4, 4, 5, 4,
		4, 3, 5, 4, 2,
		2, 2, 2, 3, 2,
	}

	for i := range neighbours {
		if neighbours[i] != g.Neighbours[i] {
			t.Errorf("Neighbours fail expected:\n%vGot:\n%v", neighbours, g.Neighbours)
			break
		}
	}
}

func TestUpdate(t *testing.T) {
	var seed int64 = 1
	g := New(5, 5, seed)
	g.Update()

}
