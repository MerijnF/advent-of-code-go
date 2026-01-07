package grid

import "errors"

type Grid struct {
	Cells  [][]rune
	Width  int
	Height int
	WrapX  bool
	WrapY  bool
}

func NewGrid(cells [][]rune, wrapX, wrapY bool) Grid {
	height := len(cells)
	width := 0
	if height > 0 {
		width = len(cells[0])
	}
	return Grid{
		Cells:  cells,
		Width:  width,
		Height: height,
		WrapX:  wrapX,
		WrapY:  wrapY,
	}
}

func (g *Grid) Get(x, y int) (rune, error) {
	if g.WrapX {
		x = Wrap(x, g.Width)
	}
	if g.WrapY {
		y = Wrap(y, g.Height)
	}
	return g.getNoWrap(x, y)
}

func (g *Grid) getNoWrap(x, y int) (rune, error) {
	if !g.inBoundsNoWrap(x, y) {
		return ' ', errors.New("out of bounds")
	}
	return g.Cells[y][x], nil
}

func (g *Grid) Set(x, y int, value rune) error {
	if g.WrapX {
		x = Wrap(x, g.Width)
	}
	if g.WrapY {
		y = Wrap(y, g.Height)
	}
	return g.setNoWrap(x, y, value)
}

func (g *Grid) setNoWrap(x, y int, value rune) error {
	if !g.inBoundsNoWrap(x, y) {
		return errors.New("out of bounds")
	}
	g.Cells[y][x] = value
	return nil
}

func (g *Grid) InBounds(x, y int) bool {
	if g.WrapX {
		x = Wrap(x, g.Width)
	}
	if g.WrapY {
		y = Wrap(y, g.Height)
	}
	return InBounds(x, y, g.Width, g.Height)
}

func (g *Grid) inBoundsNoWrap(x, y int) bool {
	return InBounds(x, y, g.Width, g.Height)
}

func Wrap(value, max int) int {
	if value < 0 {
		return (value%max + max) % max
	} else if value >= max {
		return value % max
	}
	return value
}

func InBounds(x, y, width, height int) bool {
	return x >= 0 && x < width && y >= 0 && y < height
}

type AdjacentValue struct {
	Position  Vector
	Value     rune
	Direction Direction
}

func (g *Grid) GetCardinalyAdjecent(x, y int) ([]AdjacentValue, error) {
	values := make([]AdjacentValue, 8)
	for _, dir := range CardinalDirections() {
		dirVec, ok := dir.Vector()
		if !ok {
			return nil, errors.New("unknown direction vector")
		}
		newX := x + dirVec.X
		newY := y + dirVec.Y
		if !g.InBounds(newX, newY) {
			continue
		}
		val, err := g.Get(newX, newY)
		if err != nil {
			return nil, err
		}
		values = append(values, AdjacentValue{
			Position:  Vector{newX, newY},
			Value:     val,
			Direction: dir,
		})
	}
	return values, nil
}

func (g *Grid) GetAllAdjecent(x, y int) ([]AdjacentValue, error) {
	values := make([]AdjacentValue, 8)
	for _, dir := range AllDirections() {
		dirVec, ok := dir.Vector()
		if !ok {
			return nil, errors.New("unknown direction vector")
		}
		newX := x + dirVec.X
		newY := y + dirVec.Y
		if !g.InBounds(newX, newY) {
			continue
		}
		val, err := g.Get(newX, newY)
		if err != nil {
			return nil, err
		}
		values = append(values, AdjacentValue{
			Position:  Vector{newX, newY},
			Value:     val,
			Direction: dir,
		})
	}
	return values, nil
}
