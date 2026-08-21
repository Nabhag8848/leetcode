func splitArray(nums []int, k int) int {
    low := findMax(nums)
    high := 0

    for _, value := range nums {
        high += value
    }

    for low <= high {
        mid := low + (high - low) / 2
        subarrayCount := findSubArrayCount(nums, mid)

        if subarrayCount > k {
            low = mid + 1
        } else if subarrayCount <= k {
            high = mid - 1
        }
    }

    return low
}

/*
    [1,2,3,4,5]

    9 -> 15 (12)

    9 -> 12 (11)
    9 -> 11 (10)
    9 -> 10 (9)
    9 -> 9 (9)

*/

func findSubArrayCount(nums []int, sum int) int {
    count := 1
    current_sum := 0

    for i := range nums {
        current_sum += nums[i]

        if current_sum > sum {
            current_sum = nums[i]
            count++
        }
    }

    return count
}

func findMax(arr []int) int {
	max := arr[0]
	for _, num := range arr {
		if num > max {
			max = num
		}
	}

	return max
}

