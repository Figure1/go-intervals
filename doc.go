// Package intervals allows managing of integer intervals. Inserting and deleting intervals will automatically collapse/split existing intervals as needed.
//
// Example Usage
//
// The following example creates an interval set, adds intervals, deletes intervals and prints the result.
// 	import (
// 		"fmt"
// 		"github.com/Figure1/go-intervals"
// 	)
//
// 	func main() {
//		intervalSet := intervals.New()
//		intervalSet.Insert(2, 6)
//		intervalSet.Insert(10, 13)
//		intervalSet.Insert(15, 20)
//
//		intervalSet.Contains(12) // true
//		intervalSet.Contains(14) // false
//		intervalSet.Overlaps(18, 22) // true
//		intervalSet.Overlaps(21, 24) // false
//
//		intervalSet.Delete(5, 15)
//		fmt.Println(intervalSet) // map[2:4 16:20]
//	}
package intervals
