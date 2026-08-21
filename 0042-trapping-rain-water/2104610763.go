func trap(height []int) int {
    result := 0
    prefix_max := findPrefixMax(height)
    suffix_max := findSuffixMax(height)

    for i := range height {
        if height[i] < prefix_max[i] && height[i] < suffix_max[i] {
            result += (min(prefix_max[i], suffix_max[i]) - height[i])
        }
    }

    return result
}

func findPrefixMax(height []int) []int {
    prefix_max := make([]int, len(height))
    prefix_max[0] = height[0]

    for i := 1; i < len(prefix_max); i++ {
        prefix_max[i] = max(height[i], prefix_max[i - 1])
    }

    return prefix_max
}

func findSuffixMax(height []int) [] int {
    suffix_max := make([]int, len(height))
    suffix_max[len(suffix_max) - 1] = height[len(height) -  1]

    for i := len(suffix_max) - 2; i >= 0; i-- {
        suffix_max[i] = max(height[i], suffix_max[i + 1])

    }

    return suffix_max
}


