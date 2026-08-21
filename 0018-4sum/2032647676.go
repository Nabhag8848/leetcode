func fourSum(nums []int, target int) [][]int {
        sort.Ints(nums)
        result := make([][]int, 0)
        n := len(nums)
        for i := 0; i < n-3; i++ {
            if i > 0 && nums[i] == nums[i-1] {
                continue
            }

            for j := i + 1; j < n-2; j++ {
                if j > i+1 && nums[j] == nums[j-1] {
                    continue
                }

                k := j + 1
                h := n - 1

                for k < h {
                    sum := nums[i] + nums[j] + nums[k] + nums[h]

                    if sum < target {
                        k++
                    } else if sum > target {
                        h--
                    } else {
                        result = append(result,
                            []int{nums[i], nums[j], nums[k], nums[h]})

                        k++
                        h--

                        for k < h && nums[k] == nums[k-1] {
                            k++
                        }

                        for k < h && nums[h] == nums[h+1] {
                            h--
                        }
                    }
                }
            }
        }

        return result
}

