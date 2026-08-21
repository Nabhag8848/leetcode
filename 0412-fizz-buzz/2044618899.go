func fizzBuzz(n int) []string {
    s := make([]string, 0)
    for i := 1; i <= n; i++ {
        if i % 3 == 0 && i % 5 == 0 {
            s = append(s, "FizzBuzz")
        } else if i % 3 == 0 {
            s = append(s, "Fizz")
        } else if i % 5 == 0 {
            s = append(s, "Buzz")
        } else {
            s = append(s, strconv.Itoa(i))
        }
    }

    return s
}