func checkValidString(s string) bool {
    mini := 0
    maxi := 0

    for i := range s {
        if s[i] == '(' {
            mini++
            maxi++
        } else if s[i] == ')'{
            mini--
            maxi--
        } else {
            mini--
            maxi++
        }

        if mini < 0 {
            mini = 0
        }

        if maxi < 0 {
            return false
        }
    }

    return mini == 0
}
