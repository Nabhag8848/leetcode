func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    if len(nums2) < len(nums1) {
        return findMedianSortedArrays(nums2, nums1)
    }

    low := 0
    high := len(nums1)

    for low <= high {
        cut1 := (low + high) / 2
        cut2 := (len(nums1) + len(nums2) + 1) / 2 - cut1

        var left1 int = math.MinInt32
        var left2 int = math.MinInt32
        var right1 int = math.MaxInt32
        var right2 int = math.MaxInt32

        if cut1 > 0 {
            left1 = nums1[cut1 - 1]
        }

        if cut2 > 0 {
            left2 = nums2[cut2 - 1]
        }

        if cut1 < len(nums1) {
            right1 = nums1[cut1]
        }

        if cut2 < len(nums2) {
            right2 = nums2[cut2]
        }

        if left1 <= right2 && left2 <= right1 {
            if (len(nums1) + len(nums2)) % 2 == 0 {
                return float64(float64(max(left1, left2) + min(right1, right2)) / 2.0)
            } else {
                return float64(max(left1, left2))
            }
        } else if left1 > right2 {
            high = cut1 - 1
        } else {
            low = cut1 + 1
        }
    }

    return -1
}
