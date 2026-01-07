package grid

type Direction rune

const (
	Up        Direction = 'U'
	Down      Direction = 'D'
	Left      Direction = 'L'
	Right     Direction = 'R'
	UpLeft    Direction = 'U' + 'L'
	UpRight   Direction = 'U' + 'R'
	DownLeft  Direction = 'D' + 'L'
	DownRight Direction = 'D' + 'R'
)

var upVec = Vector{0, -1}
var downVec = Vector{0, 1}
var leftVec = Vector{-1, 0}
var rightVec = Vector{1, 0}
var upLeftVec = Vector{-1, -1}
var upRightVec = Vector{1, -1}
var downLeftVec = Vector{-1, 1}
var downRightVec = Vector{1, 1}

var allDirections = []Direction{
	Up,
	Down,
	Left,
	Right,
	UpLeft,
	UpRight,
	DownLeft,
	DownRight,
}

var cardinalDirections = []Direction{
	Up,
	Down,
	Left,
	Right,
}

var directionVector = map[Direction]Vector{
	Up:        upVec,
	Down:      downVec,
	Left:      leftVec,
	Right:     rightVec,
	UpLeft:    upLeftVec,
	UpRight:   upRightVec,
	DownLeft:  downLeftVec,
	DownRight: downRightVec,
}

func (d Direction) Vector() (Vector, bool) {
	dir, ok := directionVector[d]
	return dir, ok
}

func AllDirections() []Direction {
	return allDirections
}

func CardinalDirections() []Direction {
	return cardinalDirections
}
