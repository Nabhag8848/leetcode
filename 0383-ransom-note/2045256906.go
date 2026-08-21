func canConstruct(ransomNote string, magazine string) bool {
    freqRansomNote := make([]int, 26)

    for i := range magazine {
        freqRansomNote[magazine[i] - 'a']++
    }

    for i := range ransomNote {
        freqRansomNote[ransomNote[i] - 'a']--
        if freqRansomNote[ransomNote[i] - 'a'] < 0 {
            return false
        }
    }

    return true

}