package main

import "fmt"

/**
下标为i的节点的父节点为 (i-1)/2
下标为i的节点的左节点为 i*2+1
下标为i的节点的右节点为 i*2+2
*/
func main() {
	arr := []int{9, 10, 8, 7, 6, 5, 4, 3, 2, 1}
	arr = buildMaxHeapAndSort(arr)
	fmt.Println(arr)
}

func buildMaxHeapAndSort(arr []int) []int {

	heapSize := len(arr)
	var adjustHeap func(int, int)
	adjustHeap = func(i, heapSize int) {
		largest, lson, rson := i, i*2+1, i*2+2
		if lson < heapSize && arr[largest] < arr[lson] {
			largest = lson
		}
		if rson < heapSize && arr[largest] < arr[rson] {
			largest = rson
		}
		if largest != i {
			arr[largest], arr[i] = arr[i], arr[largest]
			adjustHeap(largest, heapSize)
		}
	}
	//建堆
	for i := heapSize/2 - 1; i >= 0; i-- {
		adjustHeap(i, heapSize)
	}
	//排序
	for i := len(arr) - 1; i > 0; i-- {
		arr[i], arr[0] = arr[0], arr[i]
		heapSize--
		adjustHeap(0, heapSize)

	}

	return arr
}
