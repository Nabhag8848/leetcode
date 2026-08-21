func permute(nums []int) [][]int {
    ans := make([][]int, 0)
    ds := make([]int, 0)
    freq := make([]bool, len(nums))

    recursive(nums, ds, freq, &ans)
    return ans
}

func recursive(nums []int, ds []int, freq []bool, ans *[][]int) {
    if len(ds) == len(nums) {
        *ans = append(*ans, ds)
        return
    }

    for i:=0; i < len(freq);i++ {
        if !freq[i] {
            freq[i] = true
            temp := append(ds, nums[i])
            recursive(nums, temp, freq, ans)
            freq[i] = false
        }
    }
}