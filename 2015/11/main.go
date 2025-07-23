package main

import (
	"fmt"
)

/**
--- Day 11: Corporate Policy ---

Santa's previous password expired, and he needs help choosing a new one.

To help him remember his new password after the old one expires, Santa has devised a method of coming up with a password based on the previous one. Corporate policy dictates that passwords must be exactly eight lowercase letters (for security reasons), so he finds his new password by incrementing his old password string repeatedly until it is valid.

Incrementing is just like counting with numbers: xx, xy, xz, ya, yb, and so on. Increase the rightmost letter one step; if it was z, it wraps around to a, and repeat with the next letter to the left until one doesn't wrap around.

Unfortunately for Santa, a new Security-Elf recently started, and he has imposed some additional password requirements:

    Passwords must include one increasing straight of at least three letters, like abc, bcd, cde, and so on, up to xyz. They cannot skip letters; abd doesn't count.
    Passwords may not contain the letters i, o, or l, as these letters can be mistaken for other characters and are therefore confusing.
    Passwords must contain at least two different, non-overlapping pairs of letters, like aa, bb, or zz.

For example:

    hijklmmn meets the first requirement (because it contains the straight hij) but fails the second requirement requirement (because it contains i and l).
    abbceffg meets the third requirement (because it repeats bb and ff) but fails the first requirement.
    abbcegjk fails the third requirement, because it only has one double letter (bb).
    The next password after abcdefgh is abcdffaa.
    The next password after ghijklmn is ghjaabcc, because you eventually skip all the passwords that start with ghi..., since i is not allowed.

Given Santa's current password (your puzzle input), what should his next password be?

--- Part Two ---

Santa's password expired again. What's the next one?
*/

func incrementString(s string) string {
	chars := []byte(s)

	for i := len(chars) - 1; i >= 0; i-- {
		if chars[i] == 'z' {
			chars[i] = 'a'
		} else {
			chars[i] = chars[i] + 1
			break
		}
	}
	return string(chars)
}

func isValid(s string) bool {
	var firstDoubleLocation, secondDoubleLocation int
	var rule3 bool

	firstDoubleLocation = -1
	secondDoubleLocation = -1

	var prevChar, prevPrevChar byte

	for i := range s {
		if s[i] == 'i' || s[i] == 'o' || s[i] == 'l' {
			return false
		}

		if s[i] == prevChar {
			if firstDoubleLocation == -1 {
				firstDoubleLocation = i - 1
			} else if i > firstDoubleLocation+2 {
				secondDoubleLocation = i
			}
		}

		if prevChar == s[i]-1 && prevPrevChar == s[i]-2 {
			rule3 = true
		}

		if firstDoubleLocation != -1 && secondDoubleLocation != -1 && rule3 {
			return true
		}

		prevPrevChar = prevChar
		prevChar = s[i]
	}

	return false
}

func main() {
	password := "hxbxwxba"

	for !isValid(password) {
		password = incrementString(password)
	}
	fmt.Println(password)

	password = incrementString(password)
	for !isValid(password) {
		password = incrementString(password)
	}
	fmt.Println(password)
}
