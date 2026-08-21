func romanToInt(s string) int {
    hash_map := map[byte]int{
        'I':1,
        'V':5,
        'X':10,
        'L':50,
        'C':100,
        'D':500,
        'M':1000,
    }

    count := 0

    if len(s) == 1 {
        return hash_map[s[0]]
    }

    for i:=0; i < len(s) - 1; i++{
        if hash_map[s[i]] < hash_map[s[i + 1]] {
            count -= hash_map[s[i]]
        } else {
            count += hash_map[s[i]]
        }
    }

    return count + hash_map[s[len(s) - 1]]
}