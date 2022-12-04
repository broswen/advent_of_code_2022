package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	Start int
	End   int
}

func (r Range) String() string {
	return fmt.Sprintf("%d - %d", r.Start, r.End)
}

func (r Range) Contains(b Range) bool {
	return b.Start >= r.Start && b.End <= r.End
}

func (r Range) Overlaps(b Range) bool {
	return b.Start >= r.Start && b.Start <= r.End || b.End >= r.Start && b.End <= r.End
}

func ParseRange(s string) Range {
	parts := strings.Split(s, "-")
	start, err := strconv.ParseInt(parts[0], 10, 64)
	if err != nil {
		panic(err)
	}
	end, err := strconv.ParseInt(parts[1], 10, 64)
	if err != nil {
		panic(err)
	}
	return Range{
		Start: int(start),
		End:   int(end),
	}
}

func main() {
	data, err := os.ReadFile("4.txt")
	//data, err := os.ReadFile("example.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(data), "\n")
	overlapCount := 0
	containCount := 0
	for _, line := range lines {
		parts := strings.Split(line, ",")
		range1 := ParseRange(parts[0])
		range2 := ParseRange(parts[1])
		log.Println(range1, range2)
		if range1.Contains(range2) || range2.Contains(range1) {
			containCount++
		}
		if range1.Overlaps(range2) || range2.Overlaps(range1) {
			overlapCount++
		}
	}
	log.Println("contained", containCount)
	log.Println("overlapped", overlapCount)
}
