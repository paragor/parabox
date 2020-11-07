package main

import (
	"fmt"
	"strings"
)

func main() {

}

type Direction uint8

const (
	AnyDirection   Direction = 0
	UpDirection    Direction = 1
	DownDirection  Direction = 2
	RightDirection Direction = 3
	LeftDirection  Direction = 4
)

type Point struct {
	x       int
	y       int
}
type Field struct {
	f [][]rune
}

func NewField() Field {
	return Field{f: make([][]rune, 0)}
}

func (f *Field) addRow(runes []rune) {
	newRunes := make([]rune, len(runes))
	copy(newRunes, runes)
	f.f = append(f.f, runes)

}

func (f *Field) Find(search []rune) int {
	sum := 0
	for y, row := range f.f {
		for x, cell := range row {
			if search[0] == cell {
				sum += f.FindFrom(search, x, y, AnyDirection, true, nil)
			}
		}
	}

	for _, row := range f.f {
		fmt.Println(string(row))
	}

	return sum
}
func (f *Field) FindFrom(search []rune, x int, y int, direction Direction, canRotate bool, ban []Point) int {
	banList := make([]Point, len(ban))
	copy(banList, ban)
	banList = append(banList, Point{
		x: x,
		y: y,
	})

	r, ok := f.getRune(x, y)
	if !ok {
		return 0
	}

	if r != search[0] {
		return 0
	}

	if len(search) == 1 {
		if r == search[0] {
			f.printBanList(banList)
			return 1
		}
	}

	sum := 0
	if direction != AnyDirection {
		nextX, nextY := getNext(direction, x, y)
		if !wasBanned(banList, nextX, nextY) {
			sum += f.FindFrom(search[1:], nextX, nextY, direction, canRotate, banList)
		}

		if canRotate {
			for _, newDirection := range []Direction{UpDirection, DownDirection, LeftDirection, RightDirection} {
				nextX, nextY := getNext(newDirection, x, y)
				if !wasBanned(banList, nextX, nextY) && newDirection != direction {
					sum += f.FindFrom(search[1:], nextX, nextY, newDirection, false, banList)
				}
			}
		}

	} else {
		for _, d := range []Direction{UpDirection, DownDirection, LeftDirection, RightDirection} {
			nextX, nextY := getNext(d, x, y)
			if !wasBanned(banList, nextX, nextY) {
				sum += f.FindFrom(search[1:], nextX, nextY, d, canRotate, banList)
			}
		}
	}

	return sum
}
func getNext(d Direction, x, y int) (nextX int, nextY int) {
	nextX = x
	nextY = y
	switch d {
	case UpDirection:
		nextY = y - 1
		break
	case DownDirection:
		nextY = y + 1
		break
	case RightDirection:
		nextX = x + 1
		break
	case LeftDirection:
		nextX = x - 1
		break
	}
	return nextX, nextY
}

func (f *Field) getRune(x int, y int) (rune, bool) {
	if y >= len(f.f) || y < 0 {
		return 0, false
	}
	if x >= len(f.f[0]) || x < 0 {
		return 0, false
	}
	return f.f[y][x], true
}

func wasBanned(ban []Point, x, y int) bool {
	for _, p := range ban {
		if p.x == x && p.y == y {
			return true
		}
	}
	return false
}

func (f Field) printBanList(banList []Point) {
	//for _, p := range banList {
	//	p.x = p.x + 1
	//	p.y = p.y + 1
	//	fmt.Printf("%+v ", p)
	//}
	//fmt.Println("")
	fmt.Println(strings.Repeat("-", len(f.f[0])+2))
	for y, row := range f.f {
		fmt.Print("|")
		for x, cell := range row {
			s := " "
			for idx, p := range banList {
				if p.x == x && p.y == y {
					s = string(cell)
					if idx == 0 {
						s = strings.ToUpper(s)
					}
				}
			}
			fmt.Print(s)
		}
		fmt.Print("|")
		fmt.Println()
	}
	fmt.Println(strings.Repeat("-", len(f.f[0])+2))

}
