func mySqrt(x int) int {
    if x <= 1 {
        return x
    }


    i:= 1
    j:= x

    mid := i + (j - i) / 2
    ans := 0

    for i <= j {
        mid = i + (j - i) / 2
        if mid > x / mid {
            j = mid - 1
        } else if mid < x / mid {
            ans = mid
            i = mid + 1
        } else {
            return mid
        }
    }

    return ans
}