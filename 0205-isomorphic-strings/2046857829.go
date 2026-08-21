func isIsomorphic(s string, t string) bool {
    s_arr := []byte(s)
    r := []byte(t)

    hash_map := make(map[byte]byte)
    set := make(map[byte]struct{})

    for i := range r {
        if value, is_exist := hash_map[r[i]]; is_exist {
            if value != s_arr[i] {
                return false
            }
        } else {
            if _, is_exist_in_set := set[s_arr[i]]; is_exist_in_set {
                return false
            }
            
            hash_map[r[i]] = s_arr[i]
            set[s_arr[i]] = struct{}{}
        }
    }

    return true
}