func countDigits(num int) int {
    count := 0
    original_num := num

    for num > 0 {
        rem := num % 10
        if original_num % rem == 0 {
            count++
        }

        num = num / 10

    }

    return count
}