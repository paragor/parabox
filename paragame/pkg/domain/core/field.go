package core

type CellState int8

const (
	EmptyCell CellState = 0
)

type Field interface {
	Get(x, y int) CellState
	Set(x, y int, state CellState)
	Next()
	MaxX() int
	MaxY() int
}

type StdField struct {
	now  [][]CellState
	next [][]CellState
	maxX int
	maxY int
}

func NewStdField(maxX int, maxY int) *StdField {
	now := make([][]CellState, maxY)
	next := make([][]CellState, maxY)
	for y := 0; y < maxY; y++ {
		now[y] = make([]CellState, maxX)
		next[y] = make([]CellState, maxX)
		for x := 0; x < maxX; x++ {
			now[y][x] = EmptyCell
			next[y][x] = EmptyCell
		}
	}
	return &StdField{maxX: maxX, maxY: maxY, now: now, next: next}
}

func (f *StdField) Get(x, y int) CellState {
	return f.now[y][x]
}

func (f *StdField) Set(x, y int, state CellState) {
	f.next[y][x] = state
}

func (f *StdField) Next() {
	for y, row := range f.next {
		for x, cell := range row {
			f.now[y][x] = cell
		}
	}
}

func (f *StdField) MaxX() int {
	return f.maxX
}

func (f *StdField) MaxY() int {
	return f.maxY
}
