package main

import (
	"log"
	"os"
	"strings"
)

type Shape string
type Result string

const (
	ROCK     Shape  = "ROCK"
	PAPER    Shape  = "PAPER"
	SCISSORS Shape  = "SCISSORS"
	WIN      Result = "WIN"
	LOSE     Result = "LOSE"
	DRAW     Result = "DRAW"
)

func ParseShape(s string) Shape {
	switch s {
	case "A", "X":
		return ROCK
	case "B", "Y":
		return PAPER
	case "C", "Z":
		return SCISSORS
	default:
		panic("invalid code for shape")
	}
}

func ParseResult(s string) Result {
	switch s {
	case "X":
		return LOSE
	case "Y":
		return DRAW
	case "Z":
		return WIN
	default:
		panic("invalid result code")
	}
}

func (s Shape) Score() int {
	switch s {
	case ROCK:
		return 1
	case PAPER:
		return 2
	case SCISSORS:
		return 3
	default:
		panic("unknown shape")
	}
}

func (r Result) Score() int {
	switch r {
	case WIN:
		return 6
	case DRAW:
		return 3
	case LOSE:
		return 0
	default:
		panic("unknown result code")
	}
}

func (s Shape) Shoot(b Shape) int {
	if s == b {
		return 3
	}
	if s == ROCK && b == SCISSORS {
		return 6
	}
	if s == PAPER && b == ROCK {
		return 6
	}
	if s == SCISSORS && b == PAPER {
		return 6
	}
	return 0
}

func (s Shape) Result(r Result) Shape {
	if r == DRAW {
		return s
	}

	if s == ROCK {
		switch r {
		case WIN:
			return PAPER
		case LOSE:
			return SCISSORS
		}
	}

	if s == PAPER {
		switch r {
		case WIN:
			return SCISSORS
		case LOSE:
			return ROCK
		}
	}

	if s == SCISSORS {
		switch r {
		case WIN:
			return ROCK
		case LOSE:
			return PAPER
		}
	}

	panic("shape not supported")
}

func main() {
	data, err := os.ReadFile("2.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(data), "\n")
	score := 0
	for _, line := range lines {
		parts := strings.Split(line, " ")
		a := ParseShape(parts[0])
		b := ParseResult(parts[1])
		score += b.Score() + a.Result(b).Score()
	}
	log.Println(score)
}
