func singleNumber(nums []int) int {
   result := 0

   for i := range 32 {
      count := 0

      for _,value := range nums {
        bit := value & (1 << i)
        if bit != 0 {
            count++
        }
      }

      if count % 3 != 0 {
        if i == 31 {
            result = result - (1 << i)
        } else {
            result = result | (1 << i)
        }
      }

   }

   return result
}


