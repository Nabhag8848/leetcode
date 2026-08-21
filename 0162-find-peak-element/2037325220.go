func findPeakElement(nums []int) int {
    if len(nums) == 1 {
        return 0
    }

    if nums[0] > nums[1] {
        return 0
    }

    if nums[len(nums) - 1] > nums[len(nums) - 2] {
        return len(nums) - 1
    }

    low := 1
    high := len(nums) - 2
    mid := low + (high - low) / 2


    for low <= high {
        mid = low + (high - low) / 2

        if nums[mid] > nums[mid - 1] && nums[mid] > nums[mid + 1] {
            break
        } else if nums[mid] < nums[mid - 1] {
            high = mid - 1
        } else {
            low = mid + 1
        }
    }

    return mid
}

/*
    1, 2, 1, 3, 5, 6, 4
    
    left = math.min
    right = math.min

    left = i - 1
    right = i + 1

    peak_index = 1


*/