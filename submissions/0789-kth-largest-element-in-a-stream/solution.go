import "container/heap"

type minHeap []int

func (mh minHeap) Len() int           { return len(mh) }
func (mh minHeap) Less(i, j int) bool { return mh[i] < mh[j] }
func (mh minHeap) Swap(i, j int)      { mh[i], mh[j] = mh[j], mh[i] }

func (mh *minHeap) Push(x any) { *mh = append(*mh, x.(int)) }

func (mh *minHeap) Pop() any {
	old := *mh
	n := len(old)
	*mh = old[0 : n-1]
	return old[n-1]
}

type KthLargest struct {
	k    int
	nums *minHeap
}

func Constructor(k int, nums []int) KthLargest {
	mh := &minHeap{}
    heap.Init(mh)
	kl := KthLargest{k, mh}

	for _, num := range nums {
		kl.Add(num)
	}

	return kl
}

func (this *KthLargest) Add(val int) int {
	if this.nums.Len() < this.k {
		heap.Push(this.nums, val)
	} else if val > (*this.nums)[0] {
		heap.Pop(this.nums)
		heap.Push(this.nums, val)
	}

	return (*this.nums)[0]
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
