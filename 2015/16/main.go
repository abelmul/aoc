package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

/*
--- Day 16: Aunt Sue ---

Your Aunt Sue has given you a wonderful gift, and you'd like to send her a thank you card. However, there's a small problem: she signed it "From, Aunt Sue".

You have 500 Aunts named "Sue".

So, to avoid sending the card to the wrong person, you need to figure out which Aunt Sue (which you conveniently number 1 to 500, for sanity) gave you the gift. You open the present and, as luck would have it, good ol' Aunt Sue got you a My First Crime Scene Analysis Machine! Just what you wanted. Or needed, as the case may be.

The My First Crime Scene Analysis Machine (MFCSAM for short) can detect a few specific compounds in a given sample, as well as how many distinct kinds of those compounds there are. According to the instructions, these are what the MFCSAM can detect:

    children, by human DNA age analysis.
    cats. It doesn't differentiate individual breeds.
    Several seemingly random breeds of dog: samoyeds, pomeranians, akitas, and vizslas.
    goldfish. No other kinds of fish.
    trees, all in one group.
    cars, presumably by exhaust or gasoline or something.
    perfumes, which is handy, since many of your Aunts Sue wear a few kinds.

In fact, many of your Aunts Sue have many of these. You put the wrapping from the gift into the MFCSAM. It beeps inquisitively at you a few times and then prints out a message on ticker tape:

children: 3
cats: 7
samoyeds: 2
pomeranians: 3
akitas: 0
vizslas: 0
goldfish: 5
trees: 3
cars: 2
perfumes: 1

You make a list of the things you can remember about each Aunt Sue. Things missing from your list aren't zero - you simply don't remember the value.

What is the number of the Sue that got you the gift?

--- Part Two ---

As you're about to send the thank you note, something in the MFCSAM's instructions catches your eye. Apparently, it has an outdated retroencabulator, and so the output from the machine isn't exact values - some of them indicate ranges.

In particular, the cats and trees readings indicates that there are greater than that many (due to the unpredictable nuclear decay of cat dander and tree pollen), while the pomeranians and goldfish readings indicate that there are fewer than that many (due to the modial interaction of magnetoreluctance).

What is the number of the real Aunt Sue?
*/

func exactInt(s []byte) int {
	v, _ := strconv.Atoi(string(s))

	return v
}

func part2_test(match string, num int, mfcsam_output map[string]int) bool {
	switch match {
	case "cats", "trees":
		return mfcsam_output[match] < num
	case "pomeranians", "goldfish":
		return mfcsam_output[match] > num
	}

	return mfcsam_output[match] == num
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	regex := regexp.MustCompile("Sue ([0-9]+): ([a-z]+): ([0-9]+), ([a-z]+): ([0-9]+), ([a-z]+): ([0-9]+)")
	mfcsam_output := map[string]int{
		"children":    3,
		"cats":        7,
		"samoyeds":    2,
		"pomeranians": 3,
		"akitas":      0,
		"vizslas":     0,
		"goldfish":    5,
		"trees":       3,
		"cars":        2,
		"perfumes":    1,
	}

	ans := 0
	ans_2 := 0

	for {
		line, err := reader.ReadString('\n')

		if err == io.EOF {
			break
		}

		line = strings.TrimSpace(line)

		matches := regex.FindSubmatch([]byte(line))

		if match, num := string(matches[2]), exactInt(matches[3]); mfcsam_output[match] == num {
			if match, num := string(matches[4]), exactInt(matches[5]); mfcsam_output[match] == num {
				if match, num := string(matches[6]), exactInt(matches[7]); mfcsam_output[match] == num {
					ans = exactInt(matches[1])
				}
			}
		}

		if match, num := string(matches[2]), exactInt(matches[3]); part2_test(match, num, mfcsam_output) {
			if match, num := string(matches[4]), exactInt(matches[5]); part2_test(match, num, mfcsam_output) {
				if match, num := string(matches[6]), exactInt(matches[7]); part2_test(match, num, mfcsam_output) {
					ans_2 = exactInt(matches[1])
				}
			}
		}
	}

	fmt.Println(ans)
	fmt.Println(ans_2)
}
