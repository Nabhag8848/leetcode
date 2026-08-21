func isPerfectSquare(num int) bool {
    return mySqrt(num)
}

func mySqrt(x int) bool {
    if x <= 1 {
        return true
    }


    i:= 1
    j:= x

    mid := i + (j - i) / 2

    for i <= j {
        mid = i + (j - i) / 2
        if mid > x / mid {
            j = mid - 1
        } else if mid < x / mid {
            i = mid + 1
        } else {
            if mid * mid == x {
                return true
            }
            i = mid + 1
        }
    }

    return false
}