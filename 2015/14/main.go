package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

/*
*
--- Day 14: Reindeer Olympics ---

This year is the Reindeer Olympics! Reindeer can fly at high speeds, but must rest occasionally to recover their energy. Santa would like to know which of his reindeer is fastest, and so he has them race.

Reindeer can only either be flying (always at their top speed) or resting (not moving at all), and always spend whole seconds in either state.

For example, suppose you have the following Reindeer:

	Comet can fly 14 km/s for 10 seconds, but then must rest for 127 seconds.
	Dancer can fly 16 km/s for 11 seconds, but then must rest for 162 seconds.

After one second, Comet has gone 14 km, while Dancer has gone 16 km. After ten seconds, Comet has gone 140 km, while Dancer has gone 160 km. On the eleventh second, Comet begins resting (staying at 140 km), and Dancer continues on for a total distance of 176 km. On the 12th second, both reindeer are resting. They continue to rest until the 138th second, when Comet flies for another ten seconds. On the 174th second, Dancer flies for another 11 seconds.

In this example, after the 1000th second, both reindeer are resting, and Comet is in the lead at 1120 km (poor Dancer has only gotten 1056 km by that point). So, in this situation, Comet would win (if the race ended at 1000 seconds).

Given the descriptions of each reindeer (in your puzzle input), after exactly 2503 seconds, what distance has the winning reindeer traveled?

--- Part Two ---

Seeing how reindeer move in bursts, Santa decides he's not pleased with the old scoring system.

Instead, at the end of each second, he awards one point to the reindeer currently in the lead. (If there are multiple reindeer tied for the lead, they each get one point.) He keeps the traditional 2503 second time limit, of course, as doing otherwise would be entirely ridiculous.

Given the example reindeer from above, after the first second, Dancer is in the lead and gets one point. He stays in the lead until several seconds into Comet's second burst: after the 140th second, Comet pulls into the lead and gets his first point. Of course, since Dancer had been in the lead for the 139 seconds before that, he has accumulated 139 points by the 140th second.

After the 1000th second, Dancer has accumulated 689 points, while poor Comet, our old champion, only has 312. So, with the new scoring system, Dancer would win (if the race ended at 1000 seconds).

Again given the descriptions of each reindeer (in your puzzle input), after exactly 2503 seconds, how many points does the winning reindeer have?
*/
type DeerTracker struct {
	is_flying            bool
	in_current_state_for int
	current_distance     int
	bonus                int

	speed     int
	rest_for  int
	fly_until int
}

func calculateDistance(tracker map[string]DeerTracker, remaining_seconds int) {

	for remaining_seconds > 0 {
		current_max_distance := 0

		for k, v := range tracker {
			if v.is_flying {
				v.current_distance += v.speed

			}

			v.in_current_state_for++

			if v.is_flying && v.in_current_state_for >= v.fly_until {
				v.is_flying = false
				v.in_current_state_for = 0
			} else if v.in_current_state_for >= v.rest_for {
				v.is_flying = true
				v.in_current_state_for = 0
			}

			if current_max_distance < v.current_distance {
				current_max_distance = v.current_distance
			}

			tracker[k] = v
		}

		for k, v := range tracker {
			if v.current_distance == current_max_distance {
				v.bonus++
				tracker[k] = v
			}
		}

		remaining_seconds--
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	tracker := map[string]DeerTracker{}

	// max_distance := 0

	for {

		line, err := reader.ReadString('\n')

		if err == io.EOF {
			break
		}

		var name string
		var speed, fly_until, rest_for int

		fmt.Sscanf(line, "%s can fly %d km/s for %d seconds, but then must rest for %d seconds.", &name, &speed, &fly_until, &rest_for)

		tracker[name] = DeerTracker{
			is_flying: true,
			speed:     speed,
			fly_until: fly_until,
			rest_for:  rest_for,
		}
	}

	calculateDistance(tracker, 2503)

	max_distance := 0
	max_distance_bonus := 0

	for _, v := range tracker {
		max_distance = max(max_distance, v.current_distance)
		max_distance_bonus = max(max_distance_bonus, v.bonus)
	}

	fmt.Println(max_distance)
	fmt.Println(max_distance_bonus)
}
