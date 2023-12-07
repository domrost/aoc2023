package day_0

import (
	"fmt"
	"sort"
	"strconv"
	"time"
)

// Run function of the daily challenge
func Run(input []string, mode int) {
	var time1 time.Duration
	var time2 time.Duration

	if mode == 1 || mode == 3 {
		start1 := time.Now()
		res1 := Part1(input)
		time1 = time.Since(start1)

		fmt.Printf("Part 1 - Result: %v, Time: %v\n", res1, time1)
	}
	if mode == 2 || mode == 3 {
		start2 := time.Now()
		res2 := Part1(input)
		time2 = time.Since(start2)

		fmt.Printf("Part 2 - Result: %v, Time: %v\n", res2, time2)
	}
	if mode == 3 {
		fmt.Printf("Time Sum: %v\n", (time1 + time2))
	}
}

// getElfCalories reads the input lines and returns a slice of calories for each elf
func getElfCalories(input []string) []int {
	var calories []int
	currentElf := 0

	for _, i := range input {
		if i == "" {
			calories = append(calories, currentElf)
			currentElf = 0
		} else {
			cal, _ := strconv.Atoi(i)
			currentElf += cal
		}
	}
	calories = append(calories, currentElf)

	sort.Slice(calories, func(i, j int) bool {
		return calories[i] > calories[j]
	})

	return calories
}

// Part1 solves the first part of the exercise
func Part1(input []string) string {
	calories := getElfCalories(input)
	maxCalories := calories[0]
	return strconv.Itoa(maxCalories)
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	calories := getElfCalories(input)

	topCalories := 0
	for i := 0; i < 3; i++ {
		topCalories += calories[i]
	}

	return strconv.Itoa(topCalories)
}
