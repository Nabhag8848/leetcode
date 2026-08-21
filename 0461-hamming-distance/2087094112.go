func hammingDistance(x int, y int) int {
    start := x
    goal := y
    count := 0
    power := 0

    for start != goal {
        bit_s := (1 << power) & start
        bit_g :=  (1 << power) & goal

        if bit_s != bit_g {
            count++
            start = (1 << power) ^ start
        }

        power++
    }
    
    fmt.Println(math.MaxInt32)

    return count
}