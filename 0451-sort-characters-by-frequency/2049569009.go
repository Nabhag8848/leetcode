func frequencySort(s string) string {
    hash_map := make(map[byte]int)

    for i := range s {
        hash_map[s[i]]++
    }

    type pair struct  {
        key byte
        value int
    }

    pairs := make([]pair, 0, len(hash_map))

    for k,v := range hash_map {
        pairs = append(pairs, pair{k, v})
    }

    sort.Slice(pairs, func (i, j int) bool {
        return pairs[i].value > pairs[j].value
    })

    var builder strings.Builder

    for _,pair_val := range pairs {
        for i:=0; i < pair_val.value;i++ {
            builder.WriteByte(pair_val.key)
        }
    }

    return builder.String()
}