// func majorityElement(nums []int) []int {
//     ele1 := math.MinInt32
//     ele2 := math.MinInt32
//     count1, count2 := 0, 0

//     res := make([]int, 0, 2)

//     for i:=0; i < len(nums);i++ {

//         if count1 == 0 && ele2 != nums[i]  {
//             ele1 = nums[i]
//             count1++
//         }else if count2 == 0 && ele1 != nums[i]  {
//             ele2 = nums[i]
//             count2++
//         }else if nums[i] == ele1 {
//             count1++
//         } else if nums[i] == ele2 {
//             count2++
//         } else {
//             count1--
//             count2--
//         }
//     }

//     count := int(len(nums) / 3) + 1

//     if count1 >= count {
//         res = append(res, ele1)
//     }

//     if count2 >= count {
//         res = append(res, ele2)
//     }

//     return res

// }

func majorityElement(nums []int) []int {
    ele1 := 0
    ele2 := 0
    count1, count2 := 0, 0

    for _, num := range nums {

        if num == ele1 {
            count1++
        } else if num == ele2 {
            count2++
        } else if count1 == 0 {
            ele1 = num
            count1 = 1
        } else if count2 == 0 {
            ele2 = num
            count2 = 1
        } else {
            count1--
            count2--
        }
    }

    // verification step (required)
    count1, count2 = 0, 0
    for _, num := range nums {
        if num == ele1 {
            count1++
        } else if num == ele2 {
            count2++
        }
    }

    res := []int{}
    n := len(nums)

    if count1 > n/3 {
        res = append(res, ele1)
    }
    if count2 > n/3 {
        res = append(res, ele2)
    }

    return res
}