func mergeAlternately(word1 string, word2 string) string {
    var builder strings.Builder
    count := 0
    total_length := len(word1) + len(word2)
    i := 0
    j := 0

    for count < total_length {
        if count % 2 == 0 {
            if i < len(word1) {
                builder.WriteString(string(word1[i]))
                i++
            } else {
                builder.WriteString(string(word2[j]))
                j++
            }   
        } else {
            if j < len(word2) {
                builder.WriteString(string(word2[j]))
                j++
            } else {
                builder.WriteString(string(word1[i]))
                i++
            }
        }

        count++
    }

    return builder.String()
}