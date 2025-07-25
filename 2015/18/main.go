package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

/*
--- Day 18: Like a GIF For Your Yard ---

After the million lights incident, the fire code has gotten stricter: now, at most ten thousand lights are allowed. You arrange them in a 100x100 grid.

Never one to let you down, Santa again mails you instructions on the ideal lighting configuration. With so few lights, he says, you'll have to resort to animation.

Start by setting your lights to the included initial configuration (your puzzle input). A # means "on", and a . means "off".

Then, animate your grid in steps, where each step decides the next configuration based on the current one. Each light's next state (either on or off) depends on its current state and the current states of the eight lights adjacent to it (including diagonals). Lights on the edge of the grid might have fewer than eight neighbors; the missing ones always count as "off".

For example, in a simplified 6x6 grid, the light marked A has the neighbors numbered 1 through 8, and the light marked B, which is on an edge, only has the neighbors marked 1 through 5:

1B5...
234...
......
..123.
..8A4.
..765.

The state a light should have next is based on its current state (on or off) plus the number of neighbors that are on:

    A light which is on stays on when 2 or 3 neighbors are on, and turns off otherwise.
    A light which is off turns on if exactly 3 neighbors are on, and stays off otherwise.

All of the lights update simultaneously; they all consider the same current state before moving to the next.

Here's a few steps from an example configuration of another 6x6 grid:

Initial state:
.#.#.#
...##.
#....#
..#...
#.#..#
####..

After 1 step:
..##..
..##.#
...##.
......
#.....
#.##..

After 2 steps:
..###.
......
..###.
......
.#....
.#....

After 3 steps:
...#..
......
...#..
..##..
......
......

After 4 steps:
......
......
..##..
..##..
......
......

After 4 steps, this example has four lights on.

In your grid of 100x100 lights, given your initial configuration, how many lights are on after 100 steps?

--- Part Two ---

You flip the instructions over; Santa goes on to point out that this is all just an implementation of Conway's Game of Life. At least, it was, until you notice that something's wrong with the grid of lights you bought: four lights, one in each corner, are stuck on and can't be turned off. The example above will actually run like this:

Initial state:
##.#.#
...##.
#....#
..#...
#.#..#
####.#

After 1 step:
#.##.#
####.#
...##.
......
#...#.
#.####

After 2 steps:
#..#.#
#....#
.#.##.
...##.
.#..##
##.###

After 3 steps:
#...##
####.#
..##.#
......
##....
####.#

After 4 steps:
#.####
#....#
...#..
.##...
#.....
#.#..#

After 5 steps:
##.###
.##..#
.##...
.##...
#.#...
##...#

After 5 steps, this example now has 17 lights on.

In your grid of 100x100 lights, given your initial configuration, but with the four corners always in the on state, how many lights are on after 100 steps?
*/

func count_on_neighbours(lights [][]bool, i, j int) int {
	count := 0

	if lights[i-1][j-1] {
		count++
	}

	if lights[i-1][j] {
		count++
	}

	if lights[i-1][j+1] {
		count++
	}

	if lights[i][j-1] {
		count++
	}

	if lights[i][j+1] {
		count++
	}

	if lights[i+1][j-1] {
		count++
	}

	if lights[i+1][j] {
		count++
	}

	if lights[i+1][j+1] {
		count++
	}

	return count
}

func next_state(lights [][]bool, i, j int) bool {
	if on_neighbours := count_on_neighbours(lights, i, j); (lights[i][j] && on_neighbours == 2 || on_neighbours == 3) || (!lights[i][j] && on_neighbours == 3) {
		return true
	}

	return false
}

func new__state() [][]bool {
	new_lights := make([][]bool, 102)
	for i := range 102 {
		new_lights[i] = make([]bool, 102)
	}

	return new_lights
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	lights := new__state()

	i := 0

	for i < 100 {
		line, err := reader.ReadString('\n')

		if err == io.EOF {
			break
		}

		line = strings.TrimSpace(line)

		for j, v := range line {
			switch v {
			case '#':
				lights[i+1][j+1] = true
			}
		}

		i++
	}

	part2_lights := new__state()
	for i := range lights {
		copy(part2_lights[i], lights[i])
	}

	part2_lights[1][1] = true
	part2_lights[1][100] = true
	part2_lights[100][1] = true
	part2_lights[100][100] = true

	for range 100 {
		new_lights := new__state()
		new_part2_lights := new__state()

		for i := 1; i < 101; i++ {
			for j := 1; j < 101; j++ {
				new_lights[i][j] = next_state(lights, i, j)
				new_part2_lights[i][j] = next_state(part2_lights, i, j)
			}
		}

		lights = new_lights

		new_part2_lights[1][1] = true
		new_part2_lights[1][100] = true
		new_part2_lights[100][1] = true
		new_part2_lights[100][100] = true
		part2_lights = new_part2_lights
	}

	count := 0
	count_part2 := 0

	for i := range lights {
		for j := range lights[i] {
			if lights[i][j] {
				count++
			}
			if part2_lights[i][j] {
				count_part2++
			}
		}
	}

	fmt.Println(count)
	fmt.Println(count_part2)
}
