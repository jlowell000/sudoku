package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/jlowell000/sudoku/internal/puzzle"
)

type DataRow struct {
	Puzzle, Solution string
}

func main() {

	// open file
	f, err := os.Open("./test/sudoku.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	i := 0
	for {
		rec, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		// do something with read line
		if i > 0 {
			// 90000001
			dataRow := DataRow{rec[0], rec[1]}
			p := dataRowToPuzzle(dataRow)
			valid := p.Valid()
			// log.Default().Println(dataRow.Solution)
			// log.Default().Println(p.ReadableCurrentValueString())
			if i%10000 == 0 {
				per := i / 90000001 * 100
				log.Default().Printf("parsed %d percent (%d/90000001)", per, i)
			}
			if !valid {
				log.Default().Printf("parsed properly: %v, valid: %v", dataRow.Solution == p.CurrentValueString(), valid)
			}
		}
		i++
	}

}

func dataRowToPuzzle(dataRow DataRow) puzzle.Puzzle {
	p, err := puzzle.New(9)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(dataRow.Puzzle); i++ {
		initialV, _ := strconv.Atoi(string(dataRow.Puzzle[i]))
		trueV, _ := strconv.Atoi(string(dataRow.Solution[i]))
		pos := indexToPosition(i)

		p.TrueValues[pos] = trueV
		p.CurrentValues[pos] = trueV
		p.InitialValues[pos] = initialV
	}

	return p
}

func indexToPosition(i int) puzzle.Position {
	v := i + 1
	div := (v / 9) + 1
	mod := (v % 9)
	if div == 10 {
		div = 9
		mod = 9
	}
	if mod == 0 {
		div--
		mod = 9
	}
	return puzzle.Position{X: div, Y: mod}
}
