func getRow(rowIndex int) []int {
    res := make([]int, 0, rowIndex + 1)
    for i:=0; i <= rowIndex; i++ {
       res = append(res, binomialCoefficient(rowIndex, i))
    } 

    return res
}

func binomialCoefficient(n, k int) int {
    res := 1

    for i := 1; i <= k; i++ {
        res = res * (n - i + 1)
        res = res / i
    }

    return res
}