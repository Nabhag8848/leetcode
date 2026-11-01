func maximumUnits(boxTypes [][]int, truckSize int) int {
    sort.Slice(boxTypes, func (a, b int) bool {
        return boxTypes[a][1] > boxTypes[b][1]
    })

    result := 0
    capacity := truckSize

    for _,box := range boxTypes {
        num := box[0]
        units := box[1]

        if num <= capacity {
            result = result + (num * units)
            capacity -= num
        } else {
            result  = result + (capacity * units)
            break
        }
    }

    return result 
}
