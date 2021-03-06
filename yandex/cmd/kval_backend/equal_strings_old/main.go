package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	scanner.Scan()
	secondRow := scanner.Text()
	var importantCells []int
	for i, importantCellId := range strings.Split(secondRow, " ") {
		if i == 0 {
			continue
		}
		cellId, err := strconv.Atoi(importantCellId)
		if err != nil {
			panic(err)
		}
		importantCells = append(importantCells, cellId-1)
	}

	db := NewDataBase(importantCells)
	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}
		db.AddRow(strings.Split(scanner.Text(), "\t"))
	}
	fmt.Println(len(db.FindDuplicates()))
}

const (
	emptyValue = ""
)

type Row struct {
	Cells                  []string
	IsImportantCellsFilled bool
}

type DataBase struct {
	importantCells []int
	rows           []Row
	reverseIndex   map[int]map[string][]int // cellId => value => []rowId
}

func NewDataBase(importantCells []int) *DataBase {
	return &DataBase{
		importantCells: importantCells,
		reverseIndex:   make(map[int]map[string][]int),
	}
}

func (db *DataBase) AddRow(cells []string) {
	isImportantCellsFilled := true
	for _, idx := range db.importantCells {
		if cells[idx] == emptyValue {
			isImportantCellsFilled = false
		}

	}
	row := Row{
		Cells:                  cells,
		IsImportantCellsFilled: isImportantCellsFilled,
	}
	db.rows = append(db.rows, row)
	for cellId, value := range row.Cells {
		if value == emptyValue {
			continue
		}
		if db.reverseIndex[cellId] == nil {
			db.reverseIndex[cellId] = make(map[string][]int)
		}
		db.reverseIndex[cellId][value] = append(db.reverseIndex[cellId][value], len(db.rows)-1)
	}
}
func (db *DataBase) FindDuplicates() map[string][]int {
	pairs := make(map[string][]int)
	for _, ValueToRowIdsMap := range db.reverseIndex {
		for _, rowIds := range ValueToRowIdsMap {
			if len(rowIds) < 2 {
				continue
			}
			for i, ival := range rowIds {
				for j := i + 1; j < len(rowIds); j++ {
					jval := rowIds[j]

					if db.rows[ival].IsImportantCellsFilled || db.rows[jval].IsImportantCellsFilled {
						pairs[hashPair(ival, jval)] = []int{ival, jval}
					}
				}
			}
		}
	}

	for key, pair := range pairs {
		if !isEqualsRows(db.rows[pair[0]], db.rows[pair[1]]) {
			delete(pairs, key)
		}
	}
	return pairs
}

func isEqualsRows(a Row, b Row) bool {
	for i := 0; i < len(a.Cells); i++ {
		if a.Cells[i] == emptyValue || b.Cells[i] == emptyValue {
			continue
		}
		if a.Cells[i] != b.Cells[i] {
			return false
		}
	}

	return true
}

func hashPair(a, b int) string {
	if a > b {
		a, b = b, a
	}
	return fmt.Sprintf("%d_%d", a, b)
}
