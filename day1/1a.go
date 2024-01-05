package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func wordToNumber(word string) string {
	wordToNumberMap := map[string]string{
		"zero":  "0",
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	lowercaseWord := strings.ToLower(word)

	if val, ok := wordToNumberMap[lowercaseWord]; ok {
		return val
	}

	return ""
}

func rev(word string) string {
	res := ""
	for i := len(word) - 1; i >= 0; i-- {
		res += string(word[i])
	}
	return res
}

func main() {
	var allDigits = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	checkDigits := func(let1 string, ind int, digits []string) []string {
		var newDigits []string
		for _, digit := range digits {
			if ind < len(digit) && string(digit[ind]) == let1 {
				newDigits = append(newDigits, digit)
			}
		}
		return newDigits
	}

	b, err := os.Open("./day1.txt")
	check(err)

	content, err := io.ReadAll(b)
	check(err)

	ans := 0

	for _, line := range strings.Split(string(content), "\n") {
		lineAns := ""
		i := 0
		end := len(line)

		for i < end {
			letter, err := strconv.Atoi(string(line[i]))
			if err != nil {
				j := 1
				wordOptions := checkDigits(string(line[i]), 0, allDigits)
				for len(wordOptions) > 1 && i+j < end {
					wordOptions = checkDigits(string(line[i+j]), j, wordOptions)
					j += 1
				}

				if len(wordOptions) == 1 {
					if i+len(wordOptions[0]) <= end && string(line[i:i+len(wordOptions[0])]) == wordOptions[0] {
						lineAns += wordToNumber(wordOptions[0])
						break
					} else {
						i += 1
					}
				} else {
					i += 1
					continue
				}
			} else {
				lineAns = lineAns + fmt.Sprint(letter)
				break
			}
		}

		var reversedDigits []string
		for _, dig := range allDigits {
			reversedDigits = append(reversedDigits, rev(dig))
		}

		i = end - 1
		for i >= 0 {
			letter, err := strconv.Atoi(string(line[i]))
			if err != nil {
				j := 1
				wordOptions := checkDigits(string(line[i]), 0, reversedDigits)
				for len(wordOptions) > 1 && i-j >= 0 {
					wordOptions = checkDigits(string(line[i-j]), j, wordOptions)
					j += 1
				}

				if len(wordOptions) == 1 {
					if i-len(wordOptions[0])+1 > 0 && rev(wordOptions[0]) == string(line[i-len(wordOptions[0])+1:i+1]) {
						lineAns += wordToNumber(rev(wordOptions[0]))
						break
					} else {
						i -= 1
						continue
					}
				} else {
					i -= 1
					continue
				}

			} else {
				lineAns = lineAns + fmt.Sprint(letter)
				break
			}
		}

		ansInt, err := strconv.Atoi(lineAns)
		check(err)
		ans = ans + ansInt
	}

	fmt.Println(ans)
}
