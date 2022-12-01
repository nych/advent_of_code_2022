package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	day1 "github.com/nych/advent_of_code_2022/internal/day_1"
)

var (
	day uint
)

func init() {
	flag.UintVar(&day, "day", 1, "defines day to solve")
}

func main() {
	flag.Parse() // will panic in case of err

	f, err := os.Open(fmt.Sprintf("resources/day_%v/input.txt", day))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Printf("closing input file failed, %v", err)
		}
	}()

	var solvePartOne func(io.Reader) (string, error)
	var solvePartTwo func(io.Reader) (string, error)

	switch day {
	case 1:
		solvePartOne = day1.SolvePartOne
		solvePartTwo = day1.SolvePartTwo
	default:
		panic(fmt.Errorf("solver for day %v not yet implemented", day))
	}

	partOne, err := solvePartOne(f)
	if err != nil {
		log.Printf("solving part one failed, %v", err)
		partOne = "failed"
	}

	if _, err := f.Seek(0, 0); err != nil {
		panic(fmt.Errorf("reset reader failed, %v", err))
	}

	partTwo, err := solvePartTwo(f)
	if err != nil {
		log.Printf("solving part two failed, %v", err)
		partTwo = "failed"
	}

	fmt.Printf("partOne: %v\npartTwo: %v\n", partOne, partTwo)
}
