package day07

import (
	"fmt"
	"time"

	"github.com/domrost/aoc2023/go/days/day07_part1"
	"github.com/domrost/aoc2023/go/days/day07_part2"
)

// Run function of the daily challenge
func Run(input []string, mode int) {
	var time1 time.Duration
	var time2 time.Duration

	if mode == 1 || mode == 3 {
		start1 := time.Now()
		res1 := Part1(input)
		time1 = time.Since(start1)

		fmt.Printf("Part one - Result: %v, Time: %v\n", res1, time1)
	}
	if mode == 2 || mode == 3 {
		start2 := time.Now()
		res2 := Part2(input)
		time2 = time.Since(start2)

		fmt.Printf("Part two - Result: %v, Time: %v\n", res2, time2)
	}
	if mode == 3 {
		fmt.Printf("Time Sum: %v\n", (time1 + time2))
	}
}

func Part1(input []string) string {
	return day07_part1.Run(input)
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	return day07_part2.Run(input)
}
