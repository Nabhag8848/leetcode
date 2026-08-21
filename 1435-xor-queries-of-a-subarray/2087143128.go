func xorQueries(arr []int, queries [][]int) []int {
    answer := make([]int, len(queries))
    for i := range arr {
        if i > 0 {
            arr[i] = arr[i] ^ arr[i - 1]
        }
    }

    for i := range queries {
        left := queries[i][0]
        right := queries[i][1]

        if left > 0 {
            answer[i] = (arr[left - 1] ^ arr[right])
        } else {
            answer[i] = arr[right]
        }
    }

    return answer
}

/*

    0 - len(arr)
    x ^ y ^ z ^ w
    x ^ x ^ y

    1 - 2 (y ^ z ^ x ^ x ^ y)

*/