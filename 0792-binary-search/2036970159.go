func search(nums []int, target int) int {
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


/*
  1 2 3 4 5

 
*/  