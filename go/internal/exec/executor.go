package exec

import (
	"github.com/domrost/aoc2023/go/days/day0"
	"github.com/domrost/aoc2023/go/days/day06"
	"github.com/domrost/aoc2023/go/days/day07"
	"github.com/domrost/aoc2023/go/days/day08"
	"github.com/domrost/aoc2023/go/internal/load"
)

// RunChallenge executes the challenge of a specific day with the provided input.
func RunChallenge(day int, inputPath string, mode int) {
	input := load.LoadInputLines(inputPath)
	mapping := map[int]func([]string, int){
		0: day0.Run,
		6: day06.Run,
		7: day07.Run,
		8: day08.Run,
	}

	mapping[day](input, mode)
}
