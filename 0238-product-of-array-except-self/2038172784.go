func productExceptSelf(nums []int) []int {
    answer := make([]int, len(nums))

    for i := range answer {
        answer[i] = 1
    }

    prefix := 1

    for i:=0; i< len(answer); i++ {
        answer[i] *= prefix
        prefix *= nums[i]
    }

    suffix := 1

    for i:= len(answer) - 1; i >= 0; i-- {
        answer[i] *= suffix
        suffix *= nums[i]
    }

    return answer
}

/*
   [1, 2, 3, 4]

   1 [1,1,1,1] 1
      1,1,2,6
      24 12 8 6  
*/