func search(nums []int, target int) int {
    return recursive(nums, target, 0, len(nums) - 1)
}

func iterative(nums []int, target int) int {
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
    return -1
   }

   return mid
}

func recursive(nums []int, target, low, high int) int {
    mid := low + (high - low) / 2

    if low > high {
        return -1
    }

    if nums[mid] > target { 
       return recursive(nums, target, low, mid - 1)
    } else if nums[mid] < target {
       return recursive(nums, target, mid + 1, high)
    } 

    return mid
}


/*
  1 2 3 4 5

 
*/  