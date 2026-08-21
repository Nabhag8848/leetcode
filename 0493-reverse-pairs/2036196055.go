func reversePairs(nums []int) int {
    count := countReversePairs(nums, 0, len(nums))
    return count
}

func countPairs(nums []int, start, mid, end int) int {
    right := mid
    count := 0
    for i:= start; i < mid; i++ {
        for right < end && int64(nums[i]) > 2*int64(nums[right]) {
            right++
        }

        count+= right - mid
    }

    return count
}

func countReversePairs(nums []int, start, end int) int {
    count := 0

    if end-start <= 1 {
        return 0
    }

    mid := start + (end-start)/2

    count += countReversePairs(nums, start, mid)
    count += countReversePairs(nums, mid, end)
    count += countPairs(nums, start, mid, end)

    merge(nums, start, mid, end)
    return count
}

func merge(nums []int, start, mid, end int) {
    temp := make([]int, end-start)

    i := start
    j := mid
    k := 0

    for i < mid && j < end {

        if nums[i] <= nums[j] {
            temp[k] = nums[i]
            i++
        } else {
            temp[k] = nums[j]
            j++
        }
        

        k++
    }

    for i < mid {
        temp[k] = nums[i]
        i++
        k++
    }

    for j < end {
        temp[k] = nums[j]
        j++
        k++
    }

    copy(nums[start:end], temp)
}