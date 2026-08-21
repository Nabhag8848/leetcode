func beautySum(s string) int {
   sum := 0

   for i := range s {
      freq := make([]int, 26)
      for j:=i; j < len(s); j++ {
        freq[s[j] - 'a']++
        sum += (getMaxCount(freq) - getMinCount(freq))
      }
   }

   return sum 
}

func getMaxCount(freq []int)int {
    maxi := 0

    for i := range freq {
        maxi = max(freq[i], maxi)
    }

    return maxi
}

func getMinCount(freq []int)int {
    mini := math.MaxInt32

    for i := range freq {
        if freq[i] != 0 {   
            mini = min(freq[i], mini)
        }
    }

    return mini
}

