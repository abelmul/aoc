package main

import (
	"fmt"
	"io"
)

/**
Santa needs help figuring out which strings in his text file are naughty or nice.

A nice string is one with all of the following properties:

    It contains at least three vowels (aeiou only), like aei, xazegov, or aeiouaeiouaeiou.
    It contains at least one letter that appears twice in a row, like xx, abcdde (dd), or aabbccdd (aa, bb, cc, or dd).
    It does not contain the strings ab, cd, pq, or xy, even if they are part of one of the other requirements.

For example:

    ugknbfddgicrmopn is nice because it has at least three vowels (u...i...o...), a double letter (...dd...), and none of the disallowed substrings.
    aaa is nice because it has at least three vowels and a double letter, even though the letters used by different rules overlap.
    jchzalrnumimnmhp is naughty because it has no double letter.
    haegwjzuvuyypxyu is naughty because it contains the string xy.
    dvszwmarrgswjxmb is naughty because it contains only one vowel.

How many strings are nice?

--- Part Two ---

Realizing the error of his ways, Santa has switched to a better model of determining whether a string is naughty or nice. None of the old rules apply, as they are all clearly ridiculous.

Now, a nice string is one with all of the following properties:

    It contains a pair of any two letters that appears at least twice in the string without overlapping, like xyxy (xy) or aabcdefgaa (aa), but not like aaa (aa, but it overlaps).
    It contains at least one letter which repeats with exactly one letter between them, like xyx, abcdefeghi (efe), or even aaa.

For example:

    qjhvhtzxzqqjkmpb is nice because is has a pair that appears twice (qj) and a letter that repeats with exactly one letter between them (zxz).
    xxyxx is nice because it has a pair that appears twice and a letter that repeats with one between, even though the letters used by each rule overlap.
    uurcxstgmygtbstg is naughty because it has a pair (tg) but no repeat with a single letter between them.
    ieodomkazucvgmuy is naughty because it has a repeating letter with one between (odo), but no pair that appears twice.

How many strings are nice under these new rules?
*/

func is_good_string_1(line string) int {
	prev_char := '%'

	found_vowels := 0
	found_double_letters := 0

	for _, char := range line {
		switch char {
		case 'a', 'e', 'i', 'o', 'u':
			found_vowels++
		}

		if char == prev_char {
			found_double_letters++
		}

		if (prev_char == 'a' && char == 'b') || (prev_char == 'c' && char == 'd') || (prev_char == 'p' && char == 'q') || (prev_char == 'x' && char == 'y') {
			return 0
		}

		prev_char = char
	}

	if found_vowels > 2 && found_double_letters > 0 {
		return 1
	}

	return 0
}

func is_good_string_2(line string) int {
	prevPrevChar := line[0]
	twoCharMap := map[string]int{}

	repeating_non_overlapping_pair := false
	repeating_char_with_one_in_between_letter := false

	// add one space to line for the sliding window to work
	line = line + " "

	for i := 2; i < len(line); i++ {
		if v, ok := twoCharMap[line[i-2:i]]; ok && i > v+3 {
			repeating_non_overlapping_pair = true
		} else {
			twoCharMap[line[i-2:i]] = i - 2
		}

		if line[i] == prevPrevChar {
			repeating_char_with_one_in_between_letter = true
		} else {
			prevPrevChar = line[i-1]
		}

		if repeating_char_with_one_in_between_letter && repeating_non_overlapping_pair {
			return 1
		}
	}

	return 0
}

func main() {
	good_strings_1 := 0
	good_strings_2 := 0

	for {
		var line string

		_, err := fmt.Scanln(&line)
		if err == io.EOF {
			break
		}

		good_strings_1 += is_good_string_1(line)
		good_strings_2 += is_good_string_2(line)
	}

	fmt.Println(good_strings_1)
	fmt.Println(good_strings_2)
}
