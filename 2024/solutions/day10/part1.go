package day10

import (
	"fmt"

	"github.com/gomsim/Advent-of-code/shared/dat"
	"github.com/gomsim/Advent-of-code/shared/input"
	"github.com/gomsim/Advent-of-code/shared/registrar"
)

func init() {
	registrar.Register("2024", "day10", "part1", Part1)
}

func Part1(file string) {
	in := input.Matrix(file)

	trailheads := findTrailheads(in)
	score := scoreTrailheads(in, trailheads)

	fmt.Println(score)
}

func scoreTrailheads(matrix [][]string, trailheads map[dat.Vec[int]]bool) int {
	score := 0
	for trailhead := range trailheads {
		trails := paths(matrix, trailhead, []dat.Vec[int]{})
		distinct := make(map[dat.Vec[int]]bool)
		for _, trail := range trails {
			distinct[trail] = true
		}
		score += len(distinct)
	}
	return score
}
