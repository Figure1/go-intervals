package intervals

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestMapInsert(t *testing.T) {
	testCases := [][]Intervals{
		{{1: 5, 6: 10}, {1: 10}},
		{{1: 5, 7: 10}, {1: 5, 7: 10}},
		{{1: 5, 2: 3}, {1: 5}},
		{{1: 5, 0: 10}, {0: 10}},
		{{1: 5, 2: 10}, {1: 10}},
		{{2: 6, 0: 3}, {0: 6}},
		{{1: 5, 7: 10, 2: 8}, {1: 10}},
		{{1: 5, 6: 6}, {1: 6}},
		{{1: 5, 0: 3}, {0: 5}},
		{{6: 10, 7: 11}, {6: 11}},
		{{2: 3, 1: 5}, {1: 5}},
		{{36: 37, 33: 37}, {33: 37}},
	}

	for i := range testCases {
		input, expected := testCases[i][0], testCases[i][1]
		t.Run(fmt.Sprintf("input: %v expected: %v", input, expected), func(t *testing.T) {
			actual := make(Intervals)
			for k, v := range input {
				actual.Insert(k, v)
			}
			if !Equal(actual, expected) {
				t.Errorf("input: %v expected: %v actual: %v", input, expected, actual)
			}
		})
	}
}

func TestMapDelete(t *testing.T) {
	testCases := [][]Intervals{
		// delete contains intervals
		{{1: 5, 10: 11, 13: 15}, {10: 11}, {1: 5, 13: 15}},
		{{1: 5, 10: 11, 13: 15}, {9: 12}, {1: 5, 13: 15}},
		{{1: 5, 10: 11, 13: 15}, {0: 12}, {13: 15}},
		{{1: 5, 10: 11, 13: 15}, {0: 16}, {}},
		// no change
		{{1: 5, 10: 11, 13: 15}, {6: 7}, {1: 5, 10: 11, 13: 15}},
		// start contained
		{{1: 5, 10: 11, 13: 15}, {2: 7}, {1: 1, 10: 11, 13: 15}},
		{{1: 5, 10: 11, 13: 15}, {1: 7}, {10: 11, 13: 15}},
		// delete contained
		{{1: 5, 10: 11, 13: 15}, {2: 5}, {1: 1, 10: 11, 13: 15}},
		{{1: 5, 10: 11, 13: 15}, {1: 3}, {4: 5, 10: 11, 13: 15}},
		{{1: 5, 10: 11, 13: 15}, {2: 2}, {1: 1, 3: 5, 10: 11, 13: 15}},
		{{1: 5, 10: 11, 13: 15}, {13: 13}, {1: 5, 10: 11, 14: 15}},
		{{1: 5, 10: 11, 13: 15}, {15: 15}, {1: 5, 10: 11, 13: 14}},
		// end contained
		{{1: 5, 10: 11, 13: 15}, {12: 14}, {1: 5, 10: 11, 15: 15}},
		{{1: 5, 10: 11, 13: 15}, {12: 15}, {1: 5, 10: 11}},
		{{1: 5, 10: 11, 13: 15}, {12: 13}, {1: 5, 10: 11, 14: 15}},
	}

	for i := range testCases {
		input, toDelete, expected := testCases[i][0], testCases[i][1], testCases[i][2]
		t.Run(fmt.Sprintf("input: %v deleted: %v", input, toDelete), func(t *testing.T) {
			actual := make(Intervals)
			for k, v := range input {
				actual.Insert(k, v)
			}
			for k, v := range toDelete {
				actual.Delete(k, v)
			}
			if !Equal(actual, expected) {
				t.Errorf("input: %v deleted: %v expected: %v actual: %v", input, toDelete, expected, actual)
			}
		})
	}
}

func TestContains(t *testing.T) {
	testCases := []struct {
		intervals Intervals
		value     int
		expected  bool
	}{
		{Intervals{1: 3}, 0, false},
		{Intervals{1: 3}, 1, true},
		{Intervals{1: 3}, 2, true},
		{Intervals{1: 3}, 3, true},
		{Intervals{1: 3}, 4, false},
	}

	for _, test := range testCases {
		t.Run(fmt.Sprintf("interval: %v value: %v", test.intervals, test.value), func(t *testing.T) {
			if test.intervals.Contains(test.value) != test.expected {
				t.Errorf("interval: %v value: %v expected: %v", test.intervals, test.value, test.expected)
			}
		})
	}
}

func TestOverlaps(t *testing.T) {
	testCases := []struct {
		intervals Intervals
		start     int
		end       int
		expected  bool
	}{
		{Intervals{-3: 5}, -6, -4, false},
		{Intervals{-3: 5}, -6, -3, true},
		{Intervals{-3: 5}, -6, 5, true},
		{Intervals{-3: 5}, -6, 6, true},
		{Intervals{-3: 5}, -3, 6, true},
		{Intervals{-3: 5}, -3, 5, true},
		{Intervals{-3: 5}, -3, 3, true},
		{Intervals{-3: 5}, -2, 3, true},
		{Intervals{-3: 5}, -2, 6, true},
		{Intervals{-3: 5}, 5, 6, true},
		{Intervals{-3: 5}, 6, 8, false},
	}

	for _, test := range testCases {
		t.Run(fmt.Sprintf("intervals: %v start: %v end: %v", test.intervals, test.start, test.end), func(t *testing.T) {
			if test.intervals.Overlaps(test.start, test.end) != test.expected {
				t.Errorf("interval: %v start: %v end: %v expected: %v", test.intervals, test.start, test.end, test.expected)
			}
		})
	}
}

func BenchmarkInsertLotsOfSmall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		data := make(Intervals)
		data.Insert(0, 5)
		data.Insert(6, 15)
		data.Insert(20, 25)
	}
}

func BenchmarkInsertNoOverlaps(b *testing.B) {
	data := make(Intervals)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		data.Insert(i*2, i*2)
	}
}

func BenchmarkInsertLotsOfOverlaps(b *testing.B) {
	data := make(Intervals)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		start, end := rand.Intn(1000), rand.Intn(10)
		data.Insert(start, start+end)
	}
}

func BenchmarkDeleteLotsOfSmall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		data := make(Intervals)
		data.Insert(0, 5)
		data.Delete(20, 25)
		data.Delete(6, 15)
	}
}

func BenchmarkDeleteLotsOfOverlaps(b *testing.B) {
	data := make(Intervals)
	for i := 0; i < 1000; i++ {
		start, end := rand.Intn(1000), rand.Intn(10)
		data.Insert(start, start+end)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		start, end := rand.Intn(1000), rand.Intn(10)
		data.Delete(start, end)
	}
}
