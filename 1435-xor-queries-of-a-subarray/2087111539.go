func xorQueries(arr []int, queries [][]int) []int {
    answer := make([]int, len(queries))

    for idx := range queries {
        left := queries[idx][0]
        right := queries[idx][1]

        xor := 0

        for left != right + 1 {
            xor ^= arr[left]
            left++
        }

        answer[idx] = xor
    }

    return answer
}