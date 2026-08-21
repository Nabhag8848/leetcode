func commonChars(words []string) []string {
    len_words := len(words)
    freq := make([][]int, len_words)
    for i := range freq {
        freq[i] = make([]int, 26)
    }
    result := make([]string, 0)

    for i := range words {
        str := words[i]
        for _,value := range str {
            freq[i][value - 'a']++
        }
    }

    for i:=0; i < len(freq[0]); i++ {
        min := math.MaxInt32

        for j := range freq {
            if freq[j][i] < min {
                min = freq[j][i]
            }
        }

        if min > 0 {
            for idx:=0; idx < min;idx++ {
                result = append(result, string(i + 97))  
            }
        }
    }

    return result
}