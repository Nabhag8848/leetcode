func closeStrings(word1 string, word2 string) bool {
    if len(word1) != len(word2) {
        return false
    }

    occ1 := make([]int, 26)
    occ2 := make([]int, 26)

    for i := range word1 {
        occ1[word1[i] - 'a']++
    }

    for i := range word2 {
        occ2[word2[i] - 'a']++
    }

    for i := range 26 {
        if (occ1[i] == 0 && occ2[i] != 0) {
            return false
        }

        if (occ2[i] == 0 && occ1[i] != 0) {
            return false
        }
    }


    slices.Sort(occ1)
    slices.Sort(occ2)

    ans := slices.Equal(occ1, occ2)

    return ans

}

/*

    a b c -> b a c -> b c a

    b c a <- target

    c a b b b a  -> b a c c c a -> a b c c c a -> true

    a b b c c c


    if l(w1) != l(w2)


    hash_map1<byte,int>
    hash_map2<byte,int>

    c a b b b a

    O(N^2), O(2N)

    hash_map1<int, []int> -> O(N)  hash_map2<byte,int> -> O(N)

    1 -> [c]               a -> 1, b -> 2, c -> 3
    2 -> [a]
    3 -> [b]

    1 - 26    

    [2, 3, 1, 0 ...] ->   [0,0,0,0,0,0,0,0... 1, 2, 3]
    
    cabbba -> abbcc (<-)
*/