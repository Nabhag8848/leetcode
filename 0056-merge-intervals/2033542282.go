func merge(intervals [][]int) [][]int {
  sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] != intervals[j][0] {
			return intervals[i][0] < intervals[j][0]
		}
		return intervals[i][1] < intervals[j][1]
  })

  i := 0

  result := make([][]int, 0)

    for i < len(intervals) {
       	start := intervals[i][0]
    	end := intervals[i][1]
    	i++
    		
        for i < len(intervals) && intervals[i][0] <= end {
    		if intervals[i][1] > end {
    			end = intervals[i][1]
    		}
    	  i++
         }

     result = append(result, []int{start, end})
    }
  return result

}