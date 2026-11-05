func insert(intervals [][]int, newInterval []int) [][]int {
    arr := make([][]int, 0)

    i := 0

    for i < len(intervals) && intervals[i][1] < newInterval[0] {
        arr = append(arr, intervals[i])
        i++
    }

    for i < len(intervals) && intervals[i][0] <= newInterval[1] {
        newInterval[0] = min(newInterval[0], intervals[i][0])
        newInterval[1] = max(newInterval[1], intervals[i][1])
        i++
    }

    arr = append(arr, newInterval)

    for i < len(intervals) {
        arr = append(arr, intervals[i])
        i++
    }

    return arr
}
