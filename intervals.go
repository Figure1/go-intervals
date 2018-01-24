package intervals

// Intervals keeps a map of start integers to end integers
type Intervals map[int]int

// New will return an empty set of intervals
func New() Intervals {
	return make(Intervals)
}

// Insert adds a new interval.
// If the new interval overlaps with any existing intervals,
// the existing intervals will be collapsed/deleted as necessary.
func (ints Intervals) Insert(start, end int) {
	if start > end {
		start, end = end, start
	}
	toDelete := []int{}
	startOverlap, endOverlap := -1, -1
	for intStart, intEnd := range ints {
		if intStart <= start && end <= intEnd {
			return
		} else if start <= intStart && intEnd <= end {
			toDelete = append(toDelete, intStart)
		} else if intStart < start && start <= intEnd+1 {
			startOverlap = intStart
		} else if end >= intStart-1 && end < intEnd {
			endOverlap = intStart
		}
	}
	for i := range toDelete {
		delete(ints, toDelete[i])
	}
	if startOverlap < 0 && endOverlap < 0 {
		ints[start] = end
	} else if startOverlap >= 0 && endOverlap >= 0 {
		ints[startOverlap] = ints[endOverlap]
		delete(ints, endOverlap)
	} else if startOverlap >= 0 {
		ints[startOverlap] = end
	} else if endOverlap >= 0 {
		ints[start] = ints[endOverlap]
		delete(ints, endOverlap)
	}
}

// Delete removes the specified interval.
// If the interval overlaps with existing intervals,
// those intervals are updated/deleted as necessary.
func (ints Intervals) Delete(start, end int) {
	if start > end {
		start, end = end, start
	}
	toDelete := []int{}
	startOverlap, endOverlap := -1, -1
	for intStart, intEnd := range ints {
		if start == intStart && end == intEnd {
			toDelete = append(toDelete, intStart)
			break
		} else if intStart < start && end < intEnd {
			ints[intStart] = start - 1
			ints[end+1] = intEnd
			break
		} else if start <= intStart && intEnd <= end {
			toDelete = append(toDelete, intStart)
		} else if intStart < start && start <= intEnd {
			startOverlap = intStart
		} else if intStart <= end && end < intEnd {
			endOverlap = intStart
		}
	}
	for i := range toDelete {
		delete(ints, toDelete[i])
	}
	if startOverlap >= 0 {
		ints[startOverlap] = start - 1
	}
	if endOverlap >= 0 {
		ints[end+1] = ints[endOverlap]
		delete(ints, endOverlap)
	}
}

// Contains will check if there is an interval that contains
// the specified integer.
func (ints Intervals) Contains(x int) bool {
	for start, end := range ints {
		if start <= x && x <= end {
			return true
		}
	}
	return false
}

// Overlaps will check to see if the specified interval overlaps
// with any existing intervals
func (ints Intervals) Overlaps(x, y int) bool {
	if x > y {
		x, y = y, x
	}
	for start, end := range ints {
		if start <= x && x <= end {
			return true
		} else if start <= y && y <= end {
			return true
		} else if x <= start && end <= y {
			return true
		}
	}
	return false
}

// Equal returns whether two Intervals are equal
func Equal(x, y Intervals) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if _, ok := y[i]; !ok || y[i] != x[i] {
			return false
		}
	}
	return true
}
