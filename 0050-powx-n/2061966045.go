func myPow(x float64, n int) float64 {
   if n < 0 {
      return 1 / helper(x, -n)
   }

   return helper(x, n)
}

func helper(x float64, n int) float64 {
    if n == 0 {
        return 1
    }

    half := helper(x, n/2)

    if n%2 == 0 {
        return half * half
    }
    return half * half * x
}