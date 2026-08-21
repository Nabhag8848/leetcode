func merge(nums1 []int, m int, nums2 []int, n int)  {
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
 
    for x := 0; x < m + n; x++ {
        nums1[x] = arr[x]
    }
}