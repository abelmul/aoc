package main

import (
	"fmt"
	"io"
	"math"
)

/*
--- Day 17: No Such Thing as Too Much ---

The elves bought too much eggnog again - 150 liters this time. To fit it all into your refrigerator, you'll need to move it into smaller containers. You take an inventory of the capacities of the available containers.

For example, suppose you have containers of size 20, 15, 10, 5, and 5 liters. If you need to store 25 liters, there are four ways to do it:

    15 and 10
    20 and 5 (the first 5)
    20 and 5 (the second 5)
    15, 5, and 5

Filling all containers entirely, how many different combinations of containers can exactly fit all 150 liters of eggnog?

--- Part Two ---

While playing with all the containers in the kitchen, another load of eggnog arrives! The shipping and receiving department is requesting as many containers as you can spare.

Find the minimum number of containers that can exactly fit all 150 liters of eggnog. How many different ways can you fill that number of containers and still hold exactly 150 litres?

In the example above, the minimum number of containers was two. There were three ways to use that many containers, and so the answer there would be 3.
*/

func main() {
	containers := []int{}

	for {
		var i int
		_, err := fmt.Scanf("%d\n", &i)

		if err == io.EOF {
			break
		}

		containers = append(containers, i)
	}

	size := len(containers)
	container_count := map[int]int{}

	var backtrack func(int, int, int) int

	backtrack = func(i, remaining, count int) int {
		ans := 0

		if remaining == 0 {
			container_count[count]++
			return 1
		} else if remaining < 0 {
			return 0
		}

		if i < size {
			ans += backtrack(i+1, remaining-containers[i], count+1)
			ans += backtrack(i+1, remaining, count)
		}

		return ans
	}

	fmt.Println(backtrack(0, 150, 0))

	m := math.MaxInt32

	for k := range container_count {
		m = min(m, k)
	}

	fmt.Println(container_count[m])
}
