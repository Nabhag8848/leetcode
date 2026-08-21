func reverseWords(s string) string {
   s = strings.TrimSpace(s)
   low:=len(s) - 1
   high := len(s) - 1

   var builder strings.Builder

   for low >= 0 {
        for low >= 0 && s[low] != ' ' {
            low--
        }

        builder.WriteString(s[low + 1:high + 1])
        high = low

        for high >= 0 && s[high] == ' ' {
            high--
        }

        low = high

        if low >= 0 {
            builder.WriteString(" ")
        }
   }

   return builder.String()
}