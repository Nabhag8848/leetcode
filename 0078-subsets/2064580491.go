func subsets(nums []int) [][]int {
    result := make([][]int, 0)
    nums_len := len(nums)

    for i:=0;i < (1 << nums_len); i++ {
        res := []int{}
        for j:=0;j < nums_len;j++ {
            if (i & (1 << j)) != 0 {
                res = append(res, nums[j])
            }
        }  

        result = append(result, res) 
    }

    return result
}