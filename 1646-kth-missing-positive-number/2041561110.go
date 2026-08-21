func findKthPositive(arr []int, k int) int {
    last_element := arr[len(arr) - 1]
    j:=0
    current_missing_count := 0

    for i := 1; i <= last_element; i++ {

        if arr[j] == i {
            j++
        } else {
            current_missing_count++
        }

        if current_missing_count == k {
            return i
        }
    }

    return k - current_missing_count + last_element

}

/*
    4 - (2 + 1)

    2,3,4,7,11

    1 -> arr[len(arr) - 1]

*/  
