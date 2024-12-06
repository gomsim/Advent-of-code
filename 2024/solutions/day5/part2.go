package day5

import (
	"fmt"
	"time"

	"github.com/gomsim/Advent-of-code/shared/data"
	"github.com/gomsim/Advent-of-code/shared/input"
	"github.com/gomsim/Advent-of-code/shared/register"
)

func init() {
	register.Part2("2024", "day5", Part2)
}

func Part2(file string) {
	in := input.Slice(file)

	rules, updates := parse(in)
	incorrectUpdates := findIncorrect(rules, updates)
	newRules := newRules(rules)
	correctedUpdates := correctTheWrong(newRules, incorrectUpdates)
	sum := sum(correctedUpdates)

	fmt.Println(sum)
}

func findIncorrect(rules map[string][]string, updates []map[string]int) []map[string]int {
	incorrectUpdates := make([]map[string]int, 0, len(updates)/2)
	for _, update := range updates {
		if !correct(rules, update) {
			incorrectUpdates = append(incorrectUpdates, update)
		}
	}
	return incorrectUpdates
}

func newRules(rules map[string][]string) map[string][]string {
	newRules := make(map[string][]string, len(rules))
	for key, val := range rules {
		for _, rule := range val {
			newRules[rule] = append(newRules[rule], key)
		}
	}
	return newRules
}

func correctTheWrong(rules map[string][]string, updates []map[string]int) []map[string]int {
	corrected := make([]map[string]int, len(updates))
	fmt.Println("Timer start!")
	before := time.Now()
	for i, update := range updates {
		// convert
		linked := data.NewWinked[string]()
		for val := range update {
			linked.Append(val)
		}

		// rearrange
		linked = toCorrect(rules, linked)

		// convert back
		fixed := make(map[string]int, len(update))
		for i, val := range linked.Iter {
			fixed[val] = i
		}
		corrected[i] = fixed
	}
	fmt.Printf("Sorting took %v!\n", time.Since(before))
	return corrected
}

func toCorrect(rules map[string][]string, update data.Winked[string]) data.Winked[string] {
	modified := true
	for modified {
		modified = false
		for _, page := range update.Iter {
			for _, mustBefore := range rules[page] {
				if update.Contains(mustBefore) && !update.After(page, mustBefore) {
					update.Insert(page, mustBefore)
					modified = true
					break
				}
			}
			if modified {
				break
			}
		}
	}
	return update
}