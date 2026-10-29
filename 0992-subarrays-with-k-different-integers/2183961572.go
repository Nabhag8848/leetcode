func subarraysWithKDistinct(nums []int, k int) int {
    return subarraysWithKLessDistinct(nums ,k) - subarraysWithKLessDistinct(nums, k - 1)
}

func subarraysWithKLessDistinct(nums []int, k int) int {
    l := 0
    r := 0

    result := 0
    hash_map := make(map[int]int)

    for r < len(nums) {
        hash_map[nums[r]]++

        for len(hash_map) > k {
            hash_map[nums[l]]--

            if val, _ := hash_map[nums[l]]; val == 0 {
                delete(hash_map, nums[l])
            } 

            l++
        }

        result = result + (r - l + 1)
        r++
    }

    return result
}
