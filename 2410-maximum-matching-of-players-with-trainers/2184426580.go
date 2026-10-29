func matchPlayersAndTrainers(players []int, trainers []int) int {
    sort.Ints(players)
    sort.Ints(trainers)

    i := 0
    j := 0

    result := 0

    for i < len(players) && j < len(trainers) {
        if trainers[j] >= players[i] {
            result++
            i++
        }

        j++
    }

    return result
}
