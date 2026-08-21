func findMissingAndRepeatedValues(grid [][]int) []int {
    xor := 0
    for i:=0; i<len(grid); i++ {
        for j:=0; j < len(grid[0]); j++ {
            xor ^= grid[i][j]
        }
    }

    xorUptoN2 := 0
    for i:=1; i <= len(grid) * len(grid[0]); i++ {
        xorUptoN2 ^= i
    }

    xor ^= xorUptoN2
    bitOne := xor & ^(xor - 1)
    one := 0
    zero := 0

    for i:=0; i < len(grid); i++{
        for j:=0; j < len(grid[0]); j++ {
            if (grid[i][j] & bitOne) != 0 {
                one ^= grid[i][j]
            } else {
                zero ^= grid[i][j]
            }
        }
    }

    n2 := len(grid) * len(grid[0])

    for i := 1; i <= n2; i++ {
        if (i & bitOne) != 0 {
            one ^= i
        } else {
            zero ^= i
        }
    }

    counter := 0

    for i:=0; i < len(grid); i++{
        for j:=0; j < len(grid[0]); j++ {
            if grid[i][j] == one {
                counter++
            }
        }
    }

    if counter == 2 {
        return []int{one, zero}
    }

    return []int{zero, one}   
}


/*
    1 ^ 3 ^ 2 ^ 2 ^ (1 ^ 2 ^ 3 ^ 4) => 2 ^ 4
*/
