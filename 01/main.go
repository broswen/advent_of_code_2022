package main

import (
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Elf struct {
	Food          []int64
	TotalCalories int64
}

type Elves []Elf

func (e Elves) Len() int {
	return len(e)
}

func (e Elves) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

func (e Elves) Less(i, j int) bool {
	return e[i].TotalCalories > e[j].TotalCalories
}

func main() {
	elves := make(Elves, 0)
	data, err := os.ReadFile("1.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(data), "\n")
	currentElf := Elf{[]int64{}, 0}
	var currentTotalCalories int64 = 0
	var maxTotalCalories int64 = 0
	for _, v := range lines {
		if v == "" {
			currentElf.TotalCalories = currentTotalCalories
			elves = append(elves, currentElf)
			currentElf = Elf{[]int64{}, 0}
			if currentTotalCalories > maxTotalCalories {
				maxTotalCalories = currentTotalCalories
			}
			currentTotalCalories = 0
			continue
		}
		calories, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			panic(err)
		}
		currentElf.Food = append(currentElf.Food, calories)
		currentTotalCalories += calories
	}
	sort.Sort(elves)
	log.Println("max", maxTotalCalories)
	log.Println(elves[0].TotalCalories + elves[1].TotalCalories + elves[2].TotalCalories)
}
