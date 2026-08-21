func minBitFlips(start int, goal int) int {
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

    return count
}

/*  
    count := 0
    start != goal 

    1 0 1 0
    0 1 1 1

    0 0 1 1
    0 1 0 0

    (1 << i) & goal
    (1 << j) & start
    if != {
        start = (1 << j) ^ start
    }

*/