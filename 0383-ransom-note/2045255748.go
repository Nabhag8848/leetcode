func canConstruct(ransomNote string, magazine string) bool {
    freqRansomNote := make([]int, 26)

    for i := range magazine {
        freqRansomNote[magazine[i] - 97]++
    }

    for i := range ransomNote {
        freqRansomNote[ransomNote[i] - 97]--
        if freqRansomNote[ransomNote[i] - 97] < 0 {
            return false
        }
    }

    return true

}