func eraseOverlapIntervals(intervals [][]int) int {
    result := 1

    sort.Slice(intervals, func (i, j int) bool {
        return intervals[i][1] < intervals[j][1]
    })

    free_time := intervals[0][1]

    for i := 1; i < len(intervals); i++ {
        start := intervals[i][0]

        if start >= free_time {
            result++
            free_time = intervals[i][1]
        }

    }

    return len(intervals) - result
}
