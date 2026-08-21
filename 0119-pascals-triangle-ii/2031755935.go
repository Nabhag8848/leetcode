func getRow(rowIndex int) []int {
    res := make([]int, 0, rowIndex + 1)
   
    value := 1
    res = append(res, 1)

    if (rowIndex > 0) {
        for i:=1;i <= rowIndex; i++ {
            value = value * ((rowIndex + 1) - i)
            value = value / i

            res = append(res, value)
        }
    }

    

    return res
}