package main

import "fmt"

func main() {
    intervals := [][]int{{10, 13}, {1, 3}, {2, 6}, {8, 10}, {1, 3}, {6, 8}, {7, 8}, {1, 3},{14,17},{20,22},{18,19},{17,23}}
    fmt.Print(extractAllIntervals(intervals))
}

func extractAllIntervals(intervals [][]int) [][]int {
    extractedIntervals := [][]int{}
    for len(intervals) > 0 {
        merged := lookForMergable(&intervals)
        extractedIntervals = append(extractedIntervals, merged)
    }
    return extractedIntervals
}

func lookForMergable(intervalsPtr *[][]int) []int {
    intervals := *intervalsPtr

    ToMerge := [][]int{intervals[0]}

    intervals = append(intervals[:0], intervals[1:]...)

    for i := 0; i < len(ToMerge); i++ {
        for j := len(intervals) - 1; j >= 0; j-- {
            if decideToMerge(ToMerge[i], intervals[j]) {
                ToMerge = append(ToMerge, intervals[j])
                intervals = append(intervals[:j], intervals[j+1:]...)
            }
        }
    }
    merged := merge(ToMerge)
    *intervalsPtr = intervals
    return merged
}

func decideToMerge(a, b []int) bool {
    return a[1] >= b[0] && b[1] >= a[0]
}

func merge(newIntervals [][]int) []int {
    min := newIntervals[0][0]
    max := newIntervals[0][1]
    for i := 1; i < len(newIntervals); i++ {
        if newIntervals[i][0] < min {
            min = newIntervals[i][0]
        }
        if newIntervals[i][1] > max {
            max = newIntervals[i][1]
        }
    }
    return []int{min, max}
}
