func myAtoi(s string) int {
    s = strings.TrimSpace(s)
    return recursive(s, 0, 0, false)
}


func recursive(s string, index int, res int, isNegative bool) int {
    if len(s) == index {
         value, ok := outOfBound(isNegative, res)

        if ok {
            return value
        }

        if isNegative {
            return -res
        }

        return res
    }

    if index == 0 {
        if s[index] == '-' {
           return recursive(s, index + 1, res, true)
        } else if s[index] == '+' { 
           return recursive(s, index + 1, res, false)
        }
    }

    digit := int(s[index] - '0')

    if digit >=0 && digit <= 9 {
        res = res * 10 + digit 
        value, ok := outOfBound(isNegative, res)

        if ok {
            return value
        }

        return recursive(s, index + 1, res, isNegative)
    } else {
        if isNegative {
            return -res
        }
        
        return res
    }

    return 0
}

func outOfBound(isNegative bool, res int) (int,bool) {
    if isNegative {
        check := -res
        if check < math.MinInt32 {
            return math.MinInt32, true
        }
    } else {
        if res > math.MaxInt32 {
            return math.MaxInt32, true
        }
    }

    return 0, false
}