func searchInsert(nums []int, target int) int {
    low  := 0
    high := len(nums) - 1
    mid := low + (high - low) / 2


    for low <= high { 
      mid = low + (high - low) / 2

      if nums[mid] > target {
        high = mid - 1
      } else if nums[mid] < target {
        low = mid + 1
      } else {
        return mid
      }
    }

    if nums[mid] > target {
        for mid > -1 && nums[mid] > target {
            mid = mid - 1
        }

        return mid + 1
    }


    for mid < len(nums) && nums[mid] < target {
        mid = mid + 1
    }

    return mid
}