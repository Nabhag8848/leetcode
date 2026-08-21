func trap(height []int) int {
    left_max := 0
    right_max := 0
    total := 0

    left := 0 
    right := len(height) - 1

    for left < right {
        if height[left] <= height[right] {
            if left_max > height[left] {
                total += (left_max - height[left])
            } else {
                left_max = height[left]
            }

            left++
        } else {
            if right_max > height[right] {
                total += (right_max - height[right])
            } else {
                right_max = height[right]
            }

            right--
        }
    }

    return total
}

/*
    0 1 0 2 1 0 1 3 2 1 2 1
    l                     r

    leftMax = 0 rightMax = 0 total = 0

*/