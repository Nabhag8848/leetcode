func xorOperation(n int, start int) int {
    xor := 0
    for i := range n {
        xor ^= (start + 2 * i)
    }

    return xor
}