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

/**
--- Day 7: Some Assembly Required ---

This year, Santa brought little Bobby Tables a set of wires and bitwise logic gates! Unfortunately, little Bobby is a little under the recommended age range, and he needs help assembling the circuit.

Each wire has an identifier (some lowercase letters) and can carry a 16-bit signal (a number from 0 to 65535). A signal is provided to each wire by a gate, another wire, or some specific value. Each wire can only get a signal from one source, but can provide its signal to multiple destinations. A gate provides no signal until all of its inputs have a signal.

The included instructions booklet describes how to connect the parts together: x AND y -> z means to connect wires x and y to an AND gate, and then connect its output to wire z.

For example:

    123 -> x means that the signal 123 is provided to wire x.
    x AND y -> z means that the bitwise AND of wire x and wire y is provided to wire z.
    p LSHIFT 2 -> q means that the value from wire p is left-shifted by 2 and then provided to wire q.
    NOT e -> f means that the bitwise complement of the value from wire e is provided to wire f.

Other possible gates include OR (bitwise OR) and RSHIFT (right-shift). If, for some reason, you'd like to emulate the circuit instead, almost all programming languages (for example, C, JavaScript, or Python) provide operators for these gates.

For example, here is a simple circuit:

123 -> x
456 -> y
x AND y -> d
x OR y -> e
x LSHIFT 2 -> f
y RSHIFT 2 -> g
NOT x -> h
NOT y -> i

After it is run, these are the signals on the wires:

d: 72
e: 507
f: 492
g: 114
h: 65412
i: 65079
x: 123
y: 456

In little Bobby's kit's instructions booklet (provided as your puzzle input), what signal is ultimately provided to wire a?

--- Part Two ---

Now, take the signal you got on wire a, override wire b to that signal, and reset the other wires (including wire a). What new signal is ultimately provided to wire a?
*/

func isNumber(value string) bool {
	return value[0] >= 48 && value[0] <= 57
}

func dfs(wires map[string]string, wire string, memo map[string]int) int {
	if _, ok := memo[wire]; !ok {
		operatorRegex := regexp.MustCompile("(AND|OR|NOT|LSHIFT|RSHIFT)")

		if isNumber(wire) {
			memo[wire], _ = strconv.Atoi(wire)
		} else if matches := operatorRegex.FindSubmatch([]byte(wires[wire])); matches == nil {
			memo[wire] = dfs(wires, wires[wire], memo)
		} else {
			var left, right string

			switch string(matches[1][:]) {
			case "AND":
				fmt.Sscanf(wires[wire], "%s AND %s", &left, &right)
				memo[wire] = dfs(wires, left, memo) & dfs(wires, right, memo)
			case "OR":
				fmt.Sscanf(wires[wire], "%s OR %s", &left, &right)
				memo[wire] = dfs(wires, left, memo) | dfs(wires, right, memo)
			case "NOT":
				fmt.Sscanf(wires[wire], "NOT %s", &left)
				memo[wire] = ^dfs(wires, left, memo)
			case "LSHIFT":
				fmt.Sscanf(wires[wire], "%s LSHIFT %s", &left, &right)
				memo[wire] = dfs(wires, left, memo) << dfs(wires, right, memo)
			case "RSHIFT":
				fmt.Sscanf(wires[wire], "%s RSHIFT %s", &left, &right)
				memo[wire] = dfs(wires, left, memo) >> dfs(wires, right, memo)
			}
		}
	}

	return memo[wire]
}

func main() {
	bufio := bufio.NewReader(os.Stdin)

	wires := map[string]string{}

	for {
		line, err := bufio.ReadString('\n')

		if err == io.EOF {
			break
		}

		line = line[:len(line)-1]
		splits := strings.Split(line, " -> ")

		wires[splits[1]] = splits[0]
	}

	partOne := dfs(wires, "a", map[string]int{})
	fmt.Println(partOne)

	wires["b"] = strconv.Itoa(partOne)

	fmt.Println(dfs(wires, "a", map[string]int{}))
}
