func candy(ratings []int) int {
    left := make([]int, len(ratings))

    for i := range ratings {
        if i > 0 && ratings[i] > ratings[i - 1] {
            left[i] = left[i - 1] + 1
        } else {
            left[i] = 1
        }
    }

    right := 1
    sum := left[len(ratings) - 1]

    for i := len(ratings) - 2; i >= 0; i-- {
        if i < len(ratings) && ratings[i] > ratings[i + 1] {
            right++
        } else {
            right = 1
        }

        sum = sum + max(left[i], right)
    }

    return sum
}
