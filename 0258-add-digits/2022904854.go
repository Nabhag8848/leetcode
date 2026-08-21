func addDigits(num int) int {
    count := 0

    for {
        rem := num % 10
        count += rem
        num = num / 10

        if (num == 0) {

            if (count % 10 == count) {
                return count
            }

            num = count
            count = 0
        }

    }

    return count
}