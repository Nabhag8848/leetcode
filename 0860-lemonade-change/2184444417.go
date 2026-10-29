func lemonadeChange(bills []int) bool {
    hash_map := map[int]int{
        5: 0,
        10: 0, 
        20: 0,
    }  

    for _, bill := range bills {
        net := bill - 5
        ten := hash_map[10]
        five := hash_map[5]

        switch net {
            case 15:
                if ten > 0 && five > 0 {
                    hash_map[10]--
                    hash_map[5]--
                } else if five > 2 {
                    hash_map[5]--
                    hash_map[5]--
                    hash_map[5]--
                } else {
                    return false
                }

            case 5: 
                if five > 0 {
                    hash_map[5]--
                } else {
                    return false
                }

            default:

        }

        hash_map[bill]++
    }

    return true
}
