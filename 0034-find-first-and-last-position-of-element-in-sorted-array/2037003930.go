func searchRange(nums []int, target int) []int {
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
        break
      }
    }

    if low > high {
        return []int{-1, -1}
    }

    low = mid
    high = mid

    for low > -1 && nums[low] == target {
        low--
    }

    for high < len(nums) && nums[high] == target {
        high++
    }

    return []int{low + 1,high - 1}

}