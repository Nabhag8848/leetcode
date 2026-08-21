func threeSum(nums []int) [][]int {
    sort.Ints(nums)
    arr := make([][]int, 0)
    for i := 0; i < len(nums); i++ {
        if (i > 0 && nums[i] == nums[i - 1]) { continue }
        j := i + 1
        k := len(nums) - 1
        for j < k {
            sum := nums[i] + nums[j] + nums[k]
            if (sum < 0) {
                j++
            } else if sum > 0 {
                k--
            } else {
                arr = append(arr, []int{nums[i], nums[j], nums[k]})
                j++
                k--
                for j < k && nums[j] == nums[j - 1] {
                    j++;
                }
                 for j < k && nums[k] == nums[k + 1] {
                    k--;
                }
            }
        }
    }

    return arr
}

