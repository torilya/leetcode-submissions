type minHeap []int

func (mh minHeap) heapify(nodeIndex int) {
	if nodeIndex <= 0 {
		return
	}

	parentIndex := (nodeIndex - 1) / 2

	if mh[nodeIndex] < mh[parentIndex] {
		mh[nodeIndex], mh[parentIndex] = mh[parentIndex], mh[nodeIndex]
		mh.heapify(parentIndex)
	}
}

func (mh minHeap) push(x int) minHeap {
	mhNew := append(mh, x)
	mhNew.heapify(len(mhNew) - 1)

	return mhNew
}

func (mh minHeap) heapifyTopDown(nodeIndex int) {
	childLeftIndex := nodeIndex*2 + 1
	childRightIndex := childLeftIndex + 1

	mhLength := len(mh)

	minIndex := nodeIndex

	if childLeftIndex < mhLength && mh[childLeftIndex] < mh[minIndex] {
		minIndex = childLeftIndex
	}
	if childRightIndex < mhLength && mh[childRightIndex] < mh[minIndex] {
		minIndex = childRightIndex
	}

	if nodeIndex != minIndex {
		mh[nodeIndex], mh[minIndex] = mh[minIndex], mh[nodeIndex]
		mh.heapifyTopDown(minIndex)
	}
}

func (mh minHeap) pop() minHeap {
	mh[0] = mh[len(mh)-1]
	mhNew := mh[:len(mh)-1]
	mhNew.heapifyTopDown(0)

	return mhNew
}

type KthLargest struct {
	k    int
	nums minHeap
}

func Constructor(k int, nums []int) KthLargest {
	var mh minHeap
	kl := KthLargest{k, mh}

	for _, num := range nums {
		kl.Add(num)
	}

	return kl
}

func (this *KthLargest) Add(val int) int {
	this.nums = this.nums.push(val)

	if len(this.nums) > this.k {
		this.nums = this.nums.pop()
	}

	return this.nums[0]
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
