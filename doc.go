// Package intervals allows managing of integer intervals. Inserting and deleting intervals will automatically collapse/split existing intervals as needed.
//
// Example Usage
//
// The following example creates an interval set, adds intervals, deletes intervals and prints the result.
// 	import (
// 		"fmt"
//
// 		"github.com/Figure1/go-intervals"
// 	)
//
// 	func main() {
// 		intervalSet := intervals.New()
// 		intervalSet.Add(2, 6)
// 		intervalSet.Add(10, 13)
// 		intervalSet.add(15, 20)
//
// 		intervalSet.Contains(12) // true
// 		intervalSet.Overlaps(18, 22) // true
//
// 		intervalSet.Delete(5, 12)
// 		fmt.Println(intervalSet) // map[2:4 13:13 18:22]
//	}
package intervals
