func candy(ratings []int) int {
    i := 1
    n := len(ratings)

    sum := 1

    for i < n {
        if i < n && ratings[i] == ratings[i - 1] {
            sum = sum + 1
            i++
            continue
        }

        peak := 1

        for i < n && ratings[i] > ratings[i - 1] {
            peak++
            sum = sum + peak
            i++
        }

        down := 1

        for i < n && ratings[i] < ratings[i - 1] {
            sum = sum + down
            i++
            down++
        }

        if down > peak {
            sum += (down - peak)
        }

    }

    return sum
}
