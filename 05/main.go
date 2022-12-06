package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Instruction struct {
	Amount      int
	Source      int
	Destination int
}

func (i Instruction) String() string {
	return fmt.Sprintf("move %d from %d to %d", i.Amount, i.Source, i.Destination)
}

func parseInstruction(s string) Instruction {
	parts := strings.Split(s, " ")
	if len(parts) != 6 {
		panic("invalid instruction")
	}
	amount, err := strconv.ParseInt(parts[1], 10, 64)
	if err != nil {
		panic(err)
	}
	source, err := strconv.ParseInt(parts[3], 10, 64)
	if err != nil {
		panic(err)
	}
	destination, err := strconv.ParseInt(parts[5], 10, 64)
	if err != nil {
		panic(err)
	}
	return Instruction{
		Amount:      int(amount),
		Source:      int(source),
		Destination: int(destination),
	}
}

func parseInstructions(lines []string) []Instruction {
	instructions := make([]Instruction, 0)
	for _, line := range lines {
		instruction := parseInstruction(line)
		instructions = append(instructions, instruction)
	}
	return instructions
}

type Stack struct {
	Items []string
}

func (s Stack) String() string {
	return fmt.Sprintf("%v", s.Items)
}

func (s *Stack) Add(item string) {
	s.Items = append(s.Items, item)
}

func (s *Stack) Push(item string) {
	s.Items = append([]string{item}, s.Items...)
}

func (s *Stack) Remove() string {
	if len(s.Items) < 1 {
		return ""
	}
	item := s.Items[len(s.Items)-1]
	s.Items = s.Items[:len(s.Items)-1]
	return item
}

func newStack() *Stack {
	return &Stack{
		Items: make([]string, 0),
	}
}

func applyInstruction(ins Instruction, stacks []*Stack) {
	temp := newStack()
	for i := 0; i < ins.Amount; i++ {
		item := stacks[ins.Source-1].Remove()
		temp.Add(item)
	}
	for i := 0; i < ins.Amount; i++ {
		item := temp.Remove()
		stacks[ins.Destination-1].Add(item)
	}
}

func parseStacks(lines []string) []*Stack {
	places := []int{1, 5, 9, 13, 17, 21, 25, 29, 33}
	stacks := make([]*Stack, 0)
	for _ = range places {
		stacks = append(stacks, newStack())
	}
	for _, line := range lines {
		for i, place := range places {
			if string(line[place]) != " " {
				stacks[i].Push(string(line[place]))
				fmt.Println(i, string(line[place]))
			}
		}
	}
	return stacks
}

func main() {
	data, err := os.ReadFile("5.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(data), "\n")
	stacks := parseStacks(lines[:8])
	instructions := parseInstructions(lines[10:])
	for _, s := range stacks {
		fmt.Println(s)
	}
	fmt.Println("========")
	for _, ins := range instructions {
		applyInstruction(ins, stacks)
	}
	for _, s := range stacks {
		fmt.Println(s)
	}
}
