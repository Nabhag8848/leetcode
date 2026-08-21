func letterCasePermutation(s string) []string {
    b := []byte(s)
    hash_map := make(map[string]bool)
    return helper("", b, 0, hash_map)
}

func helper(result string, b []byte, index int, hash_map map[string]bool) []string {
    if index == len(b) {
        if _, ok := hash_map[result]; !ok {
            hash_map[result] = true
            return []string{result}
        }

        return []string{}

    }

    var response []string

    response = append(response, helper(result + strings.ToUpper(string(b[index])), b, index + 1, hash_map)...)
    response = append(response, helper(result + strings.ToLower(string(b[index])), b, index + 1, hash_map)...)

    return response
}