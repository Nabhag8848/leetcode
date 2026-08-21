func isPalindrome(x int) bool {
    if (x < 0) {
        return false
    }

    reverse_x := reverse(x)
    return reverse_x == x
}

func num_of_digits(x int) int {
    count := 0

    for x != 0 {
        count += 1
        x = x / 10
    }

    return count
}

func reverse(x int) int {
    result := 0

    num_digits := num_of_digits(x)

    for x != 0 {
        rem := x % 10
        result = result + int(math.Pow10(num_digits - 1)) * rem 
        num_digits--
        x = x / 10
    }

    return result
}
