func isPowerOfTwo(n int) bool {
    if n == 0 {
        return false
    }

    
    x := n & (n - 1)

    if x == 0 {
        return true
    }

    return false
}