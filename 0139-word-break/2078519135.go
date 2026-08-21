func wordBreak(s string, wordDict []string) bool {
	set := make(map[string]struct{})
	memo := make([]int8, len(s)+1) // 0 = unvisited, 1 = true, 2 = false

	for _, element := range wordDict {
		set[element] = struct{}{}
	}
	return helper(s, 0, set, memo)
}

func helper(s string, idx int, set map[string]struct{}, memo []int8) bool {
	if idx == len(s) {
		return true
	}
    
	if memo[idx] != 0 {
		return memo[idx] == 1
	}

	for i := idx; i < len(s); i++ {
		word := s[idx : i+1]
		if _, ok := set[word]; ok {
			if helper(s, i+1, set, memo) {
				memo[idx] = 1
				return true
			}
		}
	}

	memo[idx] = 2
	return false
}