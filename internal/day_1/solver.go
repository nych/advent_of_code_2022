package day1

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strconv"
)

func SolvePartOne(reader io.Reader) (string, error) {
	if reader == nil {
		return "", fmt.Errorf("reader arg cannot be nil")
	}

	elves, err := loadElves(reader)
	if err != nil {
		return "", fmt.Errorf("loading elves failed, %v", err)
	}

	sort.Slice(elves, func(i, j int) bool {
		return elves[i].carryingCalories() > elves[j].carryingCalories()
	})

	return strconv.Itoa(elves[0].carryingCalories()), nil
}

func SolvePartTwo(reader io.Reader) (string, error) {
	if reader == nil {
		return "", fmt.Errorf("reader arg cannot be nil")
	}

	elves, err := loadElves(reader)
	if err != nil {
		return "", fmt.Errorf("loading elves failed, %v", err)
	}

	sort.Slice(elves, func(i, j int) bool {
		return elves[i].carryingCalories() > elves[j].carryingCalories()
	})

	sum := 0
	for i := range elves {
		if i > 2 {
			break
		}
		sum += elves[i].carryingCalories()
	}

	return strconv.Itoa(sum), nil
}

func loadElves(reader io.Reader) ([]elf, error) {
	s := bufio.NewScanner(reader)

	elves := make([]elf, 0)
	elf := createElf(0)

	for s.Scan() {
		l := s.Text()

		if len(l) == 0 {
			elves = append(elves, elf)
			elf = createElf(elf.id + 1)
			continue
		}

		calories, err := strconv.Atoi(l)
		if err != nil {
			return nil, err
		}

		i := item{
			id:       calories,
			calories: calories,
		}

		if _, ok := elf.items[i]; !ok {
			elf.items[i] = 0
		}
		elf.items[i]++
	}

	if s.Err() != nil {
		return nil, s.Err()
	}

	if len(elves) == 0 {
		elves = append(elves, elf)
	}

	return elves, nil
}

type elf struct {
	id    int
	items map[item]int
}

type item struct {
	id       int
	calories int
}

func createElf(id int) elf {
	return elf{
		id:    id,
		items: make(map[item]int),
	}
}

func (e elf) carryingCalories() int {
	var r int = 0
	for i, n := range e.items {
		r += i.calories * n
	}
	return r
}
