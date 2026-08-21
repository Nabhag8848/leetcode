func reversePrefix(word string, ch byte) string {
    firstOccurence := strings.IndexByte(word, ch)

    if firstOccurence == -1 {
        return word
    }

    var builder strings.Builder

    for i := firstOccurence; i >= 0; i-- {
        builder.WriteString(string(word[i]))
    }

    for i := firstOccurence + 1; i < len(word); i++ {
        builder.WriteString(string(word[i]))
    }

    return builder.String()
}