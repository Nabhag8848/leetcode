func merge(nums1 []int, m int, nums2 []int, n int)  {
    i := m-1
    j := n-1
    k := len(nums1) - 1

    for i >= 0 && j >= 0 {

        if nums1[i] > nums2[j] {
            nums1[k] = nums1[i]
            k--
            i--
        } else if nums1[i] < nums2[j] {
            nums1[k] = nums2[j]
            k--
            j--
        } else {
            nums1[k] = nums1[i]
            i--
            k--

            nums1[k] = nums2[j]
            j--
            k--
        
        }
    }

    for i >= 0 {
        nums1[k] = nums1[i]
        k--
        i--
    }

    for j >= 0 {
        nums1[k] = nums2[j]
        j--
        k--
    }

   
}

/*
    [1, 2, 3, 0, 0, 0]
    [2, 5, 6]

    i. j
    i < j
    i == j i++
    j < i swap (i, j) i++

    [1, 2, 3, 0, 0, 0]
    [2, 5, 6]

    i=0, j=0 i++
    i=1, j=0 i++
    i =2, j=0 swap(i, j) as i > j i++ 

    i < m
    j < n -> m + j

*/