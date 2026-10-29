func findContentChildren(g []int, s []int) int {
    sort.Ints(g)
    sort.Ints(s)

    i := 0
    j := 0

    result := 0

    for i < len(g) && j < len(s) {
        if s[j] >= g[i] {
            result++
            i++
        }

        j++
    }

    return result
}
