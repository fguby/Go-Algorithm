package examples

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	nums := []int{2, 4, 1, 9, 3, 6, 10, 5}
	QuickSort(nums)
	fmt.Println(nums)
}
