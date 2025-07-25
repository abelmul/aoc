package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

/**
--- Day 6: Probably a Fire Hazard ---

Because your neighbors keep defeating you in the holiday house decorating contest year after year, you've decided to deploy one million lights in a 1000x1000 grid.

Furthermore, because you've been especially nice this year, Santa has mailed you instructions on how to display the ideal lighting configuration.

Lights in your grid are numbered from 0 to 999 in each direction; the lights at each corner are at 0,0, 0,999, 999,999, and 999,0. The instructions include whether to turn on, turn off, or toggle various inclusive ranges given as coordinate pairs. Each coordinate pair represents opposite corners of a rectangle, inclusive; a coordinate pair like 0,0 through 2,2 therefore refers to 9 lights in a 3x3 square. The lights all start turned off.

To defeat your neighbors this year, all you have to do is set up your lights by doing the instructions Santa sent you in order.

For example:

    turn on 0,0 through 999,999 would turn on (or leave on) every light.
    toggle 0,0 through 999,0 would toggle the first line of 1000 lights, turning off the ones that were on, and turning on the ones that were off.
    turn off 499,499 through 500,500 would turn off (or leave off) the middle four lights.

After following the instructions, how many lights are lit?

--- Part Two ---

You just finish implementing your winning light pattern when you realize you mistranslated Santa's message from Ancient Nordic Elvish.

The light grid you bought actually has individual brightness controls; each light can have a brightness of zero or more. The lights all start at zero.

The phrase turn on actually means that you should increase the brightness of those lights by 1.

The phrase turn off actually means that you should decrease the brightness of those lights by 1, to a minimum of zero.

The phrase toggle actually means that you should increase the brightness of those lights by 2.

What is the total brightness of all lights combined after following Santa's instructions?
*/

func main() {
	reader := bufio.NewReader(os.Stdin)
	regex := regexp.MustCompile("(turn off|turn on|toggle) ([0-9]+),([0-9]+) through ([0-9]+),([0-9]+)")

	lights, lights_interval := make([][]bool, 1000), make([][]int, 1000)

	for i := range 1000 {
		lights[i] = make([]bool, 1000)
		lights_interval[i] = make([]int, 1000)
	}

	lit_count, brightness_count := 0, 0

	for {
		line, err := reader.ReadString('\n')

		if err == io.EOF {
			break
		}

		matches := regex.FindSubmatch([]byte(line))

		var operator int

		x, _ := strconv.Atoi(string(matches[2][:]))
		y, _ := strconv.Atoi(string(matches[3][:]))
		x_max, _ := strconv.Atoi(string(matches[4][:]))
		y_max, _ := strconv.Atoi(string(matches[5][:]))

		if o := string(matches[1][:]); o == "turn off" {
			operator = 0
		} else if o == "turn on" {
			operator = 1
		} else if o == "toggle" {
			operator = 2
		}

		y_min := y
		for x <= x_max {
			for y <= y_max {
				switch operator {
				case 0:
					if lights[x][y] {
						lit_count--
						lights[x][y] = false
					}

					if lights_interval[x][y] > 0 {
						brightness_count--
						lights_interval[x][y]--
					}
				case 1:
					if !lights[x][y] {
						lit_count++
						lights[x][y] = true
					}

					brightness_count++
					lights_interval[x][y]++
				case 2:
					if !lights[x][y] {
						lit_count++
					} else {
						lit_count--
					}
					lights[x][y] = !lights[x][y]

					brightness_count += 2
					lights_interval[x][y] += 2
				}

				y++
			}

			y = y_min
			x++
		}

	}

	fmt.Println(lit_count)
	fmt.Println(brightness_count)
}
