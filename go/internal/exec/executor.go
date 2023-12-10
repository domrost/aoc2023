package exec

import (
	"github.com/domrost/aoc2023/go/days/day_0"
	"github.com/domrost/aoc2023/go/days/day_06"
	"github.com/domrost/aoc2023/go/days/day_07"
	"github.com/domrost/aoc2023/go/days/day_08"

	"github.com/domrost/aoc2023/go/internal/load"
)

// RunChallenge executes the challenge of a specific day with the provided input.
func RunChallenge(day int, inputPath string, mode int) {
	input := load.LoadInputLines(inputPath)
	mapping := map[int]func([]string, int){
		0: day_0.Run,
		6: day_06.Run,
		7: day_07.Run,
		8: day_08.Run,
	}

	mapping[day](input, mode)
}
