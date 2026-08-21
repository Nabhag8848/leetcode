func canConstruct(ransomNote string, magazine string) bool {
    freqRansomNote := make([]int, 26)
    freqMagazine := make([]int, 26)

    for i := range ransomNote {
        freqRansomNote[ransomNote[i] - 97]++
    }

    for i := range magazine {
        freqMagazine[magazine[i] - 97]++
    }

    for i := range freqMagazine {
        if freqMagazine[i] < freqRansomNote[i] {
            return false
        }
    }

    return true

}