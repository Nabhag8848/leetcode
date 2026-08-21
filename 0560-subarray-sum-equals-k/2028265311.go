func subarraySum(nums []int, k int) int {
    // better approach prefix_sum + hashmap
    hash_map := make(map[int]int)
    hash_map[0] = 1


    sum,result := 0,0

    for _,v := range nums {
        sum += v
        _, is_exist := hash_map[sum - k]

        if is_exist {
            result += hash_map[sum - k]
        }

        hash_map[sum]++
    }

    return result
    
    /*  k = 5   

        1 2 3 4 5

        0 -> 1
        1 -> 2
        3 -> 3
        6 -> 4
        10 -> 5
        15 -> 6

    */


  
}

  // brute force: 

    // n := len(nums)
    // sum := 0
    // result := 0
    // for i := 0; i < n; i++ {
    //     sum += nums[i]
    // }
    // if sum == k {
    //     result++
    // }

    // jth := n - 1
    // for window := n - 1; window > 0; window-- {
    //     sum -= nums[jth] // remove the old last element first...
    //     jth--            // ...then shrink
    //     x, y := 0, jth
    //     windowSum := sum
    //     for y < n {
    //         if windowSum == k {
    //             result++
    //         }
    //         y++
    //         if y < n {
    //             windowSum = windowSum - nums[x] + nums[y] // x is still the old left edge
    //             x++
    //         }
    //     }
    // }
    // return result

/*
    5 4 3 2 1
    i       j
    i.    j
      i.    j
    i.   j
    sum_of_whole_array

    j--
    i++

    sub_array := n -> 0

    decrease sum by element jth

    i++ decrement ith element and increment jth element

*/