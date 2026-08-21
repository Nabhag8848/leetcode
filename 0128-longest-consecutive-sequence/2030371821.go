func longestConsecutive(nums []int) int {
    if len(nums) == 0 {
        return 0
    }

    set := make(map[int]struct{})
    for _, num := range nums {
        set[num] = struct{}{}
    }

    longest := 0

    for num := range set {
        // Only start from the beginning of a sequence
        if _, exists := set[num-1]; exists {
            continue
        }

        length := 1
        curr := num

        for {
            if _, exists := set[curr+1]; !exists {
                break
            }
            curr++
            length++
        }

        if length > longest {
            longest = length
        }
    }

    return longest
}