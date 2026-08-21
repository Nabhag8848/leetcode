func hammingWeight(n int) int {
    count := 0

    for i := range 31 {
        bit := (1 << i) & n
        if  bit != 0 {
            count++
        }
    }

    return count
}