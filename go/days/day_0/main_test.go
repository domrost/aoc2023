package day_0_test

import (
	"testing"

	"github.com/domrost/aoc2023/go/days/day_0"
	"github.com/domrost/aoc2023/go/internal/load"
)

func TestPartOne(t *testing.T) {
	t.Parallel()

	input := load.LoadInputLines("input_1_test.txt")
	expectedResult := load.LoadFirstInputLine("solution_1.txt")
	result := day_0.Part1(input)

	if result != expectedResult {
		t.Errorf("expected result was %s, but got %s instead", expectedResult, result)
	}
}

func TestPartTwo(t *testing.T) {
	t.Parallel()

	input := load.LoadInputLines("input_2_test.txt")
	expectedResult := load.LoadFirstInputLine("solution_2.txt")
	result := day_0.Part2(input)

	if result != expectedResult {
		t.Errorf("expected result was %s, but got %s instead", expectedResult, result)
	}
}
