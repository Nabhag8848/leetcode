func merge(intervals [][]int) [][]int {

    sort.Slice(intervals, func (i, j int) bool {
        if intervals[i][0] != intervals[j][0] {
            return intervals[i][0] < intervals[j][0]
        }

        return intervals[i][1] < intervals[j][1]
    })

    result := make([][]int, 0)

    i := 1

    start := intervals[0][0]
    end := intervals[0][1]

    for i < len(intervals) {
        if intervals[i][0] <= end {
            end = max(end, intervals[i][1])
        } else {
            result = append(result, []int{start, end})
            start = intervals[i][0]
            end = intervals[i][1]
        }

        i++
    }

    result = append(result, []int{start, end})

    return result
}
