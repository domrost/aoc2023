package day09

import (
	"fmt"
	"strconv"
	"strings"
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

type history []int

// Part1 solves the first part of the exercise
func Part1(input []string) string {

	sum := 0
	for _, line := range input {
		history := createHistory(line)
		predictedValue := extrapolate(history)
		sum += predictedValue
	}
	return strconv.Itoa(sum)
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	sum := 0
	for _, line := range input {
		history := createHistory(line)
		predictedValue := extrapolateBack(history)
		sum += predictedValue
	}
	return strconv.Itoa(sum)
}

func createHistory(line string) history {
	valueStrings := strings.Split(strings.Trim(line, " "), " ")
	h := history{}
	for _, valueString := range valueStrings {
		value, _ := strconv.Atoi(valueString)
		h = append(h, value)
	}
	return h
}

func extrapolate(hist history) int {
	// Calculate all differences first
	isAll0 := true
	differences := history{}
	for i := 0; i < len(hist)-1; i++ {
		differences = append(differences, hist[i+1]-hist[i])
		if isAll0 && differences[i] != 0 {
			isAll0 = false
		}
	}

	// If differences are not all 0, we dive one level down
	if !isAll0 {
		prediction := extrapolate(differences)
		return hist[len(hist)-1] + prediction // And return the last element of the history plus the prediction we received from below
	} else { // Otherwise, we simply return the last element of the history we have received
		return hist[len(hist)-1]
	}
}

func extrapolateBack(hist history) int {
	// Calculate all differences first
	isAll0 := true
	differences := history{}
	for i := 0; i < len(hist)-1; i++ {
		differences = append(differences, hist[i+1]-hist[i])
		if isAll0 && differences[i] != 0 {
			isAll0 = false
		}
	}

	// If differences are not all 0, we dive one level down
	if !isAll0 {
		prediction := extrapolateBack(differences)
		return hist[0] - prediction // And return the first element of the history minus the prediction we received from below
	} else { // Otherwise, we simply return the first element of the history we have received
		return hist[0]
	}
}
