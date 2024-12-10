package day10

import (
	"fmt"

	"github.com/gomsim/Advent-of-code/shared/dat"
	"github.com/gomsim/Advent-of-code/shared/input"
	"github.com/gomsim/Advent-of-code/shared/registrar"
)

func init() {
	registrar.Register("2024", "day10", "part2", Part2)
}

func Part2(file string) {
	in := input.Matrix(file)

	trailheads := findTrailheads(in)
	score := trailRatings(in, trailheads)

	fmt.Println(score)
}

func trailRatings(matrix [][]string, trailheads map[dat.Vec[int]]bool) int {
	score := 0
	for trailhead := range trailheads {
		_, rating := trail(matrix, trailhead, make(map[dat.Vec[int]]bool))
		score += rating
	}
	return score
}
