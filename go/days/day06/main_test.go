package day06

import (
	"testing"

	"github.com/domrost/aoc2023/go/internal/load"
)

func TestPartOne(t *testing.T) {
	t.Parallel()

	input := load.LoadInputLines("input_1_test.txt")
	expectedResult := load.LoadFirstInputLine("solution_1.txt")
	result := Part1(input)

	if result != expectedResult {
		t.Errorf("expected result was %s, but got %s instead", expectedResult, result)
	}
}

func TestPartTwo(t *testing.T) {
	t.Parallel()

	input := load.LoadInputLines("input_2_test.txt")
	expectedResult := load.LoadFirstInputLine("solution_2.txt")
	result := Part2(input)

	if result != expectedResult {
		t.Errorf("expected result was %s, but got %s instead", expectedResult, result)
	}
}

func TestCalculateWins(t *testing.T) {
	race := race{time: 7, distance: 9}
	result := calculateWins(race)

	if result != 4 {
		t.Errorf("expected result was %d, but got %d instead", 4, result)
	}

}
