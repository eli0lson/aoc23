package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/eli0lson/aoc23/utilities"
)

func B() {
	b, err := os.Open("./day2.txt")
	utilities.Check(err)

	content, err := io.ReadAll(b)
	utilities.Check(err)

	ans := 0

	for _, line := range strings.Split(string(content), "\n") {
		colorToMaxMap := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}
		firstSplit := strings.Split(line, ":")
		if len(firstSplit) > 0 {
			sets := strings.Split(firstSplit[1], ";")
			for _, set := range sets {
				pulls := strings.Split(set, ",")
				for _, pull := range pulls {
					itemStrings := strings.Split(strings.Trim(pull, " "), " ")
					if max, ok := colorToMaxMap[itemStrings[1]]; ok {
						amt, err := strconv.Atoi(itemStrings[0])
						utilities.Check(err)
						colorToMaxMap[itemStrings[1]] = int(math.Max(float64(max), float64(amt)))
					}
				}
			}
			power := 1
			for _, amt := range colorToMaxMap {
				power = power * amt
			}
			ans += power
		}
	}

	// fmt.Println(colorToMaxMap)
	fmt.Println(ans)
}
