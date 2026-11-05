func insert(intervals [][]int, newInterval []int) [][]int {
    arr := make([][]int, 0)

    arr = append(arr, intervals...)
    arr = append(arr, newInterval)

    sort.Slice(arr, func (i, j int) bool {
        if arr[i][0] < arr[j][0] {
            return true
        } 
        
        if arr[i][0] == arr[j][0] {
            return arr[i][1] < arr[j][1]
        }

        return false
    })

    return mergeIntervals(arr)
}

func mergeIntervals(arr [][]int) [][]int {
    i := 0
    result := make([][]int, 0)

    for i < len(arr) {
       	start := arr[i][0]
    	end := arr[i][1]
    	i++
    		
        for i < len(arr) && arr[i][0] <= end {
    		if arr[i][1] > end {
    			end = arr[i][1]
    		}

    	    i++
         }

        result = append(result, []int{start, end})
    }

    return result
}
