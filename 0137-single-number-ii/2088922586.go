func singleNumber(nums []int) int {
    result := int32(0) 

    for i := 0; i < 32; i++ {
        bitSum := 0
        for _, num := range nums {
            bitSum += (num >> i) & 1
        }
        if bitSum%3 != 0 {
            result |= 1 << i
        }
    }

    return int(result)
}