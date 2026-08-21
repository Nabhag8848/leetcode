func xorAllNums(nums1 []int, nums2 []int) int {
    len1 := len(nums1)
    len2 := len(nums2)
    isEven1 := len1 % 2 == 0
    isEven2 := len2 % 2 == 0  

    if !isEven1 && isEven2 {
        return xor(nums2)
    } else if isEven1 && !isEven2 {
        return xor(nums1)
    } else if !isEven1 && !isEven2 {
        xor1 := xor(nums1)
        xor2 := xor(nums2)
        return xor1 ^ xor2
    }

    return 0
}

func xor(nums []int) int {
    result := 0
    for i := range nums {
        result ^= nums[i]
    }
    return result
}

/*
    (xi ^ yi)
    (x1 ^ y1) ^ (x1 ^ y2) ^ (x1 ^ y3) ^ (x2 ^ y1) ^ (x2 ^ y2) ^ (x2 ^ y3) == x1 ^ x2
    (x1 ^ y1) ^ (x1 ^ y2) ^ (x2 ^ y1) ^ (x2 ^ y2) ^  (x3 ^ y1) ^ (x3 ^ y2) == y1 ^ y2
    (x1 ^ y1) ^ (x1 ^ y2) ^ (x1 ^ y3) == (x1 ^ y1 ^ y2 ^ y3)
*/

