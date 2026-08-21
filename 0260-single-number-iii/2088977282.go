func singleNumber(nums []int) []int {
    var num int = 0

    for i := range nums {
        num ^= nums[i]
    }

    rightmost := (num & (num - 1)) ^ num

    xor1 := 0
    xor2 := 0

    for i := range nums {
        bit := nums[i] & rightmost
        if bit != 0 {
            xor1 ^= nums[i]
        } else {
            xor2 ^= nums[i]
        }
    }

    return []int{xor1, xor2}

}