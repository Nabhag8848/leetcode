func letterCasePermutation(s string) []string {
    b := []byte(s)
    return helper("", b, 0)
}

func helper(result string, b []byte, index int) []string {
    if index == len(b) {
        return []string{result}
    }

    var response []string

    if unicode.IsLetter(rune(b[index])) {
        response = append(response, helper(result + strings.ToUpper(string(b[index])), b, index + 1)...)
        response = append(response, helper(result + strings.ToLower(string(b[index])), b, index + 1)...)
    } else {
        response = append(response, helper(result + string(b[index]), b, index + 1)...)
    }


    return response
}