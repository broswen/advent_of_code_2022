package main

import (
	"log"
	"os"
	"strings"
)

type Rucksack struct {
	One map[rune]int
	Two map[rune]int
	All map[rune]int
}

func (rs Rucksack) String() string {
	s := strings.Builder{}
	for k, _ := range rs.One {
		s.WriteRune(k)
	}
	s.WriteString(" | ")
	for k, _ := range rs.Two {
		s.WriteRune(k)
	}
	return s.String()
}

func (rs Rucksack) Mismatch() rune {
	for k, _ := range rs.One {
		if _, ok := rs.Two[k]; ok {
			return k
		}
	}
	panic("mismatch not found")
}

func ParseRucksack(s string) Rucksack {
	r := Rucksack{
		One: make(map[rune]int),
		Two: make(map[rune]int),
		All: make(map[rune]int),
	}
	size := len(s) / 2
	for i, v := range s {
		if i < size {
			r.One[v] += 1
		} else {
			r.Two[v] += 1
		}
		r.All[v] += 1
	}
	return r
}

func Priority(i rune) int {
	if i <= 'Z' {
		return int(i) - 'A' + 27
	}
	return int(i) - 'a' + 1
}

func Common(lists []map[rune]int) rune {
	seen := make(map[rune]int)
	for _, list := range lists {
		for k := range list {
			seen[k] += 1
		}
	}
	for k, v := range seen {
		if v == len(lists) {
			return k
		}
	}
	panic("no common items")
}

func main() {
	data, err := os.ReadFile("3.txt")
	//data, err := os.ReadFile("example.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(data), "\n")
	rucksacks := make([]Rucksack, 0)
	for _, line := range lines {
		rucksacks = append(rucksacks, ParseRucksack(line))
	}
	totalPriority := 0
	groups := make([][]map[rune]int, 0)
	temp := make([]map[rune]int, 0)
	for _, rs := range rucksacks {
		temp = append(temp, rs.All)
		if len(temp) == 3 {
			groups = append(groups, temp)
			temp = make([]map[rune]int, 0)
		}
	}
	for _, group := range groups {
		common := Common(group)
		totalPriority += Priority(common)
	}
	log.Println(totalPriority)
}
