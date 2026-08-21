const MOD = 1_000_000_007

func countGoodNumbers(n int64) int {
    even := int64(math.Ceil(float64(n) / 2))
    odd := int64(math.Floor(float64(n) / 2))

    var pow func(x int, n int64) int
    pow = func(x int, n int64) int {
        if n == 0 {
            return 1
        }
        half := pow(x, n/2) % MOD
        half = half * half % MOD

        if n % 2 == 1 {
            half = (x * half) % MOD
        }

        return half % MOD
    }

    return (pow(5, even) * pow(4, odd)) % MOD
}







