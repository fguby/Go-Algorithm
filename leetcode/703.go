package leetcode

type KthLargest struct {
	k    int
	data []int
}

func Constructor(k int, nums []int) KthLargest {
	kl := KthLargest{
		k:    k,
		data: make([]int, 0, k),
	}

	for _, v := range nums {
		if len(kl.data) < k {
			kl.data = append(kl.data, v)
			kl.SiftUp(len(kl.data) - 1)
		} else {
			if kl.data[0] < v {
				kl.data[0] = v
				kl.SiftDown(0)
			}
		}
	}

	return kl
}

func (this *KthLargest) Add(val int) int {
	if len(this.data) < this.k {
		this.data = append(this.data, val)
		this.SiftUp(len(this.data) - 1)
	} else {
		if this.data[0] < val {
			this.data[0] = val
			this.SiftDown(0)
		}
	}

	return this.data[0]
}

func (this *KthLargest) SiftUp(k int) {

	if k <= 0 {
		return
	}

	parent := this.GetParent(k)

	if this.data[parent] > this.data[k] {
		this.data[k], this.data[parent] = this.data[parent], this.data[k]
	} else {
		return
	}

	this.SiftUp(parent)
}

func (this *KthLargest) SiftDown(k int) {

	if (2*k + 1) >= len(this.data) {
		return
	}

	left := this.GetLeftChild(k)
	right := this.GetRightChild(k)

	if right >= len(this.data) {
		if this.data[k] > this.data[left] {
			this.data[k], this.data[left] = this.data[left], this.data[k]
			this.SiftDown(left)
		}
	} else {
		if this.data[left] > this.data[right] {
			left = right
		}

		if this.data[k] > this.data[left] {
			this.data[k], this.data[left] = this.data[left], this.data[k]
			this.SiftDown(left)
		}
	}

}

func (this *KthLargest) GetLeftChild(k int) int {
	return (2 * k) + 1
}

func (this *KthLargest) GetRightChild(k int) int {
	return (2 * k) + 2
}

func (this *KthLargest) GetParent(k int) int {
	return (k - 1) / 2
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
