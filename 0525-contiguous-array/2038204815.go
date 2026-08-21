func findMaxLength(nums []int) int {
    for i:=range nums {
        if nums[i] == 0 {
            nums[i] = -1
        }
    }

    hash_map := make(map[int]int)
    hash_map[0] = -1

    // sum of largest subarray which equal zero

    /*
        -1 1 1 1 1 1 -1 -1 -1
           
           (0, 1)
           (4  5)
           (3, 6)
           (2, 7)
           (1, 8)
    */

    prefix_sum := 0
    max_len := 0

    for i:= range nums {
        prefix_sum += nums[i]

        if idx, is_exist := hash_map[prefix_sum]; is_exist {
            max_len = int(math.Max(float64(max_len), float64(i - idx)))
        } else {
          hash_map[prefix_sum] = i
        }
    }

    return max_len
}