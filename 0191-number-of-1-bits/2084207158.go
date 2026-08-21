func hammingWeight(n int) int {
    count := 0

    for n > 1 {
        if n % 2 == 1 {
            count++
        }
        n = n / 2
    }

    if n == 1 {
        count++
    }

    return count
}