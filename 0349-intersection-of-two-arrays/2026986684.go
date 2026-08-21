func intersection(nums1 []int, nums2 []int) []int {
    hash_map := make(map[int]bool)
    var result []int
    for i:=0; i < len(nums1); i++ {
        hash_map[nums1[i]] = false
    }

    for j:=0; j < len(nums2); j++ {
        is_append,exist := hash_map[nums2[j]]
        if exist {
            if (!is_append && exist) {
                result = append(result, nums2[j])
                hash_map[nums2[j]] = true
            }
        }
    }

    return result
}