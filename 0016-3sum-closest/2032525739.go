func threeSumClosest(nums []int, target int) int {
    sort.Ints(nums)
    distance := math.MaxInt32
    var ans int 
    for i := 0; i < len(nums); i++ {
        if (i > 0 && nums[i] == nums[i - 1]) { continue }
        j := i + 1
        k := len(nums) - 1
        for j < k {
            sum := nums[i] + nums[j] + nums[k]

            if (sum < target) {
                possible_distance := abs(sum - target)
                if possible_distance < distance {
                    distance = possible_distance
                    ans = sum
                }
                j++
            } else if sum > target {          
                possible_distance := abs(sum - target)
                if possible_distance < distance {
                    distance = possible_distance
                    ans = sum
                }
                k--
            } else {
                return target
            }
        }
    }

    return ans
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}