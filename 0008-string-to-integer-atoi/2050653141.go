func myAtoi(s string) int {
    s = strings.TrimSpace(s)
    if len(s) == 0 {
        return 0
    }

    res := 0
    isNegative := false

    for i := range s {

        if i == 0 {
            if s[i] == '+' {
                isNegative = false
                continue
            } else if s[i] == '-' {
                isNegative = true
                continue
            }
        }

        digit := int(s[i] - '0')

        if digit >= 0 && digit <= 9 {
            res = res * 10 + digit

            if isNegative {
                check := -res
                if check < math.MinInt32 {
                    return math.MinInt32
                }
            } else {
                if res > math.MaxInt32 {
                    return math.MaxInt32
                }
            }
        } else {
            break
        }
    }

    if isNegative {
        return -res
    }

    return res
}