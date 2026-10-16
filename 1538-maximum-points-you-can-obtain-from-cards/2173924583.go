func maxScore(cardPoints []int, k int) int {
    lSum := 0
    rSum := 0
    right := len(cardPoints) - 1

    for i := range k {
        lSum = lSum + cardPoints[i]
    }

    maxSum := lSum

    for i := k - 1; i >= 0; i-- {
        lSum = lSum - cardPoints[i]
        rSum = rSum + cardPoints[right]

        right--
        maxSum = max(maxSum, lSum + rSum)
    }

    return maxSum
}