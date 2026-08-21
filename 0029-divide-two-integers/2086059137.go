func divide(dividend int, divisor int) int {
   sign := true

   if (dividend >= 0 && divisor < 0) || (dividend < 0 && divisor > 0 ){
      sign = !sign
   }

   var n int64  = AbsInt(dividend)
   var d int64  = AbsInt(divisor)
   var q int64 = 0

   for n >= d {
      power := 0

      for n > (d << (power + 1)) {
        power++
      }
      
      q += (1 << power)
      n = n - (d << power)
   }

   if q > math.MaxInt32  {
      if sign {
        return math.MaxInt32
      }

      if !sign {
        return math.MinInt32
      }
   }

   if sign {
      return int(q)
   }

   return int(-q)

}

func AbsInt(x int) int64 {
    if x < 0 {
        return int64(-x)
    }

    return int64(x)
}   
