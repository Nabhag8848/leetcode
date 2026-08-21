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
        result = append(result, helper(num, target, end, first)...)
    }

    return result
}

func helper(num string, target int, idx int, str string) []string {
    if idx == len(num) {
        if ok := solveExpression(str, target); ok {
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

        result = append(result, helper(num, target, end, str+"+"+chunk)...)
        result = append(result, helper(num, target, end, str+"-"+chunk)...)
        result = append(result, helper(num, target, end, str+"*"+chunk)...)
    }

    return result
}

func solveExpression(str string, target int) bool {
    stack := make([]int, 0)

    for i := 0; i < len(str); i++ {
        switch str[i] {
        case '*':
            num, next := parseNumber(str, i+1)
            x := stack[len(stack)-1]
            stack[len(stack)-1] = x * num
            i = next - 1
        case '-':
            num, next := parseNumber(str, i+1)
            stack = append(stack, -num)
            i = next - 1
        case '+':
        default:
            num, next := parseNumber(str, i)
            stack = append(stack, num)
            i = next - 1
        }
    }

    sum := 0
    for _, v := range stack {
        sum += v
    }

    return sum == target
}

func parseNumber(str string, i int) (int, int) {
    num := int(str[i] - '0')
    j := i + 1
    for j < len(str) && str[j] >= '0' && str[j] <= '9' {
        num = num*10 + int(str[j]-'0')
        j++
    }
    return num, j
}