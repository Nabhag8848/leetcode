func isIdealPermutation(nums []int) bool {
    local := 0

    for i := 0; i < len(nums)-1; i++ {
        if nums[i] > nums[i+1] {
            local++
        }
    }

    global := countGlobalInversions(nums, 0, len(nums))

    return local == global
}

func countGlobalInversions(nums []int, start, end int) int {
    if end-start <= 1 {
        return 0
    }

    mid := start + (end-start)/2

    left := countGlobalInversions(nums, start, mid)
    right := countGlobalInversions(nums, mid, end)
    merge := mergeAndCount(nums, start, mid, end)

    return left + right + merge
}

func mergeAndCount(nums []int, start, mid, end int) int {
    temp := make([]int, end-start)

    i := start
    j := mid
    k := 0

    inversions := 0

    for i < mid && j < end {
        if nums[i] <= nums[j] {
            temp[k] = nums[i]
            i++
        } else {
            temp[k] = nums[j]

            // Every remaining element in left half
            // forms an inversion with nums[j]
            inversions += mid - i

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

    return inversions
}