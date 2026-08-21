func letterCombinations(digits string) []string {
    b := []byte(digits)
    hash_map := map[byte][]byte{
    	'2': {'a', 'b', 'c'},
    	'3': {'d', 'e', 'f'},
    	'4': {'g', 'h', 'i'},
    	'5': {'j', 'k', 'l'},
    	'6': {'m', 'n', 'o'},
    	'7': {'p', 'q', 'r', 's'},
    	'8': {'t', 'u', 'v'},
    	'9': {'w', 'x', 'y', 'z'},
    }

    return helper(b, "", 0, hash_map)
}

func helper(b []byte, res string, idx int, hash_map map[byte][]byte) []string {
    if idx == len(b) {
        return []string{res}
    }

    var result []string
    values := hash_map[b[idx]]
    for _,ch := range values {
       result = append(result, helper(b, res + string(ch), idx + 1, hash_map)...)
    }

    return result
}



/*
                           root
            a                b.          c
        ad  ae  af      bd   be  bf   cd ce cf 

*/