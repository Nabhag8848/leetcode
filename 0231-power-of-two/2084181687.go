func isPowerOfTwo(n int) bool {
    for i := range 31 {
        bit := (1 << i) & n 
        if bit != 0 {
            n = (1 << i) ^ n

            if n == 0 {
                return true
            }

            return false
        }
    }

    return false
}