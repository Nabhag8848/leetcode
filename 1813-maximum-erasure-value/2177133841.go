func maximumUniqueSubarray(nums []int) int {
    hash_map := make(map[int]struct{})

    l := 0
    r := 0
    sum := 0
    maxSum := 0

    for r < len(nums) {
        if _, ok := hash_map[nums[r]]; ok {
            delete(hash_map, nums[l])
            sum = sum - nums[l]
            l++
        } else {
            sum += nums[r]
            hash_map[nums[r]] = struct{}{}
            r++

            maxSum = max(maxSum, sum)
        }
    }

    return maxSum
}