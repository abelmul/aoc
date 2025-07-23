package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
)

/**
--- Day 13: Knights of the Dinner Table ---

In years past, the holiday feast with your family hasn't gone so well. Not everyone gets along! This year, you resolve, will be different. You're going to find the optimal seating arrangement and avoid all those awkward conversations.

You start by writing up a list of everyone invited and the amount their happiness would increase or decrease if they were to find themselves sitting next to each other person. You have a circular table that will be just big enough to fit everyone comfortably, and so each person will have exactly two neighbors.

For example, suppose you have only four attendees planned, and you calculate their potential happiness as follows:

Alice would gain 54 happiness units by sitting next to Bob.
Alice would lose 79 happiness units by sitting next to Carol.
Alice would lose 2 happiness units by sitting next to David.
Bob would gain 83 happiness units by sitting next to Alice.
Bob would lose 7 happiness units by sitting next to Carol.
Bob would lose 63 happiness units by sitting next to David.
Carol would lose 62 happiness units by sitting next to Alice.
Carol would gain 60 happiness units by sitting next to Bob.
Carol would gain 55 happiness units by sitting next to David.
David would gain 46 happiness units by sitting next to Alice.
David would lose 7 happiness units by sitting next to Bob.
David would gain 41 happiness units by sitting next to Carol.

Then, if you seat Alice next to David, Alice would lose 2 happiness units (because David talks so much), but David would gain 46 happiness units (because Alice is such a good listener), for a total change of 44.

If you continue around the table, you could then seat Bob next to Alice (Bob gains 83, Alice gains 54). Finally, seat Carol, who sits next to Bob (Carol gains 60, Bob loses 7) and David (Carol gains 55, David gains 41). The arrangement looks like this:

     +41 +46
+55   David    -2
Carol       Alice
+60    Bob    +54
     -7  +83

After trying every other seating arrangement in this hypothetical scenario, you find that this one is the most optimal, with a total change in happiness of 330.

What is the total change in happiness for the optimal seating arrangement of the actual guest list?

--- Part Two ---

In all the commotion, you realize that you forgot to seat yourself. At this point, you're pretty apathetic toward the whole thing, and your happiness wouldn't really go up or down regardless of who you sit next to. You assume everyone else would be just as ambivalent about sitting next to you, too.

So, add yourself to the list, and give all happiness relationships that involve you a score of 0.

What is the total change in happiness for the optimal seating arrangement that actually includes yourself?
*/

func permutations(a []string) [][]string {
	res := [][]string{}
	size := len(a)

	var backtrack func(int, []string)

	backtrack = func(l int, cur []string) {
		if l >= size {
			cpy := make([]string, size)
			copy(cpy, cur)

			res = append(res, cpy)
			return
		}

		for i := l; i < size; i++ {
			cur[l], cur[i] = cur[i], cur[l]

			backtrack(l+1, cur)

			cur[l], cur[i] = cur[i], cur[l]
		}
	}

	backtrack(0, a)

	return res
}

func maxHappiness(persons []string, happiness map[string]map[string]int) int {
	persons_permutations := permutations(persons)

	max_happiness := math.MinInt32

	for _, permutation := range persons_permutations {
		permutation = append(permutation, permutation[0])

		current_happiness := 0

		for i := 1; i < len(permutation); i++ {
			current_happiness += happiness[permutation[i-1]][permutation[i]] + happiness[permutation[i]][permutation[i-1]]
		}

		max_happiness = max(max_happiness, current_happiness)
	}

	return max_happiness
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	happiness := map[string]map[string]int{}

	persons := []string{}
	encountered_person := map[string]bool{}

	for {

		line, err := reader.ReadString('\n')

		if err == io.EOF {
			break
		}

		var person, gain_lose, next_person string
		var amount int

		fmt.Sscanf(line, "%s would %s %d happiness units by sitting next to %s", &person, &gain_lose, &amount, &next_person)

		next_person = next_person[:len(next_person)-1]

		if _, ok := happiness[person]; !ok {
			happiness[person] = map[string]int{}
		}

		if gain_lose == "gain" {
			happiness[person][next_person] = amount
		} else {
			happiness[person][next_person] = -amount
		}

		if _, ok := encountered_person[person]; !ok {
			encountered_person[person] = true
			persons = append(persons, person)
		}

		if _, ok := encountered_person[next_person]; !ok {
			encountered_person[next_person] = true
			persons = append(persons, next_person)
		}
	}

	fmt.Println(maxHappiness(persons, happiness))
	fmt.Println(maxHappiness(append(persons, "me"), happiness))
}
