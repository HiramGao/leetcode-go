package main

import "fmt"

func QuickSort(arr []int) []int {

	var sort func(int, int)
	var partition func(l, r int) int

	partition = func(l, r int) int {
		lCursor, RCursor := l+1, r
		pivot := arr[l]

		for {
			for lCursor < r && arr[lCursor] < pivot {
				lCursor++
			}
			for RCursor > l && arr[RCursor] > pivot {
				RCursor--

			}
			if lCursor < RCursor {
				arr[lCursor], arr[RCursor] = arr[RCursor], arr[lCursor]

			} else {
				break
			}
		}
		arr[l], arr[RCursor] = arr[RCursor], arr[l]
		fmt.Println(lCursor, RCursor, arr)
		fmt.Println("+++++++++++++++++++")
		return RCursor
	}

	sort = func(l int, r int) {
		if l < r {
			mid := partition(l, r)
			sort(l, mid-1)
			sort(mid+1, r)
		}
	}

	sort(0, len(arr)-1)

	return arr
}

func main() {
	arr := []int{4, 5, 6, 0, 1, 2, 3, 0}
	arr = QuickSort(arr)
	fmt.Print(arr)
}
