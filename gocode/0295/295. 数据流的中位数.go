package l0295

type MedianFinder struct {
	bigNumMinHeap   []int
	smallNumMaxHeap []int
	cnt             int
}

func Constructor() MedianFinder {
	return MedianFinder{}
}

func (this *MedianFinder) AddNum(num int) {
	this.cnt++
	n := (this.cnt + 1) / 2
	if len(this.bigNumMinHeap) == 0 {
		this.bigNumMinHeap = append(this.bigNumMinHeap, num)
		return
	}
	if len(this.bigNumMinHeap) == n {
		if this.bigNumMinHeap[0] < num {
			this.bigNumMinHeap[0], num = num, this.bigNumMinHeap[0]
			// 调整
			changeHeadMinHeap(this.bigNumMinHeap)
		}
		this.smallNumMaxHeap = append(this.smallNumMaxHeap, num)
		changeAppendMaxHeap(this.smallNumMaxHeap)
	} else {
		if this.smallNumMaxHeap[0] > num {
			this.smallNumMaxHeap[0], num = num, this.smallNumMaxHeap[0]
			// 调整
			changeHeadMaxHeap(this.smallNumMaxHeap)
		}
		this.bigNumMinHeap = append(this.bigNumMinHeap, num)
		// 调整
		changeAppendMinHeap(this.bigNumMinHeap)
	}
}

func changeHeadMaxHeap(heap []int) {
	n := len(heap)
	parent := 0
	for parent < n {
		maxN := parent
		left := 2*parent + 1
		right := 2*parent + 2
		if left < n && heap[left] > heap[maxN] {
			maxN = left
		}
		if right < n && heap[right] > heap[maxN] {
			maxN = right
		}
		if maxN == parent {
			break
		}
		heap[parent], heap[maxN] = heap[maxN], heap[parent]
		parent = maxN
	}
}

func changeHeadMinHeap(heap []int) {
	n := len(heap)
	parent := 0
	for parent < n {
		maxN := parent
		left := 2*parent + 1
		right := 2*parent + 2
		if left < n && heap[left] < heap[maxN] {
			maxN = left
		}
		if right < n && heap[right] < heap[maxN] {
			maxN = right
		}
		if maxN == parent {
			break
		}
		heap[parent], heap[maxN] = heap[maxN], heap[parent]
		parent = maxN
	}
}

func changeAppendMaxHeap(heap []int) {
	child := len(heap) - 1
	parent := (child - 1) / 2
	for parent >= 0 && heap[child] > heap[parent] {
		heap[parent], heap[child] = heap[child], heap[parent]
		child = parent
		parent = (child - 1) / 2
	}
}

func changeAppendMinHeap(heap []int) {
	child := len(heap) - 1
	parent := (child - 1) / 2
	for parent >= 0 && heap[child] < heap[parent] {
		heap[parent], heap[child] = heap[child], heap[parent]
		child = parent
		parent = (child - 1) / 2
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.cnt%2 == 1 {
		return float64(this.bigNumMinHeap[0])
	}

	return float64(this.bigNumMinHeap[0]+this.smallNumMaxHeap[0]) / 2.0
}
