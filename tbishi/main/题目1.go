package main

import "sort"

// /**
// * Definition for an interval.
// * type Interval struct {
// *	   Start int
// *	   End   int
// * }
// */
type Interval struct {
	Start int
	End int
}
func intervalIntersection(A []Interval, B []Interval) []Interval {
   if len(A)==0||len(B)==0{
   	return nil
   }
	sort.Slice(A, func(i, j int) bool {
		return A[i].Start  > A[j].Start
	})
	sort.Slice(B, func(i, j int) bool {
		return B[i].Start  > B[j].Start
	})
   j := 0

   }
}
