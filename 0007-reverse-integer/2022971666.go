func num_of_digits(x int) int {
    count := 0

    for x != 0 {
        count += 1
        x = x / 10
    }

    return count
}

func reverse(x int) int {
    is_negative := x < 0
    result := 0

    if is_negative {
        x = -x
    }

    num_digits := num_of_digits(x)

    for x != 0 {
        rem := x % 10
        result = result + int(math.Pow10(num_digits - 1)) * rem 
        num_digits--
        x = x / 10
    }

    if is_negative {  
       
        if(math.MinInt32 > -result) {
            return 0
        }
        return -result
    }

     if(math.MaxInt32 < result) {
        return 0
      }

    return result
}
