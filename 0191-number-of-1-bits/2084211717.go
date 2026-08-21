func hammingWeight(n int) int {
    count := 0

    for n > 1 {
        count += n & 1
        n = n >> 1
    }

    if n == 1 {
        count++
    }

    return count
}