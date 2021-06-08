package main

import "fmt"

func topKFrequent(nums []int, k int) []int {
	hashMap := make(map[int]int)
	var ans []int
	var adjustHeap func(int, int)
	for _, num := range nums {
		hashMap[num]++
	}
	n := len(hashMap)
	heapSize := len(hashMap)
	var arr [][2]int
	for key, value := range hashMap {
		arr = append(arr, [2]int{key, value})
	}

	adjustHeap = func(i, heapSize int) {
		largest, lson, rson := i, i*2+1, i*2+2
		if lson < heapSize && arr[largest][1] < arr[lson][1] {
			arr[largest], arr[lson] = arr[lson], arr[largest]
			adjustHeap(lson, heapSize)
		}
		if rson < heapSize && arr[largest][1] < arr[rson][1] {
			arr[largest], arr[rson] = arr[rson], arr[largest]
			adjustHeap(rson, heapSize)
		}
	}
	for i := heapSize/2 - 1; i >= 0; i-- {
		adjustHeap(i, heapSize)
	}
	for i := n - 1; i >= n-k; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		ans = append(ans, arr[i][0])
		heapSize--
		adjustHeap(0, heapSize)
	}
	return ans
}
func main() {
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3, 3, 3, 3}, 2))
}
