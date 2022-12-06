package main

import (
	"log"
	"os"
)

func unique(s string) bool {
	seen := make(map[string]bool)
	for _, c := range s {
		if _, ok := seen[string(c)]; !ok {
			seen[string(c)] = true
		} else {
			return false
		}
	}
	return true
}

func findUnique(s string, size int) int {
	for i := range s {
		if unique(s[i : i+size]) {
			return i + size
		}
	}
	return -1
}

func main() {
	data, err := os.ReadFile("6.txt")
	//data, err := os.ReadFile("example.txt")
	if err != nil {
		panic(err)
	}
	log.Println(findUnique(string(data), 14))
}
