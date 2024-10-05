package l0347

type numFrequent struct {
	num      int
	frequent int
}

func topKFrequent(nums []int, k int) []int {
	cnt := make(map[int]int)
	for _, num := range nums {
		cnt[num]++
	}

	var sortNumFrequent []*numFrequent
	for i, i2 := range cnt {
		sortNumFrequent = append(sortNumFrequent, &numFrequent{
			num:      i,
			frequent: i2,
		})
	}

	getMaxHeap(sortNumFrequent)

	var ret []int
	size := len(sortNumFrequent)
	for i := 0; i < k-1; i++ {
		ret = append(ret, sortNumFrequent[0].num)
		sortNumFrequent[0], sortNumFrequent[size-1] = sortNumFrequent[size-1], sortNumFrequent[0]
		size--
		transferMax(sortNumFrequent, 0, size)
	}
	ret = append(ret, sortNumFrequent[0].num)

	return ret
}

func getMaxHeap(sortNumFrequent []*numFrequent) {
	for i := len(sortNumFrequent)/2 - 1; i >= 0; i-- {
		transferMax(sortNumFrequent, i, len(sortNumFrequent))
	}
}

func transferMax(sortNumFrequent []*numFrequent, index, size int) {
	if index >= size || len(sortNumFrequent) < size {
		return
	}

	left, right, maxIndex := 2*index+1, 2*index+2, index

	if left < size && sortNumFrequent[left].frequent > sortNumFrequent[maxIndex].frequent {
		maxIndex = left
	}

	if right < size && sortNumFrequent[right].frequent > sortNumFrequent[maxIndex].frequent {
		maxIndex = right
	}

	if maxIndex != index {
		sortNumFrequent[maxIndex], sortNumFrequent[index] = sortNumFrequent[index], sortNumFrequent[maxIndex]
		transferMax(sortNumFrequent, maxIndex, size)
	}
}
