func addOperators(num string, target int) []string {
    if len(num) == 0 {
        return []string{}
    }

    var result []string
    n := len(num)

    for end := 1; end <= n; end++ {
        first := num[0:end]
        if len(first) > 1 && first[0] == '0' {
            break
        }
        val, _ := strconv.Atoi(first)
        result = append(result, helper(num, target, end, first, val, val)...)
    }

    return result
}

func helper(num string, target int, idx int, str string, curVal int, lastOperand int) []string {
    if idx == len(num) {
        if curVal == target {
            return []string{str}
        }
        return []string{}
    }

    var result []string

    for end := idx + 1; end <= len(num); end++ {
        chunk := num[idx:end]
        if len(chunk) > 1 && chunk[0] == '0' {
            break
        }
        val, _ := strconv.Atoi(chunk)

        result = append(result, helper(num, target, end, str+"+"+chunk, curVal+val, val)...)
        result = append(result, helper(num, target, end, str+"-"+chunk, curVal-val, -val)...)
        result = append(result, helper(num, target, end, str+"*"+chunk, curVal-lastOperand+lastOperand*val, lastOperand*val)...)
    }

    return result
}