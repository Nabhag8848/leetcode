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
        if str[i] == '*' {
            x := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            i = i + 1
            place := 10
            num := int(str[i] - '0')
            j := i + 1
            for j < len(str) && str[j] >= '0' && str[j] <= '9' {
                num = (num * place) + int(str[j] - '0')
                j++
            }

            value := x * num
            stack = append(stack, value)
            i = j -1 
        } else if str[i] == '-' {
            i = i + 1
            num := int(str[i] - '0')
            j := i + 1
            for j < len(str) && str[j] >= '0' && str[j] <= '9' {
                num = num*10 + int(str[j]-'0')
                j++
            }
            stack = append(stack, -num)
            i = j - 1
        } else if str[i] != '+' {
            place := 10
            num := int(str[i] - '0')
            j := i + 1
            for j < len(str) && str[j] >= '0' && str[j] <= '9' {
                num = (num * place) + int(str[j] - '0')
                j++
            }

            stack = append(stack, num)
            i = j - 1
        }
    }

    sum := 0
    for _, v := range stack {
        sum += v
    }

    return sum == target
}