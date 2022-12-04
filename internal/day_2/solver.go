package day2

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"
)

func SolvePartOne(reader io.Reader) (string, error) {
	s := bufio.NewScanner(reader)

	score := 0

	for s.Scan() {
		l := s.Text()
		if len(l) == 0 {
			log.Println("skip empty line")
			continue
		}

		subs := strings.Split(l, " ")
		if len(subs) != 2 {
			return "", fmt.Errorf("expected two rows, got %v", l)
		}

		var oponent hand
		switch subs[0] {
		case "A":
			oponent = rock
		case "B":
			oponent = paper
		case "C":
			oponent = scissors
		default:
			return "", fmt.Errorf("mapping %v to oponent hand failed", subs[0])
		}

		var player hand
		switch subs[1] {
		case "X":
			player = rock
		case "Y":
			player = paper
		case "Z":
			player = scissors
		default:
			return "", fmt.Errorf("mapping %v to player hand failed", subs[1])
		}

		result := play(player, oponent)
		score += player.score() + result.score()
	}

	return strconv.Itoa(score), nil
}

type hand uint

const (
	rock hand = iota
	paper
	scissors
)

func (h hand) score() int {
	return map[hand]int{
		rock:     1,
		paper:    2,
		scissors: 3,
	}[h]
}

func play(h1, h2 hand) gameResult {
	if h1 == h2 {
		return draw
	}

	if h1 == rock && h2 == scissors ||
		h1 == paper && h2 == rock ||
		h1 == scissors && h2 == paper {
		return win
	}

	return loss
}

type gameResult uint

const (
	win gameResult = iota
	loss
	draw
)

func (r gameResult) score() int {
	return map[gameResult]int{
		win:  6,
		loss: 0,
		draw: 3,
	}[r]
}

func SolvePartTwo(reader io.Reader) (string, error) {
	s := bufio.NewScanner(reader)

	score := 0

	for s.Scan() {
		l := s.Text()
		if len(l) == 0 {
			log.Println("skip empty line")
			continue
		}

		subs := strings.Split(l, " ")
		if len(subs) != 2 {
			return "", fmt.Errorf("expected two rows, got %v", l)
		}

		var oponent hand
		switch subs[0] {
		case "A":
			oponent = rock
		case "B":
			oponent = paper
		case "C":
			oponent = scissors
		default:
			return "", fmt.Errorf("mapping %v to oponent hand failed", subs[0])
		}

		var expectedResult gameResult
		switch subs[1] {
		case "X":
			expectedResult = loss
		case "Y":
			expectedResult = draw
		case "Z":
			expectedResult = win
		default:
			return "", fmt.Errorf("mapping %v to expected result failed", subs[1])
		}

		playerHand, err := inferHand(oponent, expectedResult)
		if err != nil {
			return "", err
		}

		score += playerHand.score() + expectedResult.score()
	}

	return strconv.Itoa(score), nil
}

func inferHand(oponentHand hand, expectedResult gameResult) (hand, error) {
	if expectedResult == win {
		switch oponentHand {
		case rock:
			return paper, nil
		case paper:
			return scissors, nil
		case scissors:
			return rock, nil
		default:
			return rock, fmt.Errorf("unknown hand type %v", oponentHand)
		}
	}

	if expectedResult == loss {
		switch oponentHand {
		case rock:
			return scissors, nil
		case paper:
			return rock, nil
		case scissors:
			return paper, nil
		default:
			return rock, fmt.Errorf("unknown hand type %v", oponentHand)
		}
	}

	return oponentHand, nil
}
