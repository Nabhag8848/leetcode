func maxDepth(s string) int {
   maxi := 0
   count := 0

   for i := range s {
      if s[i] == '(' {
        count++
      } else if s[i] == ')' {
        count--
      }

     maxi = max(count, maxi)
   } 

   return maxi
}