func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    total := len(nums1) + len(nums2)

    if total%2 == 1 {
        return float64(kth(nums1, nums2, total/2))
    }

    left := kth(nums1, nums2, total/2-1)
    right := kth(nums1, nums2, total/2)

    return float64(left+right) / 2.0
}

func kth(nums1, nums2 []int, k int) int {
    low := getMin(nums1, nums2)
    high := getMax(nums1, nums2)

    for low <= high {
        mid := low + (high-low)/2

        count :=
            binary_search(nums1, mid) +
                binary_search(nums2, mid)

        // count = number of elements <= mid

        if count > k {
            high = mid - 1
        } else {
            low = mid + 1
        }
    }

    return low
}

func binary_search(nums []int, target int) int {
    low := 0
    high := len(nums) - 1

    for low <= high {
        mid := low + (high-low)/2

        if nums[mid] > target {
            high = mid - 1
        } else {
            low = mid + 1
        }
    }

    return low
}

func getMin(nums1, nums2 []int) int {
    if len(nums1) == 0 {
        return nums2[0]
    }
    if len(nums2) == 0 {
        return nums1[0]
    }

    if nums1[0] < nums2[0] {
        return nums1[0]
    }
    return nums2[0]
}

func getMax(nums1, nums2 []int) int {
    if len(nums1) == 0 {
        return nums2[len(nums2)-1]
    }
    if len(nums2) == 0 {
        return nums1[len(nums1)-1]
    }

    if nums1[len(nums1)-1] > nums2[len(nums2)-1] {
        return nums1[len(nums1)-1]
    }
    
    return nums2[len(nums2)-1]
}