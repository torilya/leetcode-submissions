import (
    "fmt"
	"slices"
)

type KthLargest struct {
	k    int
	nums []int
}

func Constructor(k int, nums []int) KthLargest {

	if k > len(nums)+1 {
		panic("k > nums.length")
	}

    kl := KthLargest{k: k}

	for _, n := range nums {
        kl.Add(n)
    }

	return kl
}

func (this *KthLargest) Add(val int) int {
    if len(this.nums) < this.k || val > this.nums[0] {
        i, _ := slices.BinarySearch(this.nums, val)
        this.nums = slices.Insert(this.nums, i, val)

        if len(this.nums) > this.k {
            this.nums = this.nums[1:]
        }
    }

	return this.nums[0]
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
