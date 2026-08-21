func sortSentence(s string) string {
    words := strings.Fields(s)
    result := make([]string, len(words))

    for i := range words {
        word := words[i]
        index,_ := strconv.Atoi(word[(len(word) - 1):])
        result[index - 1] = word[:(len(word) - 1)]
    }

    return strings.Join(result, " ")
}