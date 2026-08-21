func frequencySort(s string) string {
    type pair struct  {
        key int32
        value int
    }

    pairs := make([]pair, 256)

    for idx,_ := range pairs {
        pairs[idx] = pair{-1, 0}
    }

    for idx,_ := range s {
        pairs[s[idx]].key = int32(s[idx])
        pairs[s[idx]].value = pairs[s[idx]].value + 1
    }

    sort.Slice(pairs, func (i, j int) bool {
        if pairs[i].value == pairs[j].value {
            return pairs[i].key < pairs[j].key
        }

        return pairs[i].value > pairs[j].value
    })

    var builder strings.Builder

    for _,pair_val := range pairs {
        if pair_val.key != -1 {
            for i:=0; i < pair_val.value;i++ {
                builder.WriteByte(byte(pair_val.key))
            }
        }

    }

    return builder.String()
}