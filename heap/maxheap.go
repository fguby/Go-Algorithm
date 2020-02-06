package heap

import (
	"GoAlgorithm/utils"
	"errors"
	"fmt"
)

type MaxHeap struct {
	prev       int
	arr        []interface{}
	comparator utils.Comparator
}

func assertHeapImplementation() {
	var _ heap = (*MaxHeap)(nil)
}

func NewWithComparator(c utils.Comparator, cap int) *MaxHeap {
	return &MaxHeap{
		comparator: c,
		arr:        make([]interface{}, 0, cap),
	}
}

func NewWithDefault() *MaxHeap {
	return &MaxHeap{
		comparator: utils.IntComparator,
		arr:        make([]interface{}, 0, 10),
	}
}

func (m *MaxHeap) GetSize() int {
	return len(m.arr)
}

func (m *MaxHeap) IsEmpty() bool {
	return len(m.arr) > 0
}

func (m *MaxHeap) GetChild(k int) interface{} {
	if k < 0 || k >= m.GetSize() {
		return nil
	}
	return m.arr[k]
}

func (m *MaxHeap) GetParent(k int) (int, error) {
	if k <= 0 {
		return -1, errors.New("Not a valid index")
	}
	return (k - 1) / 2, nil
}

func (m *MaxHeap) GetLeftChild(k int) (int, error) {
	if k < 0 {
		return -1, errors.New("Not a valid index")
	}
	if 2*k+1 >= m.GetSize() {
		return -1, errors.New("index out of bound")
	}
	return 2*k + 1, nil
}

func (m *MaxHeap) GetRightChild(k int) (int, error) {
	if k < 0 {
		return -1, errors.New("Not a valid index")
	}
	if 2*k+2 >= m.GetSize() {
		return -1, errors.New("index out of bound")
	}
	return 2*k + 2, nil
}

//添加元素
func (m *MaxHeap) InsertChild(e interface{}) int {
	m.arr = append(m.arr, e)
	//上浮
	m.SiftUp(len(m.arr) - 1)
	return len(m.arr)
}

//上浮
func (m *MaxHeap) SiftUp(k int) {
	if k == 0 {
		return
	}

	parent, _ := m.GetParent(k)

	v, _ := m.comparator(m.arr[parent], m.arr[k])

	if v >= 0 {
		return
	} else {
		m.arr[parent], m.arr[k] = m.arr[k], m.arr[parent]
		m.SiftUp(parent)
	}
}

//下沉
func (m *MaxHeap) SiftDown(k int) {
	if k == m.GetSize()-1 {
		return
	}

	left, _ := m.GetLeftChild(k)
	right, _ := m.GetRightChild(k)

	if left > 0 && right > 0 {
		v, _ := m.comparator(m.arr[left], m.arr[right])
		if v < 0 {
			left = right
		}
	} else if left == -1 {
		return
	}

	if left <= m.GetSize()-1 {
		value, _ := m.comparator(m.arr[left], m.arr[k])
		if value < 0 {
			return
		}
		m.arr[left], m.arr[k] = m.arr[k], m.arr[left]
		m.SiftDown(left)
	} else {
		return
	}

}

func (m *MaxHeap) Replace(k int, e interface{}) {
	//
	if k < 0 {
		return
	}

	m.arr[k] = e
	left, _ := m.GetLeftChild(k)
	if v, _ := m.comparator(e, m.arr[left]); v > 0 {
		m.SiftUp(k)
	} else if v < 0 {
		m.SiftDown(k)
	}
}

func (m *MaxHeap) Delete(k int) {
	if k < 0 {
		return
	}

	lastIndex := len(m.arr) - 1

	m.arr[k], m.arr[lastIndex] = m.arr[lastIndex], nil
	m.SiftDown(k)
}

func (m *MaxHeap) Heapify(nums []interface{}) {

	//
	if len(nums) <= 1 {
		return
	}

	m.arr = nums

	lastParentNode := (len(nums) - 1) / 2

	for i := lastParentNode; i >= 0; i-- {
		m.SiftDown(i)
	}

}

func (m *MaxHeap) PrintAll() {
	for _, v := range m.arr {
		fmt.Println(v)
	}
}
