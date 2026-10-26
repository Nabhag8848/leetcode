func numSubarraysWithSum(nums []int, goal int) int {
    hash_map := make(map[int]int)
    sum := 0
    hash_map[sum] = 1
    result := 0

    for i := range nums {
        sum += nums[i]

        if val, ok := hash_map[sum - goal]; ok {
            result = result + val
        }

        hash_map[sum]++
    }   

    return result
}


  