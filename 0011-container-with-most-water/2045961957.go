func maxArea(height []int) int {
    i := 0
    j := len(height) - 1
    area := 0

    for i < j {
        w := j - i
        h := min(height[j], height[i])

        area = max(w * h, area)

        if height[j] < height[i] {
            j--
        } else {
            i++
        }

    }

    return area
}
