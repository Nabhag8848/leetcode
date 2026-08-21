func checkInclusion(s1 string, s2 string) bool {
    if len(s2) < len(s1) {
        return false
    }

    hash_map := make(map[byte]int)
    
    for i := range s1 {
        hash_map[s1[i]]++
    }

    hash_map_substring := make(map[byte]int) 

    for i:=0;i < len(s1);i++ {
        hash_map_substring[s2[i]]++
    }

    for i := 1; i <= len(s2) - len(s1);i++ {
        if maps.Equal(hash_map_substring, hash_map) {
            return true
        }

        value := hash_map_substring[s2[i - 1]]
        if value == 1 {
            delete(hash_map_substring, s2[i - 1])
        } else {
            hash_map_substring[s2[i - 1]]--
        }

        hash_map_substring[s2[i + len(s1) - 1]]++
    }   

    if maps.Equal(hash_map_substring, hash_map) {
        return true
    }


    return false
}