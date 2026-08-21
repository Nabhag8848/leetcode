func maxVowels(s string, k int) int {
    count := 0

    for i := 0; i < k; i++ {
        if isVowel(s[i]) {
            count++
        }
    }

    maxi := count

    for i := 1; i <= len(s)-k; i++ {
        if isVowel(s[i-1]) {
            count--
        }
        if isVowel(s[i+k-1]) {
            count++
        }
        if count > maxi {
            maxi = count
        }
    }

    return maxi
}

func isVowel(c byte) bool {
    switch c {
    case 'a', 'e', 'i', 'o', 'u':
        return true
    }
    return false
}