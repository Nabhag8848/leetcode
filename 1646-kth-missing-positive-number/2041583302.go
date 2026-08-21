func findKthPositive(arr []int, k int) int {
    return optimalApproach(arr, k)
}

func optimalApproach(arr [] int, k int) int {
    low:= 0
    high := len(arr) - 1
    
    current_missing_count := 0

    for low <= high {
        missing_index := low + (high - low) / 2
        missing_count := arr[missing_index] - (missing_index + 1)
        
        if missing_count < k {
            low = missing_index + 1
            current_missing_count += missing_count
        } else if missing_count >= k {
            high = missing_index - 1
        }  

    }

    return low + k
}

func approach(arr [] int, k int) int {
    last_element := arr[len(arr) - 1]
    j:=0
    current_missing_count := 0

    for i := 1; i <= last_element; i++ {

        if arr[j] == i {
            j++
        } else {
            current_missing_count++
        }

        if current_missing_count == k {
            return i
        }
    }

    return k - current_missing_count + last_element
}

/*
    4 - (2 + 1)

    2,3,4,7,11

    1 -> arr[len(arr) - 1]

*/  
