func findMin(nums []int) int {
    low := 0
    high := len(nums) - 1
    mid := low + (high - low) / 2
    min := nums[mid]

    for low <= high {
        mid = low + (high - low) / 2

        if nums[mid] >= nums[low] {
            min = int(math.Min(float64(min), float64(nums[low])))
            low = mid + 1
        } else {
            min = int(math.Min(float64(min), float64(nums[mid])))
            high = mid - 1
        }
            
    }   

    return min
}

/*
    1,2,3,4,5

*/