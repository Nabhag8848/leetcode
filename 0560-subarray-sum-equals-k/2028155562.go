func subarraySum(nums []int, k int) int {
    n := len(nums)
    sum := 0
    result := 0
    for i := 0; i < n; i++ {
        sum += nums[i]
    }
    if sum == k {
        result++
    }

    jth := n - 1
    for window := n - 1; window > 0; window-- {
        sum -= nums[jth] // remove the old last element first...
        jth--            // ...then shrink
        x, y := 0, jth
        windowSum := sum
        for y < n {
            if windowSum == k {
                result++
            }
            y++
            if y < n {
                windowSum = windowSum - nums[x] + nums[y] // x is still the old left edge
                x++
            }
        }
    }
    return result
}

/*
    5 4 3 2 1
    i       j
    i.    j
      i.    j
    i.   j
    sum_of_whole_array

    j--
    i++

    sub_array := n -> 0

    decrease sum by element jth

    i++ decrement ith element and increment jth element

*/