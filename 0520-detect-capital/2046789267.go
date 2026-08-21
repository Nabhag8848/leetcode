func detectCapitalUse(word string) bool {
    if word == strings.ToLower(word) {
        return true
    }

    if unicode.IsLower(rune(word[0])) {
        return false
    }

    if unicode.IsUpper(rune(word[0])) {
        str := word[1:]

        if str == strings.ToLower(str) || str == strings.ToUpper(str)  {
            return true
        }
    }

    return false
}