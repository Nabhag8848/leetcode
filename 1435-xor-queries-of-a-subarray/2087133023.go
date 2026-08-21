func xorQueries(arr []int, queries [][]int) []int {
    answer := make([]int, len(queries))
    prefix_sum := make([]int, len(arr))

    xor := 0

    for i := range arr {
        if i > 0 {
            prefix_sum[i] = xor ^ arr[i - 1]
            xor = xor ^ arr[i - 1]
        }

    }

    for i := range queries {
        left := queries[i][0]
        right := queries[i][1]

        answer[i] = (prefix_sum[left] ^ prefix_sum[right] ^ arr[right])
    }

    return answer
}

/*

    0 - len(arr)
    x ^ y ^ z ^ w
    x ^ x ^ y

    1 - 2 (y ^ z ^ x ^ x)

*/