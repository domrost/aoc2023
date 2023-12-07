package day_06

import (
	"fmt"
	"math"
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

type race struct {
	time     int
	distance int
}

// Part1 solves the first part of the exercise
func Part1(input []string) string {
	times := parseTimesDistances(input[0])
	distances := parseTimesDistances(input[1])
	races := createRaces(times, distances)

	res := 1
	for _, race := range races {
		res *= calculateWins(race)
	}

	return strconv.Itoa(res)
}

func parseTimesDistances(timesDistString string) []string {
	times := strings.Fields(timesDistString)
	return times[1:]
}

func parseTimesDistancesV2(timesDistString string) []string {
	timeDist := strings.Join(strings.Fields(timesDistString)[1:], "")
	return []string{timeDist}
}

func createRaces(times, distances []string) []race {
	races := []race{}
	for i := range times {
		t, _ := strconv.Atoi(times[i])
		d, _ := strconv.Atoi(distances[i])
		races = append(races, race{time: t, distance: d})
	}
	return races
}

func calculateWins(race race) int {
	// Quadratic equation, solve with pq formula
	time := float64(race.time)
	dist := float64(race.distance)

	lowerBound := (time / 2) - math.Sqrt((math.Pow((-1*time/2.0), 2) - dist))
	upperBound := (time / 2) + math.Sqrt((math.Pow((-1*time/2.0), 2) - dist))

	// we need to check, whether we got an integer. If so, we have to add/subtract 1, because otherwise we would count draws
	if lowerBound == float64(int64(lowerBound)) {
		lowerBound++
	}

	if upperBound == float64(int64(upperBound)) {
		upperBound--
	}

	wins := int(math.Floor(upperBound)-math.Ceil(lowerBound)) + 1
	// fmt.Printf("Calculated wins for race: %v: %d\n", race, wins)
	return wins
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	times := parseTimesDistancesV2(input[0])
	distances := parseTimesDistancesV2(input[1])
	races := createRaces(times, distances)

	res := 1
	for _, race := range races {
		res *= calculateWins(race)
	}

	return strconv.Itoa(res)
}
