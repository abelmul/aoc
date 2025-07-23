package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
)

/**
--- Day 9: All in a Single Night ---

Every year, Santa manages to deliver all of his presents in a single night.

This year, however, he has some new locations to visit; his elves have provided him the distances between every pair of locations. He can start and end at any two (different) locations he wants, but he must visit each location exactly once. What is the shortest distance he can travel to achieve this?

For example, given the following distances:

London to Dublin = 464
London to Belfast = 518
Dublin to Belfast = 141

The possible routes are therefore:

Dublin -> London -> Belfast = 982
London -> Dublin -> Belfast = 605
London -> Belfast -> Dublin = 659
Dublin -> Belfast -> London = 659
Belfast -> Dublin -> London = 605
Belfast -> London -> Dublin = 982

The shortest of these is London -> Dublin -> Belfast = 605, and so the answer is 605 in this example.

What is the distance of the shortest route?

--- Part Two ---

The next year, just to show off, Santa decides to take the route with the longest distance instead.

He can still start and end at any two (different) locations he wants, and he still must visit each location exactly once.

For example, given the distances above, the longest route would be 982 via (for example) Dublin -> London -> Belfast.

What is the distance of the longest route?
*/

func nextLocationMin(locations map[string]map[string]int, current string, unexplored map[string]bool) string {
	nextCost := math.MaxInt32
	nextLocation := ""

	for k := range locations[current] {
		if _, ok := unexplored[k]; ok && nextCost > locations[current][k] {
			nextCost = locations[current][k]
			nextLocation = k
		}
	}

	return nextLocation
}

func nextLocationMax(locations map[string]map[string]int, current string, unexplored map[string]bool) string {
	nextCost := 0
	nextLocation := ""

	for k := range locations[current] {
		if _, ok := unexplored[k]; ok && nextCost < locations[current][k] {
			nextCost = locations[current][k]
			nextLocation = k
		}
	}

	return nextLocation
}

func traveling_salesman(locations map[string]map[string]int) (int, int) {
	min_cost := math.MaxInt32
	max_cost := 0

	for l := range locations {
		unexplored_min := map[string]bool{}
		unexplored_max := map[string]bool{}

		for k := range locations {
			unexplored_min[k] = true
			unexplored_max[k] = true
		}

		current_min := 0
		current_max := 0

		current := l
		for len(unexplored_min) > 1 {
			delete(unexplored_min, current)
			next := nextLocationMin(locations, current, unexplored_min)

			current_min += locations[current][next]
			current = next
		}

		current = l
		for len(unexplored_max) > 1 {
			delete(unexplored_max, current)
			next := nextLocationMax(locations, current, unexplored_max)

			current_max += locations[current][next]
			current = next
		}

		min_cost = min(min_cost, current_min)
		max_cost = max(max_cost, current_max)
	}

	return min_cost, max_cost
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	locations := map[string]map[string]int{}

	for {
		line, err := reader.ReadString('\n')

		if err == io.EOF {
			break
		}

		var location1, location2 string
		var distance int

		fmt.Sscanf(line, "%s to %s = %d", &location1, &location2, &distance)

		if _, ok := locations[location1]; !ok {
			locations[location1] = map[string]int{}
		}

		if _, ok := locations[location2]; !ok {
			locations[location2] = map[string]int{}
		}

		locations[location1][location2] = distance
		locations[location2][location1] = distance
	}

	fmt.Println(traveling_salesman(locations))
}
