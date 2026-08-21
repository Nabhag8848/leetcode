func removeDuplicates(nums []int) int {
    maps := make(map[int]bool)

    for i:=0; i < len(nums); i++ {
        maps[nums[i]] = true
    }

    set := make([]int, 0, len(maps))

    for k := range maps {
        set = append(set, k)
    }

    sort.Ints(set)

    for i:=0; i < len(set); i++ {
        nums[i] = set[i]
    }

    return len(set)
}