package main

import "fmt"

func findKthLargest(nums []int, k int) int {
	heapSize := len(nums)
	var adjustHeap func(int, int)

	adjustHeap = func(i int, heapSize int) {
		largest, lSon, rSon := i, i*2+1, i*2+2
		if lSon < heapSize && nums[largest] < nums[lSon] {
			nums[largest], nums[lSon] = nums[lSon], nums[largest]
			adjustHeap(lSon, heapSize)
		}
		if rSon < heapSize && nums[largest] < nums[rSon] {
			nums[largest], nums[rSon] = nums[rSon], nums[largest]
			adjustHeap(rSon, heapSize)
		}
	}
	//建堆
	for i := heapSize/2 - 1; i >= 0; i-- {
		adjustHeap(i, heapSize)
	}
	for i := heapSize - 1; i >= len(nums)-k+1; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		heapSize--

		adjustHeap(0, heapSize)
	}

	return nums[0]
}
func findKthLargest1(nums []int, k int) int {
	n := len(nums)
	var sort func(l, r int) int
	var partition func(l, r int) int

	partition = func(l, r int) int {
		pviot := nums[r]
		i := l - 1
		for j := l; j < r; j++ {
			if nums[j] <= pviot {
				i++
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
		nums[i+1], nums[r] = nums[r], nums[i+1]
		return i + 1
	}

	sort = func(l, r int) int {

		if l < r {
			mid := partition(l, r)
			if mid == n-k {
				return nums[mid]
			} else if mid < n-k {
				return sort(mid+1, r)
			}
			return sort(l, mid-1)
		}
		return nums[l]
	}

	return sort(0, n-1)
}
func main() {
	fmt.Println(findKthLargest1([]int{1}, 1))
}
