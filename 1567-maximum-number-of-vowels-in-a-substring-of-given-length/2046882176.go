func maxVowels(s string, k int) int {
   count := 0
   vowels := []string{"a", "e", "i", "o", "u"}
   maxi := 0

   for i:=0; i < k; i++{
      if slices.Contains(vowels, string(s[i])) {
        count++
      }
   }

   maxi = count

   for i:=1; i <= len(s) - k; i++ {
       if slices.Contains(vowels, string(s[i - 1])) {
          count--
       }

       if slices.Contains(vowels, string(s[i + k - 1])) {
          count++
       }

       maxi = max(count, maxi)
   }

   return maxi
}