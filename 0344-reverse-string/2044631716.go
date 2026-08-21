func reverseString(s []byte)  {
    i := 0
    j := len(s) - 1

    for i < j {
        x := s[i]
        s[i] = s[j]
        s[j] = x
        i++
        j--
    }
}