package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/eli0lson/aoc23/utilities"
)

func A() {
	b, err := os.Open("./day2.txt")
	utilities.Check(err)

	content, err := io.ReadAll(b)
	utilities.Check(err)

	colorToAmtMap := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	ans := 0

	for _, line := range strings.Split(string(content), "\n") {
		firstSplit := strings.Split(line, ":")
		possible := true
		if len(firstSplit) > 0 {
			id, err := strconv.Atoi(string(strings.Split(firstSplit[0], " ")[1]))
			utilities.Check(err)
			sets := strings.Split(firstSplit[1], ";")
			for _, set := range sets {
				pulls := strings.Split(set, ",")
				for _, pull := range pulls {
					itemStrings := strings.Split(strings.Trim(pull, " "), " ")
					if limit, ok := colorToAmtMap[itemStrings[1]]; ok {
						amt, err := strconv.Atoi(itemStrings[0])
						utilities.Check(err)
						possible = possible && (amt <= limit)
					}
				}
			}
			if possible {
				ans += id
			}
		}
	}

	fmt.Println(ans)
}
