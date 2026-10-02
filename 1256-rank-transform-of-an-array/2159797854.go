func arrayRankTransform(arr []int) []int {
    sorted := append([]int(nil), arr...)
	sort.Ints(sorted)

	rank := make(map[int]int)

	for _, value := range sorted {
		if _, exists := rank[value]; !exists {
			rank[value] = len(rank) + 1
		}
	}

	result := make([]int, len(arr))

	for i, value := range arr {
		result[i] = rank[value]
	}

	return result
}