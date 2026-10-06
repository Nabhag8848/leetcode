type MedianFinder struct {
    data []int    
}


func Constructor() MedianFinder {
    return MedianFinder{
        data: make([]int, 0),
    }
}


func (this *MedianFinder) AddNum(num int)  {
    i := sort.SearchInts(this.data, num)

    this.data = append(this.data, 0)
    copy(this.data[i+1:], this.data[i:])
    this.data[i] = num
}


func (this *MedianFinder) FindMedian() float64 {
    length := len(this.data)
    is_odd := (length % 2) == 1

    first := (length - 1) / 2
    second := first + 1

    if is_odd || second >= length {
        return float64(this.data[first])
    }

    
    return float64(this.data[first] + this.data[second]) / float64(2)
}



/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */