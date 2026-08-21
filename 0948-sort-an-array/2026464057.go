func sortArray(nums []int) []int {
    return mergeSort(nums)
}

func bubbleSort(nums *[]int) {
    for i := len(*nums) - 1; i > 0; i-- {
        for j := 0; j < i; j++ {
            if((*nums)[j] > (*nums)[j + 1]) {
                swap(nums, j, j + 1)
            }
        }
    }
}

func insertionSort(nums *[]int) {
    for i := 0; i < len(*nums) - 1; i++ {
        for j:=i + 1; j > 0; j-- {
            if (*nums)[j] < (*nums)[j - 1] {
                swap(nums, j, j - 1)
            }
        }
    }
}

// 2 5 3 1 ->  2 3 5 1

func selectionSort(nums *[]int) {
    var min int
    for i := 0; i < len(*nums); i++ {
        min = i
        for j := i + 1; j < len(*nums); j++ {
            if (*nums)[j] < (*nums)[min] {
                min = j
            }
        }
        
        swap(nums, i, min)
    }

}

func swap(nums *[]int, i int, j int) {
    x := (*nums)[i]
    (*nums)[i] = (*nums)[j]
    (*nums)[j] = x
}

func mergeSort(nums []int) []int{
    if (len(nums) <= 1) {
        return nums
    }

    mid := len(nums) / 2
    leftClone := make([]int, mid)
    rightClone := make([]int, len(nums) - mid)
    copy(leftClone, nums[0:mid])
    left := mergeSort(leftClone)
    copy(rightClone, nums[mid:])
    right := mergeSort(rightClone)

    return mergetwoSortedArray(left, len(left), right, len(right))
}

func mergetwoSortedArray(nums1 []int, m int, nums2 []int, n int)  []int{
    arr := make([]int, m + n)

    i := 0
    j := 0
    k := 0

    for i < m && j < n && k < m + n {
        if nums1[i] < nums2[j] {
            arr[k] = nums1[i]
            i++
        }else {
            arr[k] = nums2[j]
            j++
        }

        k++
    }

    for i < m {
        arr[k] = nums1[i]
        k++
        i++
    }


    for j < n {
        arr[k] = nums2[j]
        k++
        j++
    }

    return arr
}