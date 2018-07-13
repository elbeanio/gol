package game

// Game is the state object containing a list of cell states and the number of neighbours for each cell
type Game struct {
	width  int
	height int

	StateChars [5]string
	StateMax   int
	Neighbours []int
	State      []int
}

// New returns a new Game object
func New(width, height int, seed int64) *Game {
	g := new(Game)
	g.StateChars = [5]string{" ", ".", "o", "O", "@"}
	g.StateMax = len(g.StateChars) - 1
	g.Init(width, height, seed)
	return g
}

// Init initialises (or resets) the game state
func (g *Game) Init(width, height int, seed int64) {
	g.width = width
	g.height = height
	g.State = randomArray(width*height, seed, g.StateMax)
	g.Neighbours = make([]int, width*height)
	for y := 0; y < g.height; y++ {
		for x := 0; x < g.width; x++ {
			idx := Position(width, height, x, y)
			if g.State[idx] == g.StateMax {
				g.setNeighbours(x, y, g.StateMax)
			}
		}
	}
}

func (g *Game) setNeighbours(x int, y int, state int) {
	delta := -1
	if state == g.StateMax {
		delta = 1
	}

	for dy := -1; dy < 2; dy++ {
		for dx := -1; dx < 2; dx++ {
			if dx == 0 && dy == 0 {
				continue
			}

			idx := ToroidalPosition(g.width, g.height, x+dx, y+dy)
			g.Neighbours[idx] += delta
		}
	}
}

// Update moves the game state on
func (g *Game) Update() {
	currentN := make([]int, g.width*g.height)
	copy(currentN, g.Neighbours)

	for y := 0; y < g.height; y++ {
		for x := 0; x < g.width; x++ {
			idx := Position(g.width, g.height, x, y)
			cs := g.State[idx]
			ns := cs
			nn := currentN[idx]

			if cs == g.StateMax {
				if nn < 2 || nn > 3 {
					ns--
				}
			} else {
				if nn == 3 {
					ns = g.StateMax
				}
			}

			if ns != cs {
				g.State[idx] = ns
				g.setNeighbours(x, y, ns)
			} else if cs > 0 && ns < g.StateMax {
				g.State[idx] = ns - 1
			}
		}
	}
}
