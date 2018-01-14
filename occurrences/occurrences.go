package main

import (
	"flag"
	"fmt"
	"strings"
)

var mininum *uint

func main() {
	var s [3]string
	s[0] = "1 b 3 1 a a b"
	s[1] = "11 a 1 1 1 1 a a"
	s[2] = "-1 b 1 -4 a 1"

	minimum := flag.Uint("min", 1, "Minimum number of occurrences")
	flag.Parse()

	counts := make(map[string]uint)

	for _, line := range s {
		values := strings.Fields(line)
		for _, item := range values {
			indexItem(counts, item)
		}
	}

	for key, value := range counts {
		if value >= *minimum {
			fmt.Printf("%s -> %d\n", key, value)
		}
	}
}

func indexItem(index map[string]uint, item string) {
	_, present := index[item]
	if present {
		index[item]++
	} else {
		index[item] = 1
	}
}
