package day3

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strconv"
)

func SolvePartOne(reader io.Reader) (string, error) {
	s := bufio.NewScanner(reader)

	sum := 0

	for s.Scan() {
		l := s.Text()
		if len(l) == 0 {
			log.Println("skip empty line")
			continue
		}

		if len(l)%2 != 0 {
			return "", fmt.Errorf("invalid input line, odd number of items")
		}

		items := make(map[item]int)
		for i := 0; i < len(l)/2; i++ {
			items[item(l[i])] = 1
		}

		for i := len(l) / 2; i < len(l); i++ {
			items[item(l[i])] *= 2
		}

		for k, v := range items {
			if v > 1 {
				sum += k.priority()
			}
		}
	}

	return strconv.Itoa(sum), nil
}

type item byte

func (i item) priority() int {
	if i > 'Z' {
		return int(i - 96)
	} else {
		return int(i - 38)
	}
}

func SolvePartTwo(reader io.Reader) (string, error) {
	s := bufio.NewScanner(reader)

	sum := 0

	items := make([]map[item]int, 3)
	elf := 0

	for s.Scan() {
		l := s.Text()
		if len(l) == 0 {
			log.Println("skip empty line")
			continue
		}

		items[elf] = make(map[item]int, len(l))
		for i := range l {
			items[elf][item(l[i])] = 1
		}

		if elf < 2 {
			elf++
			continue
		}

		for k := range items[0] {
			if items[1][k]*items[2][k] != 0 {
				sum += k.priority()
			}
		}
		elf = 0
	}

	return strconv.Itoa(sum), nil
}
