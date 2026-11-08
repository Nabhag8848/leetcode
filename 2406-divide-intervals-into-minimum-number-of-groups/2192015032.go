func minGroups(intervals [][]int) int {
    events := make([][2]int, 0, 2*len(intervals))

    for _, interval := range intervals {
        events = append(events, [2]int{interval[0], 1})
        events = append(events, [2]int{interval[1] + 1, -1})
    }

    sort.Slice(events, func(i, j int) bool {
        if events[i][0] != events[j][0] {
            return events[i][0] < events[j][0]
        }

        return events[i][1] < events[j][1]
    })

    fmt.Println(events)
    count, maxCount := 0, 0

    for _, event := range events {
        count += event[1]
        maxCount = max(maxCount, count)
    }

    return maxCount
}
