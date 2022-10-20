package main

import (
	"testing"
)

func Test_getRandomToy(t *testing.T) {
	counter := make(map[Toy]int)
	for i := 0; i < 100000; i++ {
		toy := randomToy(toys)
		counter[toy]++
	}
	for toy, count := range counter {
		t.Logf("%s: %d", toy.Name, count)
	}
}
