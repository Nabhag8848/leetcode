func sortArray(nums []int) []int {
    mergeSort(&nums, 0, len(nums))
    return nums
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

func mergeSort(nums *[]int, start int, end int) {
    length := end - start
    if (length <= 1) {
        return
    }

    mid := start + (end-start)/2
    mergeSort(nums, start, mid)
    mergeSort(nums, mid, end)

    mergetwoSortedArray(nums, start, mid, end)
}

func mergetwoSortedArray(nums *[]int, start int, mid int, end int) {
    arr := make([]int, end - start)

    i := start
    j := mid
    k := 0

    for i < mid && j < end {
        if (*nums)[i] < (*nums)[j] {
            arr[k] = (*nums)[i]
            i++
        }else {
            arr[k] = (*nums)[j]
            j++
        }

        k++
    }

    for i < mid {
        arr[k] = (*nums)[i]
        k++
        i++
    }


    for j < end {
        arr[k] = (*nums)[j]
        k++
        j++
    }

    h := 0
    for x:=start; x < end; x++ {
        (*nums)[x] = arr[h]
        h++
    }
}