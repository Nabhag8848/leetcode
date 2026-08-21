func twoSum(nums []int, target int) []int {
    hash_map := make(map[int]int)

    for i:=0; i< len(nums); i++ {
        j,is_exist := hash_map[target - nums[i]]

        if is_exist { 
            return []int{i, j}
        }

        hash_map[nums[i]] = i
    }

    return []int{-1, -1}

}