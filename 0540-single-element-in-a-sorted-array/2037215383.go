func singleNonDuplicate(nums []int) int {
    if len(nums) == 1 {
        return nums[0]
    }

    if nums[0] != nums[1] {
        return nums[0]
    } 
    
    if nums[len(nums) - 1] != nums[len(nums) - 2] {
        return nums[len(nums) - 1]
    }

    low := 1
    high := len(nums) - 2
    mid := low + (high - low) / 2

    for low <= high {
        mid = low + (high - low) / 2

        if nums[mid] != nums[mid - 1] && nums[mid] != nums[mid + 1] {
            return nums[mid]
        } else if mid % 2 == 1 {
            if nums[mid] == nums[mid - 1] {
              low = mid + 1
            } else {
              high = mid - 1
            }
        } else {
            if nums[mid] == nums[mid + 1] {
                low = mid + 1
            } else {
                high = mid - 1
            }
        }
    }

    return -1
}

/*
    1 1 2 3 3 4 4 8 8

    mid = 9 / 2 (4)

    if left half is even go to right
    if right half is even go to left

    3,3,7,7, 9, 10, 10, 11, 11, 12, 12

*/