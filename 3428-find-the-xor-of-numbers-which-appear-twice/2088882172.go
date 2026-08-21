func duplicateNumbersXOR(nums []int) int {
    set := make(map[int]int)

    for i := range nums {
        set[nums[i]]++
    }

    xor := 0

    for key,value := range set {
        if value == 2 {
            xor ^= key
        }
    }

    return xor
}