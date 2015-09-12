package models

import (
	"github.com/Unknwon/log"
)

// FibonacciSequence generates a sequence of first n Fibonacci numbers
// starting from 0. It returns nil if num is negative.
func FibonacciSequence(num int) []int64 {
	if num < 0 {
		return nil
	} else if num == 0 {
		return []int64{}
	} else if num == 1 {
		return []int64{0}
	}

	// Calculate sequence.
	seq := []int64{0, 1}
	cursor := 2
	var lastNum int64 = 1
	for cursor < num {
		lastNum = seq[cursor-2] + seq[cursor-1]

		// Check overflow, the current number ovbiously
		// cannot be smaller then previous one.
		if lastNum < seq[cursor-1] {
			log.Error("Unable to produce %dth fibonacci value", cursor)
			return nil
		}
		seq = append(seq, lastNum)
		cursor++
	}
	return seq
}
