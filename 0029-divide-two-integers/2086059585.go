func divide(dividend int, divisor int) int {
   sign := true

   if (dividend >= 0 && divisor < 0) || (dividend < 0 && divisor > 0 ){
      sign = !sign
   }

   var n int  = AbsInt(dividend)
   var d int  = AbsInt(divisor)
   var q int = 0

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
      return q
   }

   return -q

}

func AbsInt(x int) int {
    if x < 0 {
        return -x
    }

    return x
}   